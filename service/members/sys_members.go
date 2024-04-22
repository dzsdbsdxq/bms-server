package members

import (
	"bms-server/global"
	"bms-server/model/members"
	membersReq "bms-server/model/members/request"
	"bms-server/utils"
	"errors"
	"gorm.io/gorm"
)

type SysMembersService struct {
}

// CreateSysMembers 创建sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) CreateSysMembers(sysMembers *members.SysMembers) (err error) {
	var user members.SysMembers
	//创建之前，判断是否用户名是否已经创建了，保证用户名唯一性
	err = global.GVA_DB.Where("username = ?", sysMembers.Username).First(&user).Error
	if err == nil {
		return errors.New("用户名已注册")
	}
	if user.Username != "" && user.Password != "" {
		return errors.New("用户名已注册")
	}
	err = global.GVA_DB.Create(sysMembers).Error
	return err
}

// DeleteSysMembers 删除sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) DeleteSysMembers(ID string) (err error) {
	err = global.GVA_DB.Delete(&members.SysMembers{}, "id = ?", ID).Error
	return err
}

// DeleteSysMembersByIds 批量删除sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) DeleteSysMembersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]members.SysMembers{}, "id in ?", IDs).Error
	return err
}

// UpdateSysMembers 更新sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) UpdateSysMembers(sysMembers members.SysMembers) (err error) {
	err = global.GVA_DB.Save(&sysMembers).Error
	return err
}

// GetSysMembers 根据ID获取sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) GetSysMembers(ID string) (sysMembers members.SysMembers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysMembers).Error
	return
}

// GetSysMembersInfoList 分页获取sysMembers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysMembersService *SysMembersService) GetSysMembersInfoList(info membersReq.SysMembersSearch) (list []members.SysMembers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&members.SysMembers{})
	var sysMemberss []members.SysMembers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysMemberss).Error
	return sysMemberss, total, err
}

// ResetPassword
// @function: resetPassword
// @description: 修改用户密码
// @param: ID uint
// @return: err error
func (sysMembersService *SysMembersService) ResetPassword(ID uint) (err error) {
	err = global.GVA_DB.Model(&members.SysMembers{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

// UpdateSysMemberField
// @function: UpdateSysMemberField
// @description: 对表字段的增加和减少
// @param: ID uint
// @return: err error
func (sysMembersService *SysMembersService) UpdateSysMemberFieldIncDec(ID uint, field string, incr bool, args ...interface{}) (err error) {
	db := global.GVA_DB.Model(&members.SysMembers{}).Where("id = ?", ID)
	if incr {
		err = db.UpdateColumn(field, gorm.Expr(field+" + ?", args)).Error
	} else {
		err = db.UpdateColumn(field, gorm.Expr(field+" - ?", args)).Error
	}
	return err
}
