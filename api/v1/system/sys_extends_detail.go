package system

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/model/system"
	"bms-server/model/system/request"
	"bms-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExtendsDetailApi struct{}

// CreateSysExtendsDetail
// @Tags      SysExtendsDetail
// @Summary   创建SysExtendsDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtendsDetail     true  "SysExtendsDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysExtendsDetail"
// @Router    /sysExtendsDetail/createSysExtendsDetail [post]
func (s *ExtendsDetailApi) CreateSysExtendsDetail(c *gin.Context) {
	var detail system.SysExtendsDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = extendsDetailService.CreateSysExtendsDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysExtendsDetail
// @Tags      SysExtendsDetail
// @Summary   删除SysExtendsDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtendsDetail     true  "SysExtendsDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysExtendsDetail"
// @Router    /sysExtendsDetail/deleteSysExtendsDetail [delete]
func (s *ExtendsDetailApi) DeleteSysExtendsDetail(c *gin.Context) {
	var detail system.SysExtendsDetail
	err := c.ShouldBindQuery(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = extendsDetailService.DeleteSysExtendsDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysExtendsDetail
// @Tags      SysExtendsDetail
// @Summary   更新SysExtendsDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysExtendsDetail     true  "更新SysExtendsDetail"
// @Success   200   {object}  response.Response{msg=string}  "更新SysExtendsDetail"
// @Router    /sysExtendsDetail/updateSysExtendsDetail [put]
func (s *ExtendsDetailApi) UpdateSysExtendsDetail(c *gin.Context) {
	var detail system.SysExtendsDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = extendsDetailService.UpdateSysExtendsDetail(&detail)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysExtendsDetail
// @Tags      SysExtendsDetail
// @Summary   用id查询SysExtendsDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysExtendsDetail                                 true  "用id查询SysExtendsDetail"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysExtendsDetail"
// @Router    /sysExtendsDetail/findSysExtendsDetail [get]
func (s *ExtendsDetailApi) FindSysExtendsDetail(c *gin.Context) {
	var detail system.SysExtendsDetail
	err := c.ShouldBindQuery(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(detail, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysExtendsDetail, err := extendsDetailService.GetSysExtendsDetail(detail.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysExtendsDetail": reSysExtendsDetail}, "查询成功", c)
}

// GetSysExtendsDetailList
// @Tags      SysExtendsDetail
// @Summary   分页获取SysExtendsDetail列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysExtendsDetailSearch                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysExtendsDetail列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysExtendsDetail/getSysExtendsDetailList [get]
func (s *ExtendsDetailApi) GetSysExtendsDetailList(c *gin.Context) {
	var pageInfo request.SysExtendsDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := extendsDetailService.GetSysExtendsDetailInfoList(pageInfo)
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
