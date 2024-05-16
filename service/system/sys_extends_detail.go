package system

import (
	"bms-server/global"
	"bms-server/model/system"
	"bms-server/model/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysExtendsDetail
//@description: 创建字典详情数据
//@param: sysExtendsDetail model.SysExtendsDetail
//@return: err error

type ExtendsDetailService struct{}

func (extendsDetailService *ExtendsDetailService) CreateSysExtendsDetail(sysExtendsDetail system.SysExtendsDetail) (err error) {
	err = global.GVA_DB.Create(&sysExtendsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysExtendsDetail
//@description: 删除字典详情数据
//@param: sysExtendsDetail model.SysExtendsDetail
//@return: err error

func (extendsDetailService *ExtendsDetailService) DeleteSysExtendsDetail(sysExtendsDetail system.SysExtendsDetail) (err error) {
	err = global.GVA_DB.Delete(&sysExtendsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysExtendsDetail
//@description: 更新字典详情数据
//@param: sysExtendsDetail *model.SysExtendsDetail
//@return: err error

func (extendsDetailService *ExtendsDetailService) UpdateSysExtendsDetail(sysExtendsDetail *system.SysExtendsDetail) (err error) {
	err = global.GVA_DB.Save(sysExtendsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysExtendsDetail
//@description: 根据id获取字典详情单条数据
//@param: id uint
//@return: sysExtendsDetail system.SysExtendsDetail, err error

func (extendsDetailService *ExtendsDetailService) GetSysExtendsDetail(id uint) (sysExtendsDetail system.SysExtendsDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysExtendsDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: list interface{}, total int64, err error

func (extendsDetailService *ExtendsDetailService) GetSysExtendsDetailInfoList(info request.SysExtendsDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysExtendsDetail{})
	var sysExtendsDetails []system.SysExtendsDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysExtendsID != 0 {
		db = db.Where("sys_extends_id = ?", info.SysExtendsID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysExtendsDetails).Error
	return sysExtendsDetails, total, err
}

// 按照字典id获取字典全部内容的方法
func (extendsDetailService *ExtendsDetailService) GetExtendsList(extendsID uint) (list []system.SysExtendsDetail, err error) {
	var sysExtendsDetails []system.SysExtendsDetail
	err = global.GVA_DB.Find(&sysExtendsDetails, "sys_extends_id = ?", extendsID).Error
	return sysExtendsDetails, err
}

// 按照字典type获取字典全部内容的方法
func (extendsDetailService *ExtendsDetailService) GetExtendsListByType(t string) (list []system.SysExtendsDetail, err error) {
	var sysExtendsDetails []system.SysExtendsDetail
	db := global.GVA_DB.Model(&system.SysExtendsDetail{}).Joins("JOIN sys_extends ON sys_extends.id = sys_extends_details.sys_extends_id")
	err = db.Debug().Find(&sysExtendsDetails, "type = ?", t).Error
	return sysExtendsDetails, err
}

// 按照字典id+字典内容value获取单条字典内容
func (extendsDetailService *ExtendsDetailService) GetExtendsInfoByValue(extendsID uint, value uint) (detail system.SysExtendsDetail, err error) {
	var sysExtendsDetail system.SysExtendsDetail
	err = global.GVA_DB.First(&sysExtendsDetail, "sys_extends_id = ? and value = ?", extendsID, value).Error
	return sysExtendsDetail, err
}

// 按照字典type+字典内容value获取单条字典内容
func (extendsDetailService *ExtendsDetailService) GetExtendsInfoByTypeValue(t string, value uint) (detail system.SysExtendsDetail, err error) {
	var sysExtendsDetails system.SysExtendsDetail
	db := global.GVA_DB.Model(&system.SysExtendsDetail{}).Joins("JOIN sys_extends ON sys_extends.id = sys_extends_details.sys_extends_id")
	err = db.First(&sysExtendsDetails, "sys_extends.type = ? and sys_extends_details.value = ?", t, value).Error
	return sysExtendsDetails, err
}
