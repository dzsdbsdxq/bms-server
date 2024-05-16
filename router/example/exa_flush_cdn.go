package example

import (
	v1 "bms-server/api/v1"
	"github.com/gin-gonic/gin"
)

type FlushCdnRouter struct{}

func (e *FlushCdnRouter) InitFlushCdnRouter(Router *gin.RouterGroup) {
	flushCdnRouter := Router.Group("sysCdnRecords")
	exaFlushCdnApi := v1.ApiGroupApp.ExampleApiGroup.FlushCdnApi
	{
		flushCdnRouter.POST("createSysCdnRecords", exaFlushCdnApi.CreateFlushCdn)
	}
}
