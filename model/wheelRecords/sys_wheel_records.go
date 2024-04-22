// 自动生成模板SysWheelRecords
package wheelRecords

import (
	"bms-server/global"
	"time"
	
)

// sysWheelRecords表 结构体  SysWheelRecords
type SysWheelRecords struct {
 global.GVA_MODEL 
      MemberId  *int `json:"memberId" form:"memberId" gorm:"column:member_id;comment:用户ID;size:19;"`  //用户ID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:冗余字段-奖品名称;size:255;"`  //冗余字段-奖品名称 
      WheelId  *int `json:"wheelId" form:"wheelId" gorm:"column:wheel_id;comment:大转盘ID;size:19;"`  //大转盘ID 
      WheelPrizeId  *int `json:"wheelPrizeId" form:"wheelPrizeId" gorm:"column:wheel_prize_id;comment:奖项ID;size:19;"`  //奖项ID 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:是否实物（0=否，1=是）;"`  //是否实物（0=否，1=是） 
      IsHit  *int `json:"isHit" form:"isHit" gorm:"column:is_hit;comment:是否中奖;"`  //是否中奖 
      Exchange  *int `json:"exchange" form:"exchange" gorm:"column:exchange;comment:是否兑换（核销）;"`  //是否兑换（核销） 
      ExchangeAt  *time.Time `json:"exchangeAt" form:"exchangeAt" gorm:"column:exchange_at;comment:兑换时间;"`  //兑换时间 
      Delivery  *int `json:"delivery" form:"delivery" gorm:"column:delivery;comment:是否发货;"`  //是否发货 
      ExpressId  string `json:"expressId" form:"expressId" gorm:"column:express_id;comment:物流单号;size:20;"`  //物流单号 
      ExpressAt  *time.Time `json:"expressAt" form:"expressAt" gorm:"column:express_at;comment:发货时间;"`  //发货时间 
}


// TableName sysWheelRecords表 SysWheelRecords自定义表名 sys_wheel_records
func (SysWheelRecords) TableName() string {
  return "sys_wheel_records"
}

