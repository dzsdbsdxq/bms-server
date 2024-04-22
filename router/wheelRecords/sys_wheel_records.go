package wheelRecords

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysWheelRecordsRouter struct {
}

// InitSysWheelRecordsRouter 初始化 sysWheelRecords表 路由信息
func (s *SysWheelRecordsRouter) InitSysWheelRecordsRouter(Router *gin.RouterGroup) {
	sysWheelRecordsRouter := Router.Group("sysWheelRecords").Use(middleware.OperationRecord())
	sysWheelRecordsRouterWithoutRecord := Router.Group("sysWheelRecords")
	var sysWheelRecordsApi = v1.ApiGroupApp.WheelRecordsApiGroup.SysWheelRecordsApi
	{
		sysWheelRecordsRouter.POST("createSysWheelRecords", sysWheelRecordsApi.CreateSysWheelRecords)   // 新建sysWheelRecords表
		sysWheelRecordsRouter.DELETE("deleteSysWheelRecords", sysWheelRecordsApi.DeleteSysWheelRecords) // 删除sysWheelRecords表
		sysWheelRecordsRouter.DELETE("deleteSysWheelRecordsByIds", sysWheelRecordsApi.DeleteSysWheelRecordsByIds) // 批量删除sysWheelRecords表
		sysWheelRecordsRouter.PUT("updateSysWheelRecords", sysWheelRecordsApi.UpdateSysWheelRecords)    // 更新sysWheelRecords表
	}
	{
		sysWheelRecordsRouterWithoutRecord.GET("findSysWheelRecords", sysWheelRecordsApi.FindSysWheelRecords)        // 根据ID获取sysWheelRecords表
		sysWheelRecordsRouterWithoutRecord.GET("getSysWheelRecordsList", sysWheelRecordsApi.GetSysWheelRecordsList)  // 获取sysWheelRecords表列表
	}
}
