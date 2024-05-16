package router

import (
	"bms-server/plugin/extends/api"
	"github.com/gin-gonic/gin"
)

type ExtendsRouter struct{}

func (s *ExtendsRouter) InitExtendsRouter(Router *gin.RouterGroup) {
	ExtendsApi := api.ApiGroupApp.ExtendsApi.GetExtendsDetail
	{
		Router.GET("get", ExtendsApi) // 获取扩展信息
	}
}
