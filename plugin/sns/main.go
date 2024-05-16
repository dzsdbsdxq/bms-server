package sns

import (
	"bms-server/plugin/sns/router"
	"github.com/gin-gonic/gin"
)

type SmsPlugin struct{}

func CreateSnsPlug() *SmsPlugin {
	return &SmsPlugin{}
}

func (*SmsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitSnsRouter(group)
}

func (*SmsPlugin) RouterPath() string {
	return "sns"
}
