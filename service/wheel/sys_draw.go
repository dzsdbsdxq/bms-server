package wheel

import (
	"bms-server/global"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var lock = sync.Mutex{}

type AwardBatchService struct {
	WheelId      int //活动ID
	PrizeId      int //奖项ID
	Uuid         string
	IsHit        int    //中奖项，0=否，1=是
	PrizeName    string // 奖品名称
	TotalBalance int64  // 剩余奖品总数
	TotalAmount  int64  // 奖品总数
	UpdateTime   int64  // 上次奖品发放时间
	StartTime    int64
	EndTime      int64
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
	return fmt.Sprintf("awardInfo_%d", wheelId)
}

func getAwardBalanceKey(wheelId int) string {
	return fmt.Sprintf("awardBalance_%d", wheelId)
}
func getAwardSourceInfoKey(wheelId int) string {
	return fmt.Sprintf("awardSourceInfo_%d", wheelId)
}
func InitAwardPool(award *AwardBatchService) {
	awardJson, _ := jsoniter.Marshal(award)
	global.GVA_REDIS.ZAdd(context.Background(), getAwardInfoKey(award.WheelId), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: award.PrizeId,
	})
	global.GVA_REDIS.HSet(context.Background(), getAwardBalanceKey(award.WheelId), award.PrizeId, award.TotalBalance)
	global.GVA_REDIS.HSet(context.Background(), getAwardSourceInfoKey(award.WheelId), award.PrizeId, string(awardJson))
}

func GetAwardBatch(wheelId int, allowWinPrize bool) *AwardBatchService {
	awardBatch := RandomGetAwardBatch(wheelId, allowWinPrize)
	if awardBatch == nil {
		log.Println("sorry, you didn't win the prize.")
		return nil
	}
	fmt.Println("中奖：", awardBatch)

	// 判断是否到达奖品释放时间点
	random := rand.New(rand.NewSource(awardBatch.UpdateTime))

	detalTime := (awardBatch.EndTime - awardBatch.StartTime) / awardBatch.TotalAmount

	// 计算下一个奖品释放的时间点
	releaseTime := awardBatch.StartTime + (awardBatch.TotalAmount-awardBatch.TotalBalance)*detalTime + int64(random.Int())%detalTime

	if time.Now().Unix() < releaseTime {
		// 未到达奖品释放的时间点，即不中奖
		log.Println("sorry, you didn't win the prize")
		return nil
	}
	return awardBatch
}

func RandomGetAwardBatch(wheelId int, allowWinPrize bool) *AwardBatchService {
	retMap := global.GVA_REDIS.HGetAll(context.Background(), getAwardBalanceKey(wheelId)).Val()
	totalBalance := int64(0)
	for _, value := range retMap {
		i, _ := strconv.ParseInt(value, 10, 64)
		totalBalance += i
	}

	if totalBalance == 0 {
		global.GVA_LOG.Info("total balance is ", zap.String("totalBalance", strconv.FormatInt(totalBalance, 10)))
		return nil
	}

	awardBatches := GetAllAwardBatch(wheelId)
	for index, awardBatch := range awardBatches {
		awardBatches[index].TotalBalance, _ = strconv.ParseInt(retMap[strconv.Itoa(awardBatch.PrizeId)], 10, 64) //awardBatch.TotalBalance //retMap[strconv.Itoa(awardBatch.PrizeId)]
	}
	random := rand.New(rand.NewSource(totalBalance))

	num := random.Int63n(totalBalance)
	for _, awardBatch := range awardBatches {
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

	awardSourceInfoStr := global.GVA_REDIS.HGetAll(context.Background(), getAwardSourceInfoKey(wheelId)).Val()

	var aw AwardBatchService
	for _, value := range values {
		parseInt, _ := strconv.ParseInt(value.Member.(string), 10, 64)
		_ = jsoniter.UnmarshalFromString(awardSourceInfoStr[strconv.FormatInt(parseInt, 10)], &aw)
		awardBatches = append(awardBatches, AwardBatchService{
			WheelId:     aw.WheelId,
			Uuid:        aw.Uuid,
			IsHit:       aw.IsHit,
			PrizeName:   aw.PrizeName,
			PrizeId:     int(parseInt),
			UpdateTime:  int64(value.Score),
			TotalAmount: aw.TotalAmount,
			StartTime:   aw.StartTime,
			EndTime:     aw.EndTime,
		})
	}
	return awardBatches
}

func WinPrize(wheelId int, allowWinPrize bool) *AwardBatchService {

	lock.Lock()

	defer lock.Unlock()
	awardBatch := GetAwardBatch(wheelId, allowWinPrize)

	if awardBatch == nil {
		return awardBatch
	}

	err := global.GVA_REDIS.Watch(context.Background(), func(tx *redis.Tx) error {
		return nil
	}, getAwardBalanceKey(awardBatch.WheelId))
	if err != nil {
		return nil
	}
	pipe := global.GVA_REDIS.TxPipeline()
	pipe.ZAdd(context.Background(), getAwardInfoKey(awardBatch.WheelId), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: awardBatch.PrizeId,
	})
	// redis中的库存-1
	pipe.HSet(context.Background(), getAwardBalanceKey(awardBatch.WheelId), awardBatch.PrizeId, awardBatch.TotalBalance-1)
	_, err = pipe.Exec(context.Background())
	if err != nil {
		lock.Unlock()
		return nil
	}
	log.Println("congratulations , you won ", awardBatch.PrizeName)

	//更新 返回的库存
	awardBatch.TotalBalance = awardBatch.TotalBalance - 1
	// 保存用户中奖纪录
	//awardTime := time.Unix(awardBatch.GetUpdateTime(), 0).Format("2006-01-02 15:04:05")
	//	userName := req.Form.Get("user_name")
	//mysqlWrapper.SaveRecords(awardBatch.GetName(), awardTime, username)

	return awardBatch
}
