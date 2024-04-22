package wheelPrizes

import (
	"bms-server/global"
	"bms-server/model/wheelPrizes"
	wheelPrizesReq "bms-server/model/wheelPrizes/request"
)

type SysWheelPrizesService struct {
}

// CreateSysWheelPrizes 创建sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) CreateSysWheelPrizes(sysWheelPrizes *wheelPrizes.SysWheelPrizes) (err error) {
	err = global.GVA_DB.Create(sysWheelPrizes).Error
	return err
}

// DeleteSysWheelPrizes 删除sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) DeleteSysWheelPrizes(ID string) (err error) {
	err = global.GVA_DB.Delete(&wheelPrizes.SysWheelPrizes{}, "id = ?", ID).Error
	return err
}

// DeleteSysWheelPrizesByIds 批量删除sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) DeleteSysWheelPrizesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]wheelPrizes.SysWheelPrizes{}, "id in ?", IDs).Error
	return err
}

// UpdateSysWheelPrizes 更新sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) UpdateSysWheelPrizes(sysWheelPrizes wheelPrizes.SysWheelPrizes) (err error) {
	err = global.GVA_DB.Save(&sysWheelPrizes).Error
	return err
}

// GetSysWheelPrizes 根据ID获取sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) GetSysWheelPrizes(ID string) (sysWheelPrizes wheelPrizes.SysWheelPrizes, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysWheelPrizes).Error
	return
}

// GetSysWheelPrizesInfoList 分页获取sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) GetSysWheelPrizesInfoList(info wheelPrizesReq.SysWheelPrizesSearch) (list []wheelPrizes.SysWheelPrizes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wheelPrizes.SysWheelPrizes{})
	var sysWheelPrizess []wheelPrizes.SysWheelPrizes
	db = db.Where("wheel_id = ?", info.ID)
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysWheelPrizess).Error
	return sysWheelPrizess, total, err
}
