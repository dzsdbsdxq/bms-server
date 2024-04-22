package integral

import (
	"bms-server/global"
	"bms-server/model/integral"
	integralReq "bms-server/model/integral/request"
	"gorm.io/gorm"
)

type SysIntegralService struct {
}

// CreateSysIntegral 创建sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) CreateSysIntegral(sysIntegral *integral.SysIntegral) (err error) {
	err = global.GVA_DB.Create(sysIntegral).Error
	return err
}

// DeleteSysIntegral 删除sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) DeleteSysIntegral(ID string) (err error) {
	err = global.GVA_DB.Delete(&integral.SysIntegral{}, "id = ?", ID).Error
	return err
}

// DeleteSysIntegralByIds 批量删除sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) DeleteSysIntegralByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]integral.SysIntegral{}, "id in ?", IDs).Error
	return err
}

// UpdateSysIntegral 更新sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) UpdateSysIntegral(sysIntegral integral.SysIntegral) (err error) {
	err = global.GVA_DB.Save(&sysIntegral).Error
	return err
}

// GetSysIntegral 根据ID获取sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) GetSysIntegral(ID string) (sysIntegral integral.SysIntegral, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysIntegral).Error
	return
}

// GetSysIntegralInfoList 分页获取sysIntegral表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysIntegralService *SysIntegralService) GetSysIntegralInfoList(info integralReq.SysIntegralSearch) (list []integral.SysIntegral, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&integral.SysIntegral{})
	var sysIntegrals []integral.SysIntegral
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

	err = db.Preload("SysMembers", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,username,nick_name")
	}).Find(&sysIntegrals).Error

	return sysIntegrals, total, err
}
