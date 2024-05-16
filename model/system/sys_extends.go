// 自动生成模板SysExtends
package system

import (
	"bms-server/global"
)

// 如果含有time.Time 请自行import time包
type SysExtends struct {
	global.GVA_MODEL
	Name              string             `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Type              string             `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`   // 字典名（英）
	Status            *bool              `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	Desc              string             `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`       // 描述
	AppId             string             `json:"appid" form:"appid" gorm:"column:appid;comment:APPID"` // APPID
	SysExtendsDetails []SysExtendsDetail `json:"sysExtendsDetails" form:"sysExtendsDetails"`
}

func (SysExtends) TableName() string {
	return "sys_extends"
}
