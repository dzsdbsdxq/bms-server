package api

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/plugin/channel/model/request"
	response2 "bms-server/plugin/channel/model/response"
	"bms-server/plugin/channel/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChannelApi struct{}

func (s *ChannelApi) GetChannelDetail(c *gin.Context) {
	var channel request.ChannelRequest
	err := c.ShouldBindQuery(&channel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	channelInfo, err := service.ServiceGroupApp.ChannelService.GetSysChannel(channel.AppId)
	if err != nil {
		global.GVA_LOG.Error("获取渠道信息失败!", zap.Error(err))
		response.FailWithMessage("获取渠道信息失败", c)
		return
	}
	if !*channelInfo.Status {
		response.FailWithMessage("渠道已停用", c)
		return
	}
	resp := response2.ChannelDetailResponse{
		Name:      channelInfo.Name,
		AppId:     channelInfo.AppId,
		JumpValue: fmt.Sprintf("%s%s", channelInfo.JumpType, channelInfo.JumpValue),
	}
	response.OkWithDetailed(resp, "获取成功", c)
}

func (s *ChannelApi) UpdateChannelFields(c *gin.Context) {
	var channel request.UpdateChannelRequest
	err := c.ShouldBindQuery(&channel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fields := map[string]bool{"views": true, "down": true}
	if _, found := fields[channel.Field]; !found {
		fmt.Println(found)
		response.FailWithMessage("参数错误", c)
		return
	}
	if channel.AppId == "" {
		//默认
		channel.AppId = "official"
	}

	err = service.ServiceGroupApp.ChannelService.UpdateSysChannelField(channel.Field, channel.AppId)
	if err != nil {
		global.GVA_LOG.Error("修改渠道信息失败!", zap.Error(err))
		response.FailWithMessage("error", c)
		return
	}
	response.Ok(c)
}
