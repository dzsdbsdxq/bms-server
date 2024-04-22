package wheel

import (
	"bms-server/global"
	"bms-server/model/wheel"
	wheelReq "bms-server/model/wheel/request"
)

type SysWheelsService struct {
}

// CreateSysWheels 创建sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) CreateSysWheels(sysWheels *wheel.SysWheels) (err error) {
	err = global.GVA_DB.Create(sysWheels).Error
	return err
}

// DeleteSysWheels 删除sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) DeleteSysWheels(ID string) (err error) {
	err = global.GVA_DB.Delete(&wheel.SysWheels{}, "id = ?", ID).Error
	return err
}

// DeleteSysWheelsByIds 批量删除sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) DeleteSysWheelsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]wheel.SysWheels{}, "id in ?", IDs).Error
	return err
}

// UpdateSysWheels 更新sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) UpdateSysWheels(sysWheels wheel.SysWheels) (err error) {
	err = global.GVA_DB.Save(&sysWheels).Error
	return err
}

// UpdatePageParams 更新PageParams记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) UpdatePageParams(id uint, pageParams string) (err error) {
	err = global.GVA_DB.Model(&wheel.SysWheels{}).Where("id = ?", id).Update("page_params", pageParams).Error
	return err
}

// GetSysWheels 根据ID获取sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) GetSysWheels(ID string) (sysWheels wheel.SysWheels, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysWheels).Error
	return
}

// GetSysWheels 根据UUID获取sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) GetSysWheelsByUUID(UUID string) (sysWheels wheel.SysWheels, err error) {
	err = global.GVA_DB.Where("uuid = ?", UUID).First(&sysWheels).Error
	return
}

// GetSysWheelsInfoList 分页获取sysWheels表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelsService *SysWheelsService) GetSysWheelsInfoList(info wheelReq.SysWheelsSearch) (list []wheel.SysWheels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wheel.SysWheels{})
	var sysWheelss []wheel.SysWheels
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

	err = db.Find(&sysWheelss).Error
	return sysWheelss, total, err
}
