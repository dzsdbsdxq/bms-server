package router

import (
	"bms-server/plugin/wheels/api"
	"github.com/gin-gonic/gin"
)

type WheelRouter struct {
}

func (s *WheelRouter) InitWheelRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WheelApi
	//{
	//	plugRouter.POST("register", plugApi.ApiName)
	//}
	{
		plugRouter.GET("/", plugApi.IndexPage)
	}
}
