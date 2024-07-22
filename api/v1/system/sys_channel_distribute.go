package system

import (
	"bms-server/global"
	"bms-server/model/common/request"
	"bms-server/model/common/response"
	"bms-server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChannelDistrubuteApi struct{}

// CreateSysChannelDistrubute
// @Tags      SysChannelDistrubute
// @Summary   创建SysChannelDistrubute
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysChannelDistrubute      true  "SysChannelDistrubute模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysChannelDistrubute"
// @Router    /sysChannelDistrubute/createSysChannelDistrubute [post]
func (s *ChannelDistrubuteApi) CreateSysChannelDistrubute(c *gin.Context) {
	var dictionary system.SysChannelDistrubute
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = channelDistrubuteService.CreateSysChannelDistrubute(dictionary)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysChannelDistrubute
// @Tags      SysChannelDistrubute
// @Summary   删除SysChannelDistrubute
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysChannelDistrubute     true  "SysChannelDistrubute模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysChannelDistrubute"
// @Router    /sysExtends/deleteChannelDistrubute [delete]
func (s *ChannelDistrubuteApi) DeleteSysChannelDistrubute(c *gin.Context) {
	var dictionary system.SysChannelDistrubute
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = channelDistrubuteService.DeleteSysChannelDistrubute(dictionary)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysChannelDistrubute
// @Tags      SysExtends
// @Summary   更新SysExtends
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysChannelDistrubute      true  "SysChannelDistrubute模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysExtends"
// @Router    /sysChannelDistrubute/updateSysChannelDistrubute [put]
func (s *ChannelDistrubuteApi) UpdateSysChannelDistrubute(c *gin.Context) {
	var dictionary system.SysChannelDistrubute
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = channelDistrubuteService.UpdateSysChannelDistrubute(&dictionary)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysChannelDistrubute
// @Tags      SysChannelDistrubute
// @Summary   用id查询SysChannelDistrubute
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysChannelDistrubute                              true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysChannelDistrubute"
// @Router    /sysChannelDistrubute/findSysChannelDistrubute [get]
func (s *ChannelDistrubuteApi) FindSysChannelDistrubute(c *gin.Context) {
	var dictionary system.SysChannelDistrubute
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysChannelDistrubute, err := channelDistrubuteService.GetSysChannelDistrubute(dictionary.ID)
	if err != nil {
		global.GVA_LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(sysChannelDistrubute, "查询成功", c)
}

// GetSysChannelDistrubuteList
// @Tags      SysChannelDistrubute
// @Summary   分页获取SysChannelDistrubute列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysChannelDistrubute列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysChannelDistrubute/getSysChannelDistrubuteList [get]
func (s *ChannelDistrubuteApi) GetSysChannelDistrubuteList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := channelDistrubuteService.GetSysChannelDistrubuteInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
