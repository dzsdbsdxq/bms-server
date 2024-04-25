package router

import (
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
		//middleware.JWTMemberAuth(),
		Router.POST("memberInfo", plugApi.GetMemberInfo)
		// middleware.JWTMemberAuth(),
		Router.POST("getAwardInfo", plugApi.GetAwardInfo)
	}
}
