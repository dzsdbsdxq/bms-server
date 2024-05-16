package system

import (
	v1 "bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type ExtendsDetailRouter struct{}

func (s *ExtendsDetailRouter) InitSysExtendsDetailRouter(Router *gin.RouterGroup) {
	extendsDetailRouter := Router.Group("sysExtendsDetails").Use(middleware.OperationRecord())
	extendsDetailRouterWithoutRecord := Router.Group("sysExtendsDetails")
	sysExtendsDetailApi := v1.ApiGroupApp.SystemApiGroup.ExtendsDetailApi
	{
		extendsDetailRouter.POST("createSysExtendsDetails", sysExtendsDetailApi.CreateSysExtendsDetail)   // 新建SysExtendsDetail
		extendsDetailRouter.DELETE("deleteSysExtendsDetails", sysExtendsDetailApi.DeleteSysExtendsDetail) // 删除SysExtendsDetail
		extendsDetailRouter.PUT("updateSysExtendsDetails", sysExtendsDetailApi.UpdateSysExtendsDetail)    // 更新SysExtendsDetail
	}
	{
		extendsDetailRouterWithoutRecord.GET("findSysExtendsDetails", sysExtendsDetailApi.FindSysExtendsDetail)       // 根据ID获取SysExtendsDetail
		extendsDetailRouterWithoutRecord.GET("getSysExtendsDetailsList", sysExtendsDetailApi.GetSysExtendsDetailList) // 获取SysExtendsDetail列表
	}
}
