package api

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/plugin/extends/model/request"
	"bms-server/plugin/extends/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExtendsApi struct{}

func (s *ExtendsApi) GetExtendsDetail(c *gin.Context) {
	var extends request.ExtendsRequest
	err := c.ShouldBindQuery(&extends)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	extendInfo, err := service.ServiceGroupApp.ExtendsService.GetSysExtends(extends.AppId)
	if err != nil {
		global.GVA_LOG.Error("获取扩展信息失败!", zap.Error(err))
		response.FailWithMessage("获取扩展信息失败", c)
		return
	}
	if !*extendInfo.Status {
		response.FailWithMessage("扩展已停用", c)
		return
	}
	lists := make(map[string]string, 0)
	extendDetailInfo, err := service.ServiceGroupApp.ExtendsDetailService.ExtendsDetailInfo(extends.AppId)
	if len(extendDetailInfo) > 0 {
		for _, detail := range extendDetailInfo {
			lists[detail.Key] = detail.Value
		}

	}
	response.OkWithDetailed(lists, "获取成功", c)
}
