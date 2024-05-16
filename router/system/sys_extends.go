package system

import (
	v1 "bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type ExtendsRouter struct{}

func (s *ExtendsRouter) InitSysExtendsRouter(Router *gin.RouterGroup) {
	sysExtendsRouter := Router.Group("sysExtends").Use(middleware.OperationRecord())
	sysExtendsRouterWithoutRecord := Router.Group("sysExtends")
	sysExtendsApi := v1.ApiGroupApp.SystemApiGroup.ExtendsApi
	{
		sysExtendsRouter.POST("createSysExtends", sysExtendsApi.CreateSysExtends)   // 新建SysExtends
		sysExtendsRouter.DELETE("deleteSysExtends", sysExtendsApi.DeleteSysExtends) // 删除SysExtends
		sysExtendsRouter.PUT("updateSysExtends", sysExtendsApi.UpdateSysExtends)    // 更新SysExtends
	}
	{
		sysExtendsRouterWithoutRecord.GET("findSysExtends", sysExtendsApi.FindSysExtends)       // 根据ID获取SysExtends
		sysExtendsRouterWithoutRecord.GET("getSysExtendsList", sysExtendsApi.GetSysExtendsList) // 获取SysExtends列表
	}
}
