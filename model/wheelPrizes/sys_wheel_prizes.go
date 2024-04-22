// 自动生成模板SysWheelPrizes
package wheelPrizes

import (
	"bms-server/global"
)

// sysWheelPrizes表 结构体  SysWheelPrizes
type SysWheelPrizes struct {
	global.GVA_MODEL
	Name      string   `json:"name" form:"name" gorm:"column:name;comment:奖品名称;size:255;"`                    //奖品名称
	ShortName string   `json:"shortName" form:"shortName" gorm:"column:short_name;comment:简短奖品名称;size:255;"`  //奖品名称
	WheelId   *int     `json:"wheelId" form:"wheelId" gorm:"column:wheel_id;comment:大转盘ID;size:19;"`          //大转盘ID
	Uuid      string   `json:"uuid" form:"uuid" gorm:"column:uuid;comment:唯一UUID;size:191;"`                  //唯一UUID
	ThumbUrl  string   `json:"thumbUrl" form:"thumbUrl" gorm:"column:thumb_url;comment:奖品图片;size:255;"`       //奖品图片
	Nums      *int     `json:"nums" form:"nums" gorm:"column:nums;comment:奖品数量;size:10;"`                     //奖品数量
	UseNums   *int     `json:"useNums" form:"useNums" gorm:"column:use_nums;comment:已抽数量;size:10;"`           //已抽数量
	Perc      *float64 `json:"perc" form:"perc" gorm:"column:perc;comment:奖品概率;size:19;"`                     //奖品概率
	Type      *int     `json:"type" form:"type" gorm:"column:type;comment:是否实物（0=否，1=是）;"`                    //是否实物（0=否，1=是）
	NeedAddr  *int     `json:"needAddr" form:"needAddr" gorm:"column:need_addr;comment:实物是否需要填写地址（0=否，1=是）;"` //实物是否需要填写地址（0=否，1=是）
	IsHit     *int     `json:"isHit" form:"isHit" gorm:"column:is_hit;comment:是否中奖;"`                         //是否中奖
}

// TableName sysWheelPrizes表 SysWheelPrizes自定义表名 sys_wheel_prizes
func (SysWheelPrizes) TableName() string {
	return "sys_wheel_prizes"
}
