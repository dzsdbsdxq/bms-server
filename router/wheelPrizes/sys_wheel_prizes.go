package wheelPrizes

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysWheelPrizesRouter struct {
}

// InitSysWheelPrizesRouter 初始化 sysWheelPrizes表 路由信息
func (s *SysWheelPrizesRouter) InitSysWheelPrizesRouter(Router *gin.RouterGroup) {
	sysWheelPrizesRouter := Router.Group("sysWheelPrizes").Use(middleware.OperationRecord())
	sysWheelPrizesRouterWithoutRecord := Router.Group("sysWheelPrizes")
	var sysWheelPrizesApi = v1.ApiGroupApp.WheelPrizesApiGroup.SysWheelPrizesApi
	{
		sysWheelPrizesRouter.POST("createSysWheelPrizes", sysWheelPrizesApi.CreateSysWheelPrizes)   // 新建sysWheelPrizes表
		sysWheelPrizesRouter.DELETE("deleteSysWheelPrizes", sysWheelPrizesApi.DeleteSysWheelPrizes) // 删除sysWheelPrizes表
		sysWheelPrizesRouter.DELETE("deleteSysWheelPrizesByIds", sysWheelPrizesApi.DeleteSysWheelPrizesByIds) // 批量删除sysWheelPrizes表
		sysWheelPrizesRouter.PUT("updateSysWheelPrizes", sysWheelPrizesApi.UpdateSysWheelPrizes)    // 更新sysWheelPrizes表
	}
	{
		sysWheelPrizesRouterWithoutRecord.GET("findSysWheelPrizes", sysWheelPrizesApi.FindSysWheelPrizes)        // 根据ID获取sysWheelPrizes表
		sysWheelPrizesRouterWithoutRecord.GET("getSysWheelPrizesList", sysWheelPrizesApi.GetSysWheelPrizesList)  // 获取sysWheelPrizes表列表
	}
}
