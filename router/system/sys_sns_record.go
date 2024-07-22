package system

import (
	v1 "bms-server/api/v1"
	"github.com/gin-gonic/gin"
)

type SnsRecordRouter struct{}

func (s *SnsRecordRouter) InitSysSnsRecordRouter(Router *gin.RouterGroup) {
	snsRecordRouter := Router.Group("sysSnsRecords")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.SnsRecordApi
	{
		snsRecordRouter.POST("createSysSnsRecords", authorityMenuApi.CreateSysSnsRecord)             // 新建SysSnsRecord
		snsRecordRouter.DELETE("deleteSysSnsRecords", authorityMenuApi.DeleteSysSnsRecord)           // 删除SysSnsRecord
		snsRecordRouter.DELETE("deleteSysSnsRecordsByIds", authorityMenuApi.DeleteSysSnsRecordByIds) // 批量删除SysSnsRecord
		snsRecordRouter.GET("findSysSnsRecords", authorityMenuApi.FindSysSnsRecord)                  // 根据ID获取SysSnsRecord
		snsRecordRouter.GET("getSysSnsRecordsList", authorityMenuApi.GetSysSnsRecordList)            // 获取SysSnsRecord列表

	}
}
