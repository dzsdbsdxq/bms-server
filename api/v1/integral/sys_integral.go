package integral

import (
	"bms-server/global"
    "bms-server/model/integral"
    integralReq "bms-server/model/integral/request"
    "bms-server/model/common/response"
    "bms-server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysIntegralApi struct {
}

var sysIntegralService = service.ServiceGroupApp.IntegralServiceGroup.SysIntegralService


// CreateSysIntegral 创建sysIntegral表
// @Tags SysIntegral
// @Summary 创建sysIntegral表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body integral.SysIntegral true "创建sysIntegral表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysIntegral/createSysIntegral [post]
func (sysIntegralApi *SysIntegralApi) CreateSysIntegral(c *gin.Context) {
	var sysIntegral integral.SysIntegral
	err := c.ShouldBindJSON(&sysIntegral)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysIntegralService.CreateSysIntegral(&sysIntegral); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysIntegral 删除sysIntegral表
// @Tags SysIntegral
// @Summary 删除sysIntegral表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body integral.SysIntegral true "删除sysIntegral表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysIntegral/deleteSysIntegral [delete]
func (sysIntegralApi *SysIntegralApi) DeleteSysIntegral(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysIntegralService.DeleteSysIntegral(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysIntegralByIds 批量删除sysIntegral表
// @Tags SysIntegral
// @Summary 批量删除sysIntegral表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysIntegral/deleteSysIntegralByIds [delete]
func (sysIntegralApi *SysIntegralApi) DeleteSysIntegralByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysIntegralService.DeleteSysIntegralByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysIntegral 更新sysIntegral表
// @Tags SysIntegral
// @Summary 更新sysIntegral表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body integral.SysIntegral true "更新sysIntegral表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysIntegral/updateSysIntegral [put]
func (sysIntegralApi *SysIntegralApi) UpdateSysIntegral(c *gin.Context) {
	var sysIntegral integral.SysIntegral
	err := c.ShouldBindJSON(&sysIntegral)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysIntegralService.UpdateSysIntegral(sysIntegral); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysIntegral 用id查询sysIntegral表
// @Tags SysIntegral
// @Summary 用id查询sysIntegral表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query integral.SysIntegral true "用id查询sysIntegral表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysIntegral/findSysIntegral [get]
func (sysIntegralApi *SysIntegralApi) FindSysIntegral(c *gin.Context) {
	ID := c.Query("ID")
	if resysIntegral, err := sysIntegralService.GetSysIntegral(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysIntegral": resysIntegral}, c)
	}
}

// GetSysIntegralList 分页获取sysIntegral表列表
// @Tags SysIntegral
// @Summary 分页获取sysIntegral表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query integralReq.SysIntegralSearch true "分页获取sysIntegral表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysIntegral/getSysIntegralList [get]
func (sysIntegralApi *SysIntegralApi) GetSysIntegralList(c *gin.Context) {
	var pageInfo integralReq.SysIntegralSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysIntegralService.GetSysIntegralInfoList(pageInfo); err != nil {
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
