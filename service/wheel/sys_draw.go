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

type AwardBatch struct {
	wheelId int //活动ID
	prizeId int //奖项ID
	Uuid    string
	// 奖品名称
	prizeName string
	// 剩余奖品总数
	totalBalance int64
	// 奖品总数
	totalAmount int64
	// 上次奖品发放时间
	updateTime int64
	startTime  int64
	endTime    int64
}

func (award *AwardBatch) GetName() string {
	return award.prizeName
}

func (award *AwardBatch) GetTotalBalance() int64 {
	return award.totalBalance
}

func (award *AwardBatch) GetTotalAmount() int64 {
	return award.totalAmount
}

func (award *AwardBatch) GetUpdateTime() int64 {
	return award.updateTime
}

func (award *AwardBatch) SetTotalBalance(totalBalance int64) {
	award.totalBalance = totalBalance
}

func (award *AwardBatch) SetTotalAmount(totalAmount int64) {
	award.totalAmount = totalAmount
}

func (award *AwardBatch) SetUpdateTime(updateTime int64) {
	award.updateTime = updateTime
}

func (award *AwardBatch) SetName(name string) {
	award.prizeName = name

}

func (award *AwardBatch) SetAwardBatch(aw AwardBatch) {
	award.Uuid = aw.Uuid
	award.wheelId = aw.wheelId
	award.prizeId = aw.prizeId
	award.prizeName = aw.prizeName
	award.totalBalance = aw.totalBalance
	award.totalAmount = aw.totalAmount
	award.updateTime = aw.updateTime
	award.startTime = aw.startTime
	award.endTime = aw.endTime
}
func (award *AwardBatch) getAwardInfoKey() string {
	return fmt.Sprintf("awardInfo_%d", award.wheelId)
}

func (award *AwardBatch) getAwardBalanceKey() string {
	return fmt.Sprintf("awardBalance_%d", award.wheelId)
}

func (award *AwardBatch) InitAwardPool() {

	global.GVA_REDIS.ZAdd(context.Background(), award.getAwardInfoKey(), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: award.prizeId,
	})
	global.GVA_REDIS.HSet(context.Background(), award.getAwardBalanceKey(), award.prizeId, award.totalBalance)
}

func (award *AwardBatch) GetAwardBatch() *AwardBatch {
	awardBatch := award.RandomGetAwardBatch()
	if awardBatch == nil {
		log.Println("sorry, you didn't win the prize.")
		return nil
	}

	// 判断是否到达奖品释放时间点
	random := rand.New(rand.NewSource(award.GetUpdateTime()))

	detalTime := (award.endTime - award.startTime) / awardBatch.totalAmount

	// 计算下一个奖品释放的时间点
	releaseTime := award.startTime + (award.totalAmount-award.totalBalance)*detalTime + int64(random.Int())%detalTime

	log.Println("relaseTime : ", time.Unix(releaseTime, 0).Format("2006-01-02 15:04:05"))
	if time.Now().Unix() < releaseTime {
		// 未到达奖品释放的时间点，即不中奖
		log.Println("sorry, you didn't win the prize")
		return nil
	}
	return awardBatch
}

func (award *AwardBatch) RandomGetAwardBatch() *AwardBatch {
	// global.GVA_REDIS.Int64Map()
	retMap := global.GVA_REDIS.HGetAll(context.Background(), award.getAwardBalanceKey())

	fmt.Println("retMap:", retMap)

	totalBalance := int64(0)
	for _, value := range retMap.Val() {
		i, _ := strconv.ParseInt(value, 10, 64)
		totalBalance += i
	}

	if totalBalance == 0 {
		global.GVA_LOG.Info("total balance is ", zap.String("totalBalance", strconv.FormatInt(totalBalance, 10)))
		return nil
	}

	awardBatches := award.GetAllAwardBatch()
	for index, awardBatch := range awardBatches {
		awardBatches[index].totalBalance = awardBatch.totalBalance
	}

	log.Println("awardBatches :", awardBatches)

	random := rand.New(rand.NewSource(totalBalance))

	num := random.Int63n(totalBalance)

	for _, awardBatch := range awardBatches {

		// 奖品已经抽完
		if awardBatch.GetTotalBalance() <= 0 {
			log.Println("奖品已经抽完")
			continue
		}

		num = num - awardBatch.GetTotalBalance()
		if num < 0 {
			return &awardBatch
		}
	}

	return nil

}

func (award *AwardBatch) GetAllAwardBatch() []AwardBatch {
	var awardBatches []AwardBatch
	// 获取所有奖品数
	awardInfoKey := award.getAwardInfoKey()

	values := global.GVA_REDIS.ZRangeWithScores(context.Background(), awardInfoKey, 0, -1).Val()
	if len(values) == 0 {
		log.Println("get all award redis error ")
	}

	for index, value := range values {
		if index%2 == 0 {
			awardBatches = append(awardBatches, AwardBatch{
				prizeId: value.Member.(int),
			})
		} else {
			lastUpdateTime, _ := strconv.ParseInt(string(value.Member.([]byte)), 10, 64)

			awardBatches[index/2].SetUpdateTime(lastUpdateTime)

		}
	}

	// 填充 totalAmount
	for index, awardBatch := range awardBatches {
		awardBatches[index].SetTotalAmount(awardBatch.totalAmount)
	}

	return awardBatches

}

func (award *AwardBatch) WinPrize() {

}
