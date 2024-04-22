package integral

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysIntegralRouter struct {
}

// InitSysIntegralRouter 初始化 sysIntegral表 路由信息
func (s *SysIntegralRouter) InitSysIntegralRouter(Router *gin.RouterGroup) {
	sysIntegralRouter := Router.Group("sysIntegral").Use(middleware.OperationRecord())
	sysIntegralRouterWithoutRecord := Router.Group("sysIntegral")
	var sysIntegralApi = v1.ApiGroupApp.IntegralApiGroup.SysIntegralApi
	{
		sysIntegralRouter.POST("createSysIntegral", sysIntegralApi.CreateSysIntegral)   // 新建sysIntegral表
		sysIntegralRouter.DELETE("deleteSysIntegral", sysIntegralApi.DeleteSysIntegral) // 删除sysIntegral表
		sysIntegralRouter.DELETE("deleteSysIntegralByIds", sysIntegralApi.DeleteSysIntegralByIds) // 批量删除sysIntegral表
		sysIntegralRouter.PUT("updateSysIntegral", sysIntegralApi.UpdateSysIntegral)    // 更新sysIntegral表
	}
	{
		sysIntegralRouterWithoutRecord.GET("findSysIntegral", sysIntegralApi.FindSysIntegral)        // 根据ID获取sysIntegral表
		sysIntegralRouterWithoutRecord.GET("getSysIntegralList", sysIntegralApi.GetSysIntegralList)  // 获取sysIntegral表列表
	}
}
