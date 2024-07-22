package system

import (
	"bms-server/model/common/request"
	"errors"

	"bms-server/global"
	"bms-server/model/system"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type ExtendsService struct{}

func (extendsService *ExtendsService) CreateSysExtends(sysExtends system.SysExtends) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysExtends{}, "type = ?", sysExtends.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysExtends).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysExtends
//@description: 删除字典数据
//@param: sysExtends model.SysExtends
//@return: err error

func (extendsService *ExtendsService) DeleteSysExtends(sysExtends system.SysExtends) (err error) {
	err = global.GVA_DB.Where("id = ?", sysExtends.ID).Preload("SysExtendsDetails").First(&sysExtends).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysExtends).Error
	if err != nil {
		return err
	}

	if sysExtends.SysExtendsDetails != nil {
		return global.GVA_DB.Where("sys_extends_id=?", sysExtends.ID).Delete(sysExtends.SysExtendsDetails).Error
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysExtends
//@description: 更新字典数据
//@param: sysExtends *model.SysExtends
//@return: err error

func (extendsService *ExtendsService) UpdateSysExtends(sysExtends *system.SysExtends) (err error) {
	sysExtendsMap := map[string]interface{}{
		"Name":   sysExtends.Name,
		"Type":   sysExtends.Type,
		"Status": sysExtends.Status,
		"Desc":   sysExtends.Desc,
		"DType":  sysExtends.DType,
	}
	err = global.GVA_DB.Model(&system.SysExtends{}).Where("id = ?", sysExtends.ID).Updates(sysExtendsMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysExtends
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysExtends model.SysExtends

func (extendsService *ExtendsService) GetSysExtends(Id uint) (sysExtends system.SysExtends, err error) {
	err = global.GVA_DB.Where("id = ?", Id).First(&sysExtends).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysExtendsInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (extendsService *ExtendsService) GetSysExtendsInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	var sysExtends []system.SysExtends

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GVA_DB.Model(&system.SysExtends{})
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&sysExtends).Error
	return sysExtends, total, err
}
