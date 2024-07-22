package system

import (
	"bms-server/model/common/request"
	"errors"

	"bms-server/global"
	"bms-server/model/system"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysChannelDistrubute
//@description: 壹深圳分发下载
//@param: sysChannelDistrubute model.SysChannelDistrubute
//@return: err error

type ChannelDistrubuteService struct{}

func (channelDistrubuteService *ChannelDistrubuteService) CreateSysChannelDistrubute(sysChannelDistrubute system.SysChannelDistrubute) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysExtends{}, "appid = ?", sysChannelDistrubute.AppId).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysChannelDistrubute).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysChannelDistrubute
//@description: 删除
//@param: sysChannelDistrubute model.SysChannelDistrubute
//@return: err error

func (channelDistrubuteService *ChannelDistrubuteService) DeleteSysChannelDistrubute(sysChannelDistrubute system.SysChannelDistrubute) (err error) {
	err = global.GVA_DB.Where("id = ?", sysChannelDistrubute.ID).First(&sysChannelDistrubute).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysChannelDistrubute).Error
	if err != nil {
		return err
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysChannelDistrubute
//@description: 更新
//@param: sysChannelDistrubute *model.SysChannelDistrubute
//@return: err error

func (channelDistrubuteService *ChannelDistrubuteService) UpdateSysChannelDistrubute(sysChannelDistrubute *system.SysChannelDistrubute) (err error) {
	sysChannelDistrubuteMap := map[string]interface{}{
		"Name":      sysChannelDistrubute.Name,
		"AppId":     sysChannelDistrubute.AppId,
		"Status":    sysChannelDistrubute.Status,
		"JumpType":  sysChannelDistrubute.JumpType,
		"JumpValue": sysChannelDistrubute.JumpValue,
	}
	err = global.GVA_DB.Model(&system.SysChannelDistrubute{}).Where("id = ?", sysChannelDistrubute.ID).Updates(sysChannelDistrubuteMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysChannelDistrubute
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysChannelDistrubute model.SysChannelDistrubute

func (channelDistrubuteService *ChannelDistrubuteService) GetSysChannelDistrubute(Id uint) (sysChannelDistrubute system.SysChannelDistrubute, err error) {
	err = global.GVA_DB.Where("id = ?", Id).First(&sysChannelDistrubute).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysExtendsInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (channelDistrubuteService *ChannelDistrubuteService) GetSysChannelDistrubuteInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	var sysChannelDistrubute []system.SysChannelDistrubute

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GVA_DB.Model(&system.SysChannelDistrubute{})
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&sysChannelDistrubute).Error
	return sysChannelDistrubute, total, err
}
