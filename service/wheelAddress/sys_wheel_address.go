package wheelAddress

import (
	"bms-server/global"
	"bms-server/model/wheelAddress"
    wheelAddressReq "bms-server/model/wheelAddress/request"
)

type SysWheelAddressService struct {
}

// CreateSysWheelAddress 创建sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService) CreateSysWheelAddress(sysWheelAddress *wheelAddress.SysWheelAddress) (err error) {
	err = global.GVA_DB.Create(sysWheelAddress).Error
	return err
}

// DeleteSysWheelAddress 删除sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService)DeleteSysWheelAddress(ID string) (err error) {
	err = global.GVA_DB.Delete(&wheelAddress.SysWheelAddress{},"id = ?",ID).Error
	return err
}

// DeleteSysWheelAddressByIds 批量删除sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService)DeleteSysWheelAddressByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]wheelAddress.SysWheelAddress{},"id in ?",IDs).Error
	return err
}

// UpdateSysWheelAddress 更新sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService)UpdateSysWheelAddress(sysWheelAddress wheelAddress.SysWheelAddress) (err error) {
	err = global.GVA_DB.Save(&sysWheelAddress).Error
	return err
}

// GetSysWheelAddress 根据ID获取sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService)GetSysWheelAddress(ID string) (sysWheelAddress wheelAddress.SysWheelAddress, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysWheelAddress).Error
	return
}

// GetSysWheelAddressInfoList 分页获取sysWheelAddress表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelAddressService *SysWheelAddressService)GetSysWheelAddressInfoList(info wheelAddressReq.SysWheelAddressSearch) (list []wheelAddress.SysWheelAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wheelAddress.SysWheelAddress{})
    var sysWheelAddresss []wheelAddress.SysWheelAddress
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&sysWheelAddresss).Error
	return  sysWheelAddresss, total, err
}
