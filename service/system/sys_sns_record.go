package system

import (
	"bms-server/global"
	"bms-server/model/common/request"
	"bms-server/model/system"
	systemReq "bms-server/model/system/request"
)

type SnsRecordService struct{}

func (snsRecordService *SnsRecordService) CreateSysSnsRecord(sysSnsRecord system.SysSnsRecord) (err error) {
	err = global.GVA_DB.Create(&sysSnsRecord).Error
	return err
}

func (snsRecordService *SnsRecordService) DeleteSysSnsRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysSnsRecord{}, "id in (?)", ids.Ids).Error
	return err
}

func (snsRecordService *SnsRecordService) DeleteSysSnsRecord(sysSnsRecord system.SysSnsRecord) (err error) {
	err = global.GVA_DB.Delete(&sysSnsRecord).Error
	return err
}

func (snsRecordService *SnsRecordService) GetSysSnsRecord(id uint) (sysSnsRecord system.SysSnsRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysSnsRecord).Error
	return
}

func (snsRecordService *SnsRecordService) GetSysSnsRecordInfoList(info systemReq.SysSnsRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysSnsRecord{})
	var sysSnsRecords []system.SysSnsRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Mobile != "" {
		db = db.Where("mobile = ?", info.Mobile)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&sysSnsRecords).Error
	return sysSnsRecords, total, err
}
