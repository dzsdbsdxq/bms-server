package wheel

import (
	"bms-server/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type AwardBatchService struct {
	WheelId int //活动ID
	PrizeId int //奖项ID
	Uuid    string
	// 奖品名称
	PrizeName string
	// 剩余奖品总数
	TotalBalance int64
	// 奖品总数
	TotalAmount int64
	// 上次奖品发放时间
	UpdateTime int64
	StartTime  int64
	EndTime    int64
}

func (award *AwardBatchService) SetAwardBatchService(aw *AwardBatchService) *AwardBatchService {
	award.WheelId = aw.WheelId
	award.PrizeId = aw.PrizeId
	award.Uuid = aw.Uuid
	award.PrizeName = aw.PrizeName
	award.TotalBalance = aw.TotalBalance
	award.TotalAmount = aw.TotalAmount
	award.UpdateTime = aw.UpdateTime
	award.StartTime = aw.StartTime
	award.EndTime = aw.EndTime
	return award
}

func getAwardInfoKey(wheelId int) string {
	fmt.Println("awardInfo_", wheelId)
	return fmt.Sprintf("awardInfo_%d", wheelId)
}

func getAwardBalanceKey(wheelId int) string {
	return fmt.Sprintf("awardBalance_%d", wheelId)
}

func InitAwardPool(award *AwardBatchService) {

	global.GVA_REDIS.ZAdd(context.Background(), getAwardInfoKey(award.WheelId), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: award.PrizeId,
	})
	global.GVA_REDIS.HSet(context.Background(), getAwardBalanceKey(award.WheelId), award.PrizeId, award.TotalBalance)
}

func GetAwardBatch(wheelId int) *AwardBatchService {
	awardBatch := RandomGetAwardBatch(wheelId)
	if awardBatch == nil {
		log.Println("sorry, you didn't win the prize.")
		return nil
	}

	// 判断是否到达奖品释放时间点
	random := rand.New(rand.NewSource(awardBatch.UpdateTime))

	detalTime := (awardBatch.EndTime - awardBatch.StartTime) / awardBatch.TotalAmount

	// 计算下一个奖品释放的时间点
	releaseTime := awardBatch.StartTime + (awardBatch.TotalAmount-awardBatch.TotalBalance)*detalTime + int64(random.Int())%detalTime

	log.Println("relaseTime : ", time.Unix(releaseTime, 0).Format("2006-01-02 15:04:05"))
	if time.Now().Unix() < releaseTime {
		// 未到达奖品释放的时间点，即不中奖
		log.Println("sorry, you didn't win the prize")
		return nil
	}
	return awardBatch
}

func RandomGetAwardBatch(wheelId int) *AwardBatchService {
	// global.GVA_REDIS.Int64Map()
	retMap := global.GVA_REDIS.HGetAll(context.Background(), getAwardBalanceKey(wheelId))
	totalBalance := int64(0)
	fmt.Println("retMap:", retMap)
	for _, value := range retMap.Val() {
		i, _ := strconv.ParseInt(value, 10, 64)
		totalBalance += i
	}
	fmt.Println("totalBalance:", totalBalance)

	if totalBalance == 0 {
		global.GVA_LOG.Info("total balance is ", zap.String("totalBalance", strconv.FormatInt(totalBalance, 10)))
		return nil
	}

	awardBatches := GetAllAwardBatch(wheelId)
	for index, awardBatch := range awardBatches {
		awardBatches[index].TotalBalance = awardBatch.TotalBalance
	}
	log.Println("awardBatches :", awardBatches)

	random := rand.New(rand.NewSource(totalBalance))

	log.Println("random :", random)
	num := random.Int63n(totalBalance)
	log.Println("num :", num)
	for _, awardBatch := range awardBatches {
		log.Println("awardBatch.TotalBalance :", awardBatch.TotalBalance)
		// 奖品已经抽完
		if awardBatch.TotalBalance <= 0 {
			log.Println("奖品已经抽完")
			continue
		}

		num = num - awardBatch.TotalBalance
		if num < 0 {
			return &awardBatch
		}
	}

	return nil

}

func GetAllAwardBatch(wheelId int) []AwardBatchService {
	var awardBatches []AwardBatchService
	// 获取所有奖品数
	values := global.GVA_REDIS.ZRangeWithScores(context.Background(), getAwardInfoKey(wheelId), 0, -1).Val()
	if len(values) == 0 {
		log.Println("get all award redis error ")
	}
	fmt.Println("values:", values)

	for index, value := range values {
		fmt.Println("value.Member:", value.Member, value.Member.(string))
		if index%2 == 0 {
			parseInt, err := strconv.ParseInt(value.Member.(string), 10, 64)
			if err != nil {
				return nil
			}
			awardBatches = append(awardBatches, AwardBatchService{
				PrizeId: int(parseInt),
			})
		} else {
			lastUpdateTime, _ := strconv.ParseInt(value.Member.(string), 10, 64)
			fmt.Println("lastUpdateTime:", lastUpdateTime)
			awardBatches[index/2].UpdateTime = lastUpdateTime

		}
	}
	fmt.Println("awardBatches:", awardBatches)

	// 填充 totalAmount
	for index, awardBatch := range awardBatches {
		awardBatches[index].TotalAmount = awardBatch.TotalAmount
	}
	fmt.Println("awardBatches2:", awardBatches)

	return awardBatches

}

func WinPrize(wheelId int) *AwardBatchService {
	awardBatch := GetAwardBatch(wheelId)

	if awardBatch == nil {
		return awardBatch
	}
	fmt.Println("awardBatch:", awardBatch)

	err := global.GVA_REDIS.Watch(context.Background(), nil, getAwardBalanceKey(awardBatch.WheelId))
	if err != nil {
		return nil
	}
	pipe := global.GVA_REDIS.TxPipeline()
	pipe.ZAdd(context.Background(), getAwardInfoKey(awardBatch.WheelId), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: awardBatch.PrizeId,
	})
	pipe.HSet(context.Background(), getAwardBalanceKey(awardBatch.WheelId), awardBatch.PrizeId, awardBatch.TotalBalance-1)
	_, err = pipe.Exec(context.Background())
	if err != nil {
		return nil
	}
	log.Println("congratulations , you won ", awardBatch.PrizeName)

	// 保存用户中奖纪录
	//awardTime := time.Unix(awardBatch.GetUpdateTime(), 0).Format("2006-01-02 15:04:05")
	//	userName := req.Form.Get("user_name")
	//mysqlWrapper.SaveRecords(awardBatch.GetName(), awardTime, username)

	return awardBatch
}
