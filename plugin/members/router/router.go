package router

import (
	"bms-server/middleware"
	"bms-server/plugin/members/api"
	"github.com/gin-gonic/gin"
)

type PluginMemberRouter struct {
}

func (s *PluginMemberRouter) InitRegisterRouter(Router *gin.RouterGroup) {
	plugApi := api.ApiGroupApp.PluginMemberApi
	{
		Router.POST("register", plugApi.Register)
		Router.POST("login", plugApi.Login)
		Router.POST("memberInfo", middleware.JWTMemberAuth(), plugApi.GetMemberInfo)
	}
}
