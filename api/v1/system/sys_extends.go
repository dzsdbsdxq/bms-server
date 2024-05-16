package system

import (
	"bms-server/global"
	"bms-server/model/common/request"
	"bms-server/model/common/response"
	"bms-server/model/system"
	"bms-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExtendsApi struct{}

// CreateSysExtends
// @Tags      SysExtends
// @Summary   创建SysExtends
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtends           true  "SysExtends模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysExtends"
// @Router    /sysExtends/createSysExtends [post]
func (s *ExtendsApi) CreateSysExtends(c *gin.Context) {
	var dictionary system.SysExtends
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dictionary.AppId = fmt.Sprintf("YSZH5-%s", utils.GetRandomId(8))
	err = extendsService.CreateSysExtends(dictionary)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysExtends
// @Tags      SysExtends
// @Summary   删除SysExtends
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtends          true  "SysExtends模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionary"
// @Router    /sysExtends/deleteExtends [delete]
func (s *ExtendsApi) DeleteSysExtends(c *gin.Context) {
	var dictionary system.SysExtends
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = extendsService.DeleteSysExtends(dictionary)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysExtends
// @Tags      SysExtends
// @Summary   更新SysExtends
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtends           true  "SysExtends模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysExtends"
// @Router    /sysExtends/updateSysExtends [put]
func (s *ExtendsApi) UpdateSysExtends(c *gin.Context) {
	var dictionary system.SysExtends
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = extendsService.UpdateSysExtends(&dictionary)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysExtends
// @Tags      SysExtends
// @Summary   用id查询SysExtends
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysExtends                                       true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysExtends"
// @Router    /sysExtends/findSysExtends [get]
func (s *ExtendsApi) FindSysExtends(c *gin.Context) {
	var dictionary system.SysExtends
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysExtends, err := extendsService.GetSysExtends(dictionary.ID)
	if err != nil {
		global.GVA_LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysExtends": sysExtends}, "查询成功", c)
}

// GetSysExtendsList
// @Tags      SysExtends
// @Summary   分页获取SysExtends列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysExtends列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysExtends/getSysExtendsList [get]
func (s *ExtendsApi) GetSysExtendsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := extendsService.GetSysExtendsInfoList(pageInfo)
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
