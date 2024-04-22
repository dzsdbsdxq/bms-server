package wheel

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysWheelsRouter struct {
}

// InitSysWheelsRouter 初始化 sysWheels表 路由信息
func (s *SysWheelsRouter) InitSysWheelsRouter(Router *gin.RouterGroup) {
	sysWheelsRouter := Router.Group("sysWheels").Use(middleware.OperationRecord())
	sysWheelsRouterWithoutRecord := Router.Group("sysWheels")
	var sysWheelsApi = v1.ApiGroupApp.WheelApiGroup.SysWheelsApi
	{
		sysWheelsRouter.POST("createSysWheels", sysWheelsApi.CreateSysWheels)             // 新建sysWheels表
		sysWheelsRouter.DELETE("deleteSysWheels", sysWheelsApi.DeleteSysWheels)           // 删除sysWheels表
		sysWheelsRouter.DELETE("deleteSysWheelsByIds", sysWheelsApi.DeleteSysWheelsByIds) // 批量删除sysWheels表
		sysWheelsRouter.PUT("updateSysWheels", sysWheelsApi.UpdateSysWheels)              // 更新sysWheels表
		sysWheelsRouter.PUT("updatePageParams", sysWheelsApi.UpdatePageParams)            // 更新sysWheels表
	}
	{
		sysWheelsRouterWithoutRecord.GET("getTemplateInfo", sysWheelsApi.GetSysWheelsTemplateInfo)     // 获取模版配置信息
		sysWheelsRouterWithoutRecord.GET("getTemplateConfig", sysWheelsApi.GetSysWheelsTemplateConfig) // 获取模版配置信息
		sysWheelsRouterWithoutRecord.GET("findSysWheels", sysWheelsApi.FindSysWheels)                  // 根据ID获取sysWheels表
		sysWheelsRouterWithoutRecord.GET("getSysWheelsList", sysWheelsApi.GetSysWheelsList)            // 获取sysWheels表列表
	}
}
