package wheelPrizes

import (
	"bms-server/global"
	"bms-server/model/wheelPrizes"
	wheelPrizesReq "bms-server/model/wheelPrizes/request"
	"bms-server/service/wheel"
	"gorm.io/gorm"
	"sync"
)

type SysWheelPrizesService struct {
}

// CreateSysWheelPrizes 创建sysWheelPrizes表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) CreateSysWheelPrizes(sysWheelPrizes *wheelPrizes.SysWheelPrizes) (err error) {
	err = global.GVA_DB.Create(sysWheelPrizes).Error
	return err
}

func (sysWheelPrizesService *SysWheelPrizesService) InitWheelPrizesPool(info wheelPrizesReq.SysWheelPrizesSearch, st uint, end uint) error {
	list, _, err := sysWheelPrizesService.GetSysWheelPrizesInfoList(info)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, aw := range list {
		aw := wheel.AwardBatchService{
			WheelId:      *aw.WheelId,
			PrizeId:      int(aw.ID),
			PrizeType:    *aw.Type,
			IsHit:        *aw.IsHit,
			Probability:  *aw.Perc,
			Uuid:         aw.Uuid,
			PrizeName:    aw.Name,
			TotalBalance: int64(*aw.Nums),
			TotalAmount:  int64(*aw.Nums),
			UpdateTime:   0,
			StartTime:    int64(st),
			EndTime:      int64(end),
		}
		wg.Add(1)
		go func(aw wheel.AwardBatchService) {
			wheel.InitAwardPool(&aw)
			wg.Done()
		}(aw)
	}
	wg.Wait()
	return nil
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

// UpdateSysWheelPrizesField 更新sysWheelPrizes表某字段值
// Author [piexlmax](https://github.com/piexlmax)
func (sysWheelPrizesService *SysWheelPrizesService) UpdateSysWheelPrizesField(field string, ID string) error {
	err := global.GVA_DB.Where("id = ?", ID).UpdateColumn(field, gorm.Expr(field+" + ?", 1)).Error
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
