package initialize

import (
	"os"

	"bms-server/global"
	"bms-server/model/example"
	"bms-server/model/system"

	"bms-server/model/integral"
	"bms-server/model/members"
	"bms-server/model/wheel"
	"bms-server/model/wheelAddress"
	"bms-server/model/wheelPrizes"
	"bms-server/model/wheelRecords"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysExportTemplate{},
		system.Condition{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, members.SysMembers{}, integral.SysIntegral{}, wheel.SysWheels{}, wheelPrizes.SysWheelPrizes{}, wheelAddress.SysWheelAddress{}, wheelRecords.SysWheelRecords{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
