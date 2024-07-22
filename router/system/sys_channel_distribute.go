package system

import (
	v1 "bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type ChannelDistrubuteRouter struct{}

func (c *ChannelDistrubuteRouter) InitSysChannelDistrubuteRouter(Router *gin.RouterGroup) {
	sysChannelDistrubuteRouter := Router.Group("sysChannelDistribute").Use(middleware.OperationRecord())
	sysChannelDistrubuteRouterWithoutRecord := Router.Group("sysChannelDistribute")
	sysChannelDistrubuteApi := v1.ApiGroupApp.SystemApiGroup.ChannelDistrubuteApi
	{
		sysChannelDistrubuteRouter.POST("createSysChannelDistribute", sysChannelDistrubuteApi.CreateSysChannelDistrubute) // 新建SysChannelDistrubute
		sysChannelDistrubuteRouter.POST("deleteSysChannelDistribute", sysChannelDistrubuteApi.DeleteSysChannelDistrubute) // 删除SysChannelDistrubute
		sysChannelDistrubuteRouter.POST("updateSysChannelDistribute", sysChannelDistrubuteApi.UpdateSysChannelDistrubute) // 更新SysChannelDistrubute
	}
	{
		sysChannelDistrubuteRouterWithoutRecord.GET("findSysChannelDistribute", sysChannelDistrubuteApi.FindSysChannelDistrubute)       // 根据ID获取SysChannelDistrubute
		sysChannelDistrubuteRouterWithoutRecord.GET("getSysChannelDistributeList", sysChannelDistrubuteApi.GetSysChannelDistrubuteList) // 获取SysChannelDistrubute列表
	}
}
