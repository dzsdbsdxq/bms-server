package extends

import (
	"bms-server/plugin/extends/router"
	"github.com/gin-gonic/gin"
)

type ExtendsPlugin struct{}

func CreateExtendsPlug() *ExtendsPlugin {
	return &ExtendsPlugin{}
}

func (*ExtendsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitExtendsRouter(group)
}

func (*ExtendsPlugin) RouterPath() string {
	return "extends"
}
