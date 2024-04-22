package wheels

import (
	"bms-server/plugin/wheels/router"
	"github.com/gin-gonic/gin"
)

type WheelPlugin struct {
}

func CreateWheelPlug() *WheelPlugin {
	return &WheelPlugin{}
}

func (w *WheelPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitWheelRouter(group)
}

func (w *WheelPlugin) RouterPath() string {
	return "wheels"
}
