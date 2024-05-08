package wheel

import (
	"bms-server/global"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock = sync.Mutex{}

type AwardBatchService struct {
	WheelId      int   //活动ID
	PrizeId      int   //奖项ID
	Probability  int   //概率
	PrizeType    int   //奖项类型 0=虚拟，1=实物
	IsHit        int   //中奖项，0=否，1=是
	TotalBalance int   // 剩余奖品总数
	TotalAmount  int   // 奖品总数
	UpdateTime   int64 // 上次奖品发放时间
	StartTime    int64
	EndTime      int64
	Perc         float64
	Uuid         string
	PrizeName    string // 奖品名称
}

func (award *AwardBatchService) SetAwardBatchService(aw *AwardBatchService) *AwardBatchService {
	award.WheelId = aw.WheelId
	award.PrizeId = aw.PrizeId
	award.PrizeType = aw.PrizeType
	award.Probability = aw.Probability
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

func GetAwardBatch(wheelId int, allowWinPrize bool, algType int) *AwardBatchService {
	awardBatch, noHitAwardBatchs := RandomGetAwardBatch(wheelId, allowWinPrize, algType)
	tmpRd := rand.New(rand.NewSource(int64(len(noHitAwardBatchs))))
	//中奖标志，默认设置未中奖项
	prizeAwardBatch := noHitAwardBatchs[tmpRd.Int63n(int64(len(noHitAwardBatchs)))]
	// 如果是不允许中奖，直接返回，无需判断奖品释放时间点
	if !allowWinPrize || awardBatch == nil {
		return prizeAwardBatch
	}
	// 判断是否到达奖品释放时间点
	random := rand.New(rand.NewSource(awardBatch.UpdateTime))

	detalTime := (awardBatch.EndTime - awardBatch.StartTime) / awardBatch.TotalAmount

	// 计算下一个奖品释放的时间点
	releaseTime := awardBatch.StartTime + (awardBatch.TotalAmount-awardBatch.TotalBalance)*detalTime + int64(random.Int())%detalTime

	if time.Now().Unix() < releaseTime {
		// 未到达奖品释放的时间点，即不中奖，不中奖，返回谢谢参与
		if len(noHitAwardBatchs) != 0 {
			return prizeAwardBatch
		}
	}
	return awardBatch
}

func RandomGetAwardBatch(wheelId int, allowWinPrize bool, algType int) (*AwardBatchService, []*AwardBatchService) {
	retMap := global.GVA_REDIS.HGetAll(context.Background(), getAwardBalanceKey(wheelId)).Val()
	totalBalance := int64(0)
	for _, value := range retMap {
		i, _ := strconv.ParseInt(value, 10, 64)
		totalBalance += i
	}

	//临时存储未中奖项
	noHitAwardBatch := make([]*AwardBatchService, 0)

	awardBatches := GetAllAwardBatch(wheelId)
	for index, awardBatch := range awardBatches {
		awardBatches[index].TotalBalance, _ = strconv.ParseInt(retMap[strconv.Itoa(awardBatch.PrizeId)], 10, 64) //awardBatch.TotalBalance //retMap[strconv.Itoa(awardBatch.PrizeId)]
		if awardBatches[index].IsHit == 0 {
			noHitAwardBatch = append(noHitAwardBatch, &awardBatches[index])
		}
	}
	if totalBalance <= 0 {
		global.GVA_LOG.Info("total balance is ", zap.String("totalBalance", strconv.FormatInt(totalBalance, 10)))
		return nil, noHitAwardBatch
	}

	if algType == 1 {
		random := rand.New(rand.NewSource(totalBalance))
		num := random.Int63n(totalBalance)
		for _, awardBatch := range awardBatches {
			if awardBatch.TotalBalance <= 0 {
				// 奖品已经抽完
				continue
			}
			// 是否允许中奖，如果不允许中奖，则返回谢谢参与奖
			if allowWinPrize {
				num = num - awardBatch.TotalBalance
				if num < 0 {
					return &awardBatch, noHitAwardBatch
				}
			}
		}
	} else {
		index := aliasMethod(awardBatches)
		fmt.Println("index:", index)
		//index := dispersed(awardBatches)
		if awardBatches[index].TotalBalance > 0 {
			return &awardBatches[index], noHitAwardBatch
		}
	}
	return nil, noHitAwardBatch
}

func GetAllAwardBatch(wheelId int) []AwardBatchService {
	var awardBatches []AwardBatchService
	// 获取所有奖品数
	values := global.GVA_REDIS.ZRangeWithScores(context.Background(), getAwardInfoKey(wheelId), 0, -1).Val()
	if len(values) == 0 {
		log.Println("get all award redis error ")
		return nil
	}

	awardSourceInfoStr := global.GVA_REDIS.HGetAll(context.Background(), getAwardSourceInfoKey(wheelId)).Val()

	var aw AwardBatchService
	for _, value := range values {
		parseInt, _ := strconv.ParseInt(value.Member.(string), 10, 64)
		_ = jsoniter.UnmarshalFromString(awardSourceInfoStr[strconv.FormatInt(parseInt, 10)], &aw)
		awardBatches = append(awardBatches, AwardBatchService{
			WheelId:     aw.WheelId,
			Uuid:        aw.Uuid,
			PrizeType:   aw.PrizeType,
			Probability: aw.Probability,
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

func WinPrize(wheelId int, allowWinPrize bool, algType int) *AwardBatchService {
	lock.Lock()
	defer lock.Unlock()
	awardBatch := GetAwardBatch(wheelId, allowWinPrize, algType)
	if !allowWinPrize {
		return awardBatch
	}
	_ = global.GVA_REDIS.Watch(context.Background(), func(tx *redis.Tx) error {
		return nil
	}, getAwardBalanceKey(awardBatch.WheelId))

	pipe := global.GVA_REDIS.TxPipeline()
	pipe.ZAdd(context.Background(), getAwardInfoKey(awardBatch.WheelId), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: awardBatch.PrizeId,
	})
	// redis中的库存-1
	pipe.HSet(context.Background(), getAwardBalanceKey(awardBatch.WheelId), awardBatch.PrizeId, awardBatch.TotalBalance-1)
	_, _ = pipe.Exec(context.Background())
	log.Println("congratulations , you won ", awardBatch.PrizeName)
	//更新 返回的库存
	awardBatch.TotalBalance = awardBatch.TotalBalance - 1
	return awardBatch
}

// /https://juejin.cn/post/7070063466203054088
// 构造容量为10000的容器
func violence(gifts []AwardBatchService) int {
	length := 0
	data := ""
	for index, value := range gifts {
		length += value.Probability
		position := fmt.Sprintf("%d,", index)
		data += strings.Repeat(position, value.Probability)
	}
	random := rand.New(rand.NewSource(10000))
	// 获取[10000,0) 随机数
	index := random.Int63n(10000)

	arr := strings.Split(data, ",")
	giftIndex := cast.ToInt(arr[index])
	return giftIndex
}

// 离散算法
func dispersed(gifts []AwardBatchService) int {
	data := make([]int, 0)
	for index, value := range gifts {
		if index > 0 {
			data = append(data, value.Probability+data[index-1])
		} else {
			data = append(data, value.Probability)
		}
	}

	// 获取[1,0) 随机数
	random := rand.New(rand.NewSource(10000))
	// 获取[10000,0) 随机数
	index := random.Int63n(10000)
	res := binarySearchV3(data, int(index))
	return res
}

// 二分法下界（重复元素 第一个元素）
func binarySearchV3(data []int, target int) int {
	left, right := 0, len(data)

	for left <= right {
		mid := left + (right-left)/2

		if data[mid] > target {
			right = mid - 1
		} else if data[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || data[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return left
}

// 别名算法
func aliasMethod(gifts []AwardBatchService) int {
	// pdf := []float64{0.1, 0.2, 0.3, 0.4}
	// pdf := []float64{0.8, 0.1, 0.1}
	pdf := make([]float64, len(gifts))
	for index, value := range gifts {
		pdf[index] = float64(value.Probability) / 100
	}

	lenth := len(pdf)

	// 原始数据
	probInfo := make([]float64, lenth)

	// 别名补充数据
	alias := make([]int, lenth)

	small := []int{}
	large := []int{}

	// 构造拼装出每一列和都唯一的矩阵
	for i := 0; i < lenth; i++ {
		pdf[i] *= float64(lenth)
		if pdf[i] < 1.0 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}

	// 构造所有矩形概率值都等于1
	for len(small) != 0 && len(large) != 0 {
		s_index := small[0]
		small = small[1:]
		l_index := large[0]
		large = large[1:]

		probInfo[s_index] = pdf[s_index]
		alias[s_index] = l_index

		// 1.2 -= 1.0- 0.4
		pdf[l_index] -= 1.0 - pdf[s_index]
		if pdf[l_index] < 1.0 {
			small = append(small, l_index)
		} else {
			large = append(large, l_index)
		}
	}

	for len(small) > 0 {
		temp := small[0]
		small = small[1:]
		probInfo[temp] = 1.0
	}

	for len(large) > 0 {
		temp := large[0]
		large = large[1:]
		probInfo[temp] = 1.0
	}

	// 随机取出一个列
	random := rand.New(rand.NewSource(int64(lenth)))
	column := random.Int63n(int64(lenth))

	// 获取一个随机小数
	random = rand.New(rand.NewSource(10000))
	// 获取[10000,0) 随机数
	randomNumber := random.Int63n(10000) / 10000

	if randomNumber < int64(probInfo[column]) {
		return int(column)
	}
	return alias[column]
}
