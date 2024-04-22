package wheelAddress

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysWheelAddressRouter struct {
}

// InitSysWheelAddressRouter 初始化 sysWheelAddress表 路由信息
func (s *SysWheelAddressRouter) InitSysWheelAddressRouter(Router *gin.RouterGroup) {
	sysWheelAddressRouter := Router.Group("sysWheelAddress").Use(middleware.OperationRecord())
	sysWheelAddressRouterWithoutRecord := Router.Group("sysWheelAddress")
	var sysWheelAddressApi = v1.ApiGroupApp.WheelAddressApiGroup.SysWheelAddressApi
	{
		sysWheelAddressRouter.POST("createSysWheelAddress", sysWheelAddressApi.CreateSysWheelAddress)   // 新建sysWheelAddress表
		sysWheelAddressRouter.DELETE("deleteSysWheelAddress", sysWheelAddressApi.DeleteSysWheelAddress) // 删除sysWheelAddress表
		sysWheelAddressRouter.DELETE("deleteSysWheelAddressByIds", sysWheelAddressApi.DeleteSysWheelAddressByIds) // 批量删除sysWheelAddress表
		sysWheelAddressRouter.PUT("updateSysWheelAddress", sysWheelAddressApi.UpdateSysWheelAddress)    // 更新sysWheelAddress表
	}
	{
		sysWheelAddressRouterWithoutRecord.GET("findSysWheelAddress", sysWheelAddressApi.FindSysWheelAddress)        // 根据ID获取sysWheelAddress表
		sysWheelAddressRouterWithoutRecord.GET("getSysWheelAddressList", sysWheelAddressApi.GetSysWheelAddressList)  // 获取sysWheelAddress表列表
	}
}
