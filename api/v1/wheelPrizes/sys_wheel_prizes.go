package wheelPrizes

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/model/wheelPrizes"
	wheelPrizesReq "bms-server/model/wheelPrizes/request"
	"bms-server/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type SysWheelPrizesApi struct {
}

var sysWheelPrizesService = service.ServiceGroupApp.WheelPrizesServiceGroup.SysWheelPrizesService

// CreateSysWheelPrizes 创建sysWheelPrizes表
// @Tags SysWheelPrizes
// @Summary 创建sysWheelPrizes表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelPrizes.SysWheelPrizes true "创建sysWheelPrizes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysWheelPrizes/createSysWheelPrizes [post]
func (sysWheelPrizesApi *SysWheelPrizesApi) CreateSysWheelPrizes(c *gin.Context) {
	var sysWheelPrizes wheelPrizes.SysWheelPrizes
	err := c.ShouldBindJSON(&sysWheelPrizes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysWheelPrizes.Uuid = uuid.NewV4().String()

	if err := sysWheelPrizesService.CreateSysWheelPrizes(&sysWheelPrizes); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysWheelPrizes 删除sysWheelPrizes表
// @Tags SysWheelPrizes
// @Summary 删除sysWheelPrizes表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelPrizes.SysWheelPrizes true "删除sysWheelPrizes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysWheelPrizes/deleteSysWheelPrizes [delete]
func (sysWheelPrizesApi *SysWheelPrizesApi) DeleteSysWheelPrizes(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysWheelPrizesService.DeleteSysWheelPrizes(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysWheelPrizesByIds 批量删除sysWheelPrizes表
// @Tags SysWheelPrizes
// @Summary 批量删除sysWheelPrizes表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysWheelPrizes/deleteSysWheelPrizesByIds [delete]
func (sysWheelPrizesApi *SysWheelPrizesApi) DeleteSysWheelPrizesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysWheelPrizesService.DeleteSysWheelPrizesByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysWheelPrizes 更新sysWheelPrizes表
// @Tags SysWheelPrizes
// @Summary 更新sysWheelPrizes表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelPrizes.SysWheelPrizes true "更新sysWheelPrizes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysWheelPrizes/updateSysWheelPrizes [put]
func (sysWheelPrizesApi *SysWheelPrizesApi) UpdateSysWheelPrizes(c *gin.Context) {
	var sysWheelPrizes wheelPrizes.SysWheelPrizes
	err := c.ShouldBindJSON(&sysWheelPrizes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelPrizesService.UpdateSysWheelPrizes(sysWheelPrizes); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysWheelPrizes 用id查询sysWheelPrizes表
// @Tags SysWheelPrizes
// @Summary 用id查询sysWheelPrizes表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelPrizes.SysWheelPrizes true "用id查询sysWheelPrizes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysWheelPrizes/findSysWheelPrizes [get]
func (sysWheelPrizesApi *SysWheelPrizesApi) FindSysWheelPrizes(c *gin.Context) {
	ID := c.Query("ID")
	if resysWheelPrizes, err := sysWheelPrizesService.GetSysWheelPrizes(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysWheelPrizes": resysWheelPrizes}, c)
	}
}

// GetSysWheelPrizesList 分页获取sysWheelPrizes表列表
// @Tags SysWheelPrizes
// @Summary 分页获取sysWheelPrizes表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelPrizesReq.SysWheelPrizesSearch true "分页获取sysWheelPrizes表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysWheelPrizes/getSysWheelPrizesList [get]
func (sysWheelPrizesApi *SysWheelPrizesApi) GetSysWheelPrizesList(c *gin.Context) {
	var pageInfo wheelPrizesReq.SysWheelPrizesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysWheelPrizesService.GetSysWheelPrizesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
