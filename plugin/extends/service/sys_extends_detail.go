package service

import (
	"bms-server/global"
	"bms-server/model/system"
)

type ExtendsDetailService struct{}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (e *ExtendsDetailService) ExtendsDetailInfo(appId string) (list []system.SysExtendsDetail, err error) {
	// 创建db
	db := global.GVA_DB.Model(&system.SysExtendsDetail{})
	var sysExtendsDetails []system.SysExtendsDetail

	db = db.Where("status = ? AND appid = ?", 1, appId)

	err = db.Select("id,label,`key`,`value`").Scan(&sysExtendsDetails).Error
	return sysExtendsDetails, err
}
