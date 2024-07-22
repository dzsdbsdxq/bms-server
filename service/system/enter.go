package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	ExtendsService
	ExtendsDetailService
	AuthorityBtnService
	SysExportTemplateService
	SnsRecordService
	ChannelDistrubuteService
}
