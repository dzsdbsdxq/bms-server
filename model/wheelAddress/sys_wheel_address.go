// 自动生成模板SysWheelAddress
package wheelAddress

import (
	"bms-server/global"
	
	
)

// sysWheelAddress表 结构体  SysWheelAddress
type SysWheelAddress struct {
 global.GVA_MODEL 
      MemberId  *int `json:"memberId" form:"memberId" gorm:"column:member_id;comment:用户ID;size:19;"`  //用户ID 
      WheelId  *int `json:"wheelId" form:"wheelId" gorm:"column:wheel_id;comment:大转盘ID;size:19;"`  //大转盘ID 
      WheelRecordsId  *int `json:"wheelRecordsId" form:"wheelRecordsId" gorm:"column:wheel_records_id;comment:抽奖记录ID;size:19;"`  //抽奖记录ID 
      WheelPrizeId  *int `json:"wheelPrizeId" form:"wheelPrizeId" gorm:"column:wheel_prize_id;comment:奖项ID;size:19;"`  //奖项ID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:用户名;size:255;"`  //用户名 
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:用户电话;"`  //用户电话 
      Address  string `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;"`  //用户地址 
}


// TableName sysWheelAddress表 SysWheelAddress自定义表名 sys_wheel_address
func (SysWheelAddress) TableName() string {
  return "sys_wheel_address"
}

