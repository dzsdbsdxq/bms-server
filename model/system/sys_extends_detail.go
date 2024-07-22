// 自动生成模板SysExtendsDetail
package system

import (
	"bms-server/global"
)

// 如果含有time.Time 请自行import time包
type SysExtendsDetail struct {
	global.GVA_MODEL
	Status       *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`                     // 启用状态
	Sort         int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`                           // 排序标记
	SysExtendsID int    `json:"sysExtendsID" form:"sysExtendsID" gorm:"column:sys_extends_id;comment:关联标记"` // 关联标记
	Label        string `json:"label" form:"label" gorm:"column:label;comment:展示值"`                         // 展示值
	Key          string `json:"key" form:"key" gorm:"column:key;comment:字典值"`                               // 字典值
	Value        string `json:"value" form:"value" gorm:"column:value;comment:扩展值"`                         // 扩展值
	AppId        string `json:"appId" form:"appId" gorm:"column:appid;comment:关联标记"`                        // 扩展值
	Type         string `json:"type" form:"type" gorm:"column:type;comment:类型(text=文本，object=对象，array=数组)"` // 扩展值类型
}

func (SysExtendsDetail) TableName() string {
	return "sys_extends_details"
}
