// 自动生成模板SysExtends
package system

import (
	"bms-server/global"
)

// 如果含有time.Time 请自行import time包
type SysChannelDistrubute struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称"`
	Status    *bool  `json:"status" form:"status" gorm:"column:status;comment:状态"`
	Views     uint   `json:"views" form:"views" gorm:"column:views;comment:查看数"`
	Down      uint   `json:"down" form:"down" gorm:"column:down;comment:下载数"`
	JumpType  string `json:"jump_type" form:"jump_type" gorm:"column:jump_type;comment:跳转类型"`
	JumpValue string `json:"jump_value" form:"jump_value" gorm:"column:jump_value;comment:跳转值"`
	AppId     string `json:"appid" form:"appid" gorm:"column:appid;comment:分发ID"`
	DType     string `json:"d_type" form:"d_type" gorm:"column:d_type;comment:类型(text=文本，object=对象，array=数组)"` // 扩展值类型
}

func (SysChannelDistrubute) TableName() string {
	return "sys_channel_distribute"
}
