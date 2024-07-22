package system

import (
	"bms-server/global"
	"bms-server/model/common/request"
	"bms-server/model/common/response"
	"bms-server/model/system"
	systemReq "bms-server/model/system/request"
	"bms-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SnsRecordApi struct{}

func (s *SnsRecordApi) CreateSysSnsRecord(c *gin.Context) {
	var sysSnsRecord system.SysSnsRecord
	err := c.ShouldBindJSON(&sysSnsRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = snsRecordService.CreateSysSnsRecord(sysSnsRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *SnsRecordApi) DeleteSysSnsRecord(c *gin.Context) {
	var sysSnsRecord system.SysSnsRecord
	err := c.ShouldBindQuery(&sysSnsRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = snsRecordService.DeleteSysSnsRecord(sysSnsRecord)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *SnsRecordApi) DeleteSysSnsRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = snsRecordService.DeleteSysSnsRecordByIds(IDS)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

func (s *SnsRecordApi) FindSysSnsRecord(c *gin.Context) {
	var sysSnsRecord system.SysSnsRecord
	err := c.ShouldBindQuery(&sysSnsRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(sysSnsRecord, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysSnsRecord, err := snsRecordService.GetSysSnsRecord(sysSnsRecord.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysSnsRecord": reSysSnsRecord}, "查询成功", c)
}

func (s *SnsRecordApi) GetSysSnsRecordList(c *gin.Context) {
	var pageInfo systemReq.SysSnsRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := snsRecordService.GetSysSnsRecordInfoList(pageInfo)
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
