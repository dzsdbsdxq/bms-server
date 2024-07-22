package channel

import (
	"bms-server/plugin/channel/router"
	"github.com/gin-gonic/gin"
)

type ChannelPlugin struct{}

func CreateChannelPlug() *ChannelPlugin {
	return &ChannelPlugin{}
}

func (*ChannelPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitChannelRouter(group)
}

func (*ChannelPlugin) RouterPath() string {
	return "channel"
}
