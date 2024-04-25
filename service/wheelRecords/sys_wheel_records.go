package wheelRecords

import (
	"bms-server/global"
	"bms-server/model/wheelRecords"
	wheelRecordsReq "bms-server/model/wheelRecords/request"
)

type SysWheelRecordsService struct {
}

// CreateSysWheelRecords 创建sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) CreateSysWheelRecords(sysWheelRecords *wheelRecords.SysWheelRecords) (err error) {
	err = global.GVA_DB.Create(sysWheelRecords).Error
	return err
}

// DeleteSysWheelRecords 删除sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) DeleteSysWheelRecords(ID string) (err error) {
	err = global.GVA_DB.Delete(&wheelRecords.SysWheelRecords{}, "id = ?", ID).Error
	return err
}

// DeleteSysWheelRecordsByIds 批量删除sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) DeleteSysWheelRecordsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]wheelRecords.SysWheelRecords{}, "id in ?", IDs).Error
	return err
}

// UpdateSysWheelRecords 更新sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) UpdateSysWheelRecords(sysWheelRecords wheelRecords.SysWheelRecords) (err error) {
	err = global.GVA_DB.Save(&sysWheelRecords).Error
	return err
}

// GetSysWheelRecords 根据ID获取sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) GetSysWheelRecords(ID string) (sysWheelRecords wheelRecords.SysWheelRecords, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysWheelRecords).Error
	return
}

// GetSysWheelRecordsInfoList 分页获取sysWheelRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelRecordsService *SysWheelRecordsService) GetSysWheelRecordsInfoList(info wheelRecordsReq.SysWheelRecordsSearch) (list []wheelRecords.SysWheelRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wheelRecords.SysWheelRecords{})
	var sysWheelRecordss []wheelRecords.SysWheelRecords
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

	err = db.Find(&sysWheelRecordss).Error
	return sysWheelRecordss, total, err
}

func (sysWheelRecordsService *SysWheelRecordsService) GetSysWheelRecordsInfoListBy(wheelId int, memberId int) []wheelRecords.SysWheelRecords {
	// 创建db
	db := global.GVA_DB.Model(&wheelRecords.SysWheelRecords{})
	var sysWheelRecordss []wheelRecords.SysWheelRecords
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("wheel_id= ? AND member_id=?", wheelId, memberId)
	err := db.Find(&sysWheelRecordss).Error
	if err != nil {
		return nil
	}
	return sysWheelRecordss
}
