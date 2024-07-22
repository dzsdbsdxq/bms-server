package api

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/plugin/extends/model/request"
	response2 "bms-server/plugin/extends/model/response"
	"bms-server/plugin/extends/service"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
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
	extendDetailInfo, _ := service.ServiceGroupApp.ExtendsDetailService.ExtendsDetailInfo(extends.AppId)

	if extendInfo.DType == `gps` {
		lists := make([]response2.ExtendDetailResponse, 0)
		if len(extendDetailInfo) > 0 {
			for _, detail := range extendDetailInfo {
				var tmpJson interface{}
				_ = jsoniter.Unmarshal([]byte(detail.Value), &tmpJson)
				lists = append(lists, response2.ExtendDetailResponse{
					SysExtendsDetail: detail,
					Jsx:              tmpJson,
				})
			}
		}
		response.OkWithDetailed(lists, "获取成功", c)
	} else if extendInfo.DType == `array` {
		lists := make(map[string]interface{}, 0)
		for _, detail := range extendDetailInfo {
			var tmpJson interface{}
			_ = jsoniter.Unmarshal([]byte(detail.Value), &tmpJson)
			lists[detail.Key] = tmpJson
		}
		response.OkWithDetailed(lists, "获取成功", c)
	} else {
		lists := make(map[string]string, 0)
		if len(extendDetailInfo) > 0 {
			for _, detail := range extendDetailInfo {
				lists[detail.Key] = detail.Value
			}
		}
		response.OkWithDetailed(lists, "获取成功", c)
	}
}
