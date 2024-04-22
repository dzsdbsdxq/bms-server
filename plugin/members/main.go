package members

import (
	"bms-server/plugin/members/router"
	"github.com/gin-gonic/gin"
)

type RegisterPlugin struct {
}

func CreateMemberPlug() *RegisterPlugin {
	return &RegisterPlugin{}
}

func (*RegisterPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRegisterRouter(group)
}

func (*RegisterPlugin) RouterPath() string {
	return "members"
}
