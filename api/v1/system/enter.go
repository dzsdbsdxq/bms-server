package system

import "bms-server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	AutoCodeHistoryApi
	DictionaryDetailApi
	AuthorityBtnApi
	SysExportTemplateApi
	ExtendsApi
	ExtendsDetailApi
	SnsRecordApi
	ChannelDistrubuteApi
}

var (
	apiService               = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService               = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService              = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService              = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService            = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService            = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	autoCodeService          = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	baseMenuService          = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService         = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService        = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	extendsService           = service.ServiceGroupApp.SystemServiceGroup.ExtendsService
	extendsDetailService     = service.ServiceGroupApp.SystemServiceGroup.ExtendsDetailService
	systemConfigService      = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService   = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	snsRecordService         = service.ServiceGroupApp.SystemServiceGroup.SnsRecordService
	autoCodeHistoryService   = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService  = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	authorityBtnService      = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
	channelDistrubuteService = service.ServiceGroupApp.SystemServiceGroup.ChannelDistrubuteService
)
