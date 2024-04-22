package members

import (
	"bms-server/api/v1"
	"bms-server/middleware"
	"github.com/gin-gonic/gin"
)

type SysMembersRouter struct {
}

// InitSysMembersRouter 初始化 sysMembers表 路由信息
func (s *SysMembersRouter) InitSysMembersRouter(Router *gin.RouterGroup) {
	sysMembersRouter := Router.Group("sysMembers").Use(middleware.OperationRecord())
	sysMembersRouterWithoutRecord := Router.Group("sysMembers")
	var sysMembersApi = v1.ApiGroupApp.MembersApiGroup.SysMembersApi
	{
		sysMembersRouter.POST("createSysMembers", sysMembersApi.CreateSysMembers)             // 新建sysMembers表
		sysMembersRouter.DELETE("deleteSysMembers", sysMembersApi.DeleteSysMembers)           // 删除sysMembers表
		sysMembersRouter.DELETE("deleteSysMembersByIds", sysMembersApi.DeleteSysMembersByIds) // 批量删除sysMembers表
		sysMembersRouter.PUT("updateSysMembers", sysMembersApi.UpdateSysMembers)              // 更新sysMembers表
		sysMembersRouter.POST("resetPassword", sysMembersApi.ResetPassword)                   // 重置密码
		sysMembersRouter.POST("scoreUpdate", sysMembersApi.ScoreUpdate)                       // 积分变更
	}
	{
		sysMembersRouterWithoutRecord.GET("findSysMembers", sysMembersApi.FindSysMembers)       // 根据ID获取sysMembers表
		sysMembersRouterWithoutRecord.GET("getSysMembersList", sysMembersApi.GetSysMembersList) // 获取sysMembers表列表
	}
}
