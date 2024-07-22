package router

import (
	"bms-server/plugin/channel/api"
	"github.com/gin-gonic/gin"
)

type ChannelRouter struct{}

func (s *ChannelRouter) InitChannelRouter(Router *gin.RouterGroup) {
	ChannelApi := api.ApiGroupApp.ChannelApi
	{
		Router.GET("get", ChannelApi.GetChannelDetail)       // 获取渠道信息
		Router.GET("update", ChannelApi.UpdateChannelFields) // 修改渠道信息
	}
}
