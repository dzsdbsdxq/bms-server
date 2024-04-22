package wheelAddress

import (
	"bms-server/global"
    "bms-server/model/wheelAddress"
    wheelAddressReq "bms-server/model/wheelAddress/request"
    "bms-server/model/common/response"
    "bms-server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysWheelAddressApi struct {
}

var sysWheelAddressService = service.ServiceGroupApp.WheelAddressServiceGroup.SysWheelAddressService


// CreateSysWheelAddress 创建sysWheelAddress表
// @Tags SysWheelAddress
// @Summary 创建sysWheelAddress表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelAddress.SysWheelAddress true "创建sysWheelAddress表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysWheelAddress/createSysWheelAddress [post]
func (sysWheelAddressApi *SysWheelAddressApi) CreateSysWheelAddress(c *gin.Context) {
	var sysWheelAddress wheelAddress.SysWheelAddress
	err := c.ShouldBindJSON(&sysWheelAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelAddressService.CreateSysWheelAddress(&sysWheelAddress); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysWheelAddress 删除sysWheelAddress表
// @Tags SysWheelAddress
// @Summary 删除sysWheelAddress表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelAddress.SysWheelAddress true "删除sysWheelAddress表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysWheelAddress/deleteSysWheelAddress [delete]
func (sysWheelAddressApi *SysWheelAddressApi) DeleteSysWheelAddress(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysWheelAddressService.DeleteSysWheelAddress(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysWheelAddressByIds 批量删除sysWheelAddress表
// @Tags SysWheelAddress
// @Summary 批量删除sysWheelAddress表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysWheelAddress/deleteSysWheelAddressByIds [delete]
func (sysWheelAddressApi *SysWheelAddressApi) DeleteSysWheelAddressByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysWheelAddressService.DeleteSysWheelAddressByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysWheelAddress 更新sysWheelAddress表
// @Tags SysWheelAddress
// @Summary 更新sysWheelAddress表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelAddress.SysWheelAddress true "更新sysWheelAddress表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysWheelAddress/updateSysWheelAddress [put]
func (sysWheelAddressApi *SysWheelAddressApi) UpdateSysWheelAddress(c *gin.Context) {
	var sysWheelAddress wheelAddress.SysWheelAddress
	err := c.ShouldBindJSON(&sysWheelAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelAddressService.UpdateSysWheelAddress(sysWheelAddress); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysWheelAddress 用id查询sysWheelAddress表
// @Tags SysWheelAddress
// @Summary 用id查询sysWheelAddress表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelAddress.SysWheelAddress true "用id查询sysWheelAddress表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysWheelAddress/findSysWheelAddress [get]
func (sysWheelAddressApi *SysWheelAddressApi) FindSysWheelAddress(c *gin.Context) {
	ID := c.Query("ID")
	if resysWheelAddress, err := sysWheelAddressService.GetSysWheelAddress(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysWheelAddress": resysWheelAddress}, c)
	}
}

// GetSysWheelAddressList 分页获取sysWheelAddress表列表
// @Tags SysWheelAddress
// @Summary 分页获取sysWheelAddress表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelAddressReq.SysWheelAddressSearch true "分页获取sysWheelAddress表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysWheelAddress/getSysWheelAddressList [get]
func (sysWheelAddressApi *SysWheelAddressApi) GetSysWheelAddressList(c *gin.Context) {
	var pageInfo wheelAddressReq.SysWheelAddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysWheelAddressService.GetSysWheelAddressInfoList(pageInfo); err != nil {
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
