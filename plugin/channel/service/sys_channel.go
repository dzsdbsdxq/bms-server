package service

import (
	"bms-server/global"
	"bms-server/model/system"
	"gorm.io/gorm"
)

type ChannelService struct{}

func (e *ChannelService) GetSysChannel(appId string) (sysChannel system.SysChannelDistrubute, err error) {
	err = global.GVA_DB.Where("status = 1 AND appid = ?", appId).First(&sysChannel).Error
	return
}
func (e *ChannelService) UpdateSysChannelField(field string, appId string) error {
	err := global.GVA_DB.Model(&system.SysChannelDistrubute{}).Where("appid = ?", appId).UpdateColumn(field, gorm.Expr(field+" + ?", 1)).Error
	return err
}
