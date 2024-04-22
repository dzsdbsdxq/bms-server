package wheelRecords

import (
	"bms-server/global"
    "bms-server/model/wheelRecords"
    wheelRecordsReq "bms-server/model/wheelRecords/request"
    "bms-server/model/common/response"
    "bms-server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysWheelRecordsApi struct {
}

var sysWheelRecordsService = service.ServiceGroupApp.WheelRecordsServiceGroup.SysWheelRecordsService


// CreateSysWheelRecords 创建sysWheelRecords表
// @Tags SysWheelRecords
// @Summary 创建sysWheelRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelRecords.SysWheelRecords true "创建sysWheelRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysWheelRecords/createSysWheelRecords [post]
func (sysWheelRecordsApi *SysWheelRecordsApi) CreateSysWheelRecords(c *gin.Context) {
	var sysWheelRecords wheelRecords.SysWheelRecords
	err := c.ShouldBindJSON(&sysWheelRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelRecordsService.CreateSysWheelRecords(&sysWheelRecords); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysWheelRecords 删除sysWheelRecords表
// @Tags SysWheelRecords
// @Summary 删除sysWheelRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelRecords.SysWheelRecords true "删除sysWheelRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysWheelRecords/deleteSysWheelRecords [delete]
func (sysWheelRecordsApi *SysWheelRecordsApi) DeleteSysWheelRecords(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysWheelRecordsService.DeleteSysWheelRecords(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysWheelRecordsByIds 批量删除sysWheelRecords表
// @Tags SysWheelRecords
// @Summary 批量删除sysWheelRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysWheelRecords/deleteSysWheelRecordsByIds [delete]
func (sysWheelRecordsApi *SysWheelRecordsApi) DeleteSysWheelRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysWheelRecordsService.DeleteSysWheelRecordsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysWheelRecords 更新sysWheelRecords表
// @Tags SysWheelRecords
// @Summary 更新sysWheelRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheelRecords.SysWheelRecords true "更新sysWheelRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysWheelRecords/updateSysWheelRecords [put]
func (sysWheelRecordsApi *SysWheelRecordsApi) UpdateSysWheelRecords(c *gin.Context) {
	var sysWheelRecords wheelRecords.SysWheelRecords
	err := c.ShouldBindJSON(&sysWheelRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelRecordsService.UpdateSysWheelRecords(sysWheelRecords); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysWheelRecords 用id查询sysWheelRecords表
// @Tags SysWheelRecords
// @Summary 用id查询sysWheelRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelRecords.SysWheelRecords true "用id查询sysWheelRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysWheelRecords/findSysWheelRecords [get]
func (sysWheelRecordsApi *SysWheelRecordsApi) FindSysWheelRecords(c *gin.Context) {
	ID := c.Query("ID")
	if resysWheelRecords, err := sysWheelRecordsService.GetSysWheelRecords(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysWheelRecords": resysWheelRecords}, c)
	}
}

// GetSysWheelRecordsList 分页获取sysWheelRecords表列表
// @Tags SysWheelRecords
// @Summary 分页获取sysWheelRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelRecordsReq.SysWheelRecordsSearch true "分页获取sysWheelRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysWheelRecords/getSysWheelRecordsList [get]
func (sysWheelRecordsApi *SysWheelRecordsApi) GetSysWheelRecordsList(c *gin.Context) {
	var pageInfo wheelRecordsReq.SysWheelRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysWheelRecordsService.GetSysWheelRecordsInfoList(pageInfo); err != nil {
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
