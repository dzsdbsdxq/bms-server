package wheel

import (
	"bms-server/global"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type AwardBatch struct {
	wheelId int //活动ID
	prizeId int //奖项ID
	// 奖品名称
	prizeName string
	// 剩余奖品总数
	totalBalance int64
	// 奖品总数
	totalAmount int64
	// 上次奖品发放时间
	updateTime int64
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

func (award *AwardBatch) InitAwardPool() {
	global.GVA_REDIS.ZAdd(context.Background(), "xs_22", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: "A",
	})
	global.GVA_REDIS.ZAdd(context.Background(), "xs_22", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: "B",
	})
	global.GVA_REDIS.ZAdd(context.Background(), "xs_22", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: "C",
	})

	global.GVA_REDIS.HSet(context.Background(), "award_balance", "A", 100)
	global.GVA_REDIS.HSet(context.Background(), "award_balance", "B", 200)
	global.GVA_REDIS.HSet(context.Background(), "award_balance", "C", 500)
}

func (award *AwardBatch) WinPrize() {

}
