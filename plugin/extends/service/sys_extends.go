package service

import (
	"bms-server/global"
	"bms-server/model/system"
)

type ExtendsService struct{}

func (e *ExtendsService) GetSysExtends(appId string) (sysExtends system.SysExtends, err error) {
	err = global.GVA_DB.Where("status = 1 AND appid = ?", appId).First(&sysExtends).Error
	return
}
