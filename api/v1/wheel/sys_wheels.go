package wheel

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/model/wheel"
	wheelReq "bms-server/model/wheel/request"
	response2 "bms-server/model/wheel/response"
	"bms-server/service"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type SysWheelsApi struct {
}

var sysWheelsService = service.ServiceGroupApp.WheelServiceGroup.SysWheelsService

// CreateSysWheels 创建sysWheels表
// @Tags SysWheels
// @Summary 创建sysWheels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheel.SysWheels true "创建sysWheels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysWheels/createSysWheels [post]
func (sysWheelsApi *SysWheelsApi) CreateSysWheels(c *gin.Context) {
	var sysWheels wheel.SysWheels
	err := c.ShouldBindJSON(&sysWheels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysWheels.Uuid = uuid.NewV4().String()

	if err := sysWheelsService.CreateSysWheels(&sysWheels); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysWheels 删除sysWheels表
// @Tags SysWheels
// @Summary 删除sysWheels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheel.SysWheels true "删除sysWheels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysWheels/deleteSysWheels [delete]
func (sysWheelsApi *SysWheelsApi) DeleteSysWheels(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysWheelsService.DeleteSysWheels(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysWheelsByIds 批量删除sysWheels表
// @Tags SysWheels
// @Summary 批量删除sysWheels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysWheels/deleteSysWheelsByIds [delete]
func (sysWheelsApi *SysWheelsApi) DeleteSysWheelsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysWheelsService.DeleteSysWheelsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysWheels 更新sysWheels表
// @Tags SysWheels
// @Summary 更新sysWheels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wheel.SysWheels true "更新sysWheels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysWheels/updateSysWheels [put]
func (sysWheelsApi *SysWheelsApi) UpdateSysWheels(c *gin.Context) {
	var sysWheels wheel.SysWheels
	err := c.ShouldBindJSON(&sysWheels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelsService.UpdateSysWheels(sysWheels); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysWheels 用id查询sysWheels表
// @Tags SysWheels
// @Summary 用id查询sysWheels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheel.SysWheels true "用id查询sysWheels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysWheels/findSysWheels [get]
func (sysWheelsApi *SysWheelsApi) FindSysWheels(c *gin.Context) {
	ID := c.Query("ID")
	if resysWheels, err := sysWheelsService.GetSysWheels(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysWheels": resysWheels}, c)
	}
}

// GetSysWheelsList 分页获取sysWheels表列表
// @Tags SysWheels
// @Summary 分页获取sysWheels表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wheelReq.SysWheelsSearch true "分页获取sysWheels表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysWheels/getSysWheelsList [get]
func (sysWheelsApi *SysWheelsApi) GetSysWheelsList(c *gin.Context) {
	var pageInfo wheelReq.SysWheelsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := sysWheelsService.GetSysWheelsInfoList(pageInfo); err != nil {
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

func (sysWheelsApi *SysWheelsApi) GetSysWheelsTemplateInfo(c *gin.Context) {
	//获取配置文件中的目录
	templateDir := global.GVA_CONFIG.System.WheelsTemplateDir
	if templateDir == "" {
		templateDir, _ = os.Getwd()
		templateDir += "/templates"
	}
	templateDir += "/wheels"
	// 读取目录内容
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	wheelTemplateDir := make([]response2.WheelTemplateDirInfo, 0)
	// 遍历目录内容，过滤出文件夹
	for _, entry := range entries {
		if entry.IsDir() {
			wheelTemplateDir = append(wheelTemplateDir, response2.WheelTemplateDirInfo{
				Name:  entry.Name(),
				Thumb: "",
			})
		}
	}
	response.OkWithData(wheelTemplateDir, c)
	return
}

func (sysWheelsApi *SysWheelsApi) GetSysWheelsTemplateConfig(c *gin.Context) {
	//获取配置文件中的目录
	templateDir := global.GVA_CONFIG.System.WheelsTemplateDir
	if templateDir == "" {
		templateDir, _ = os.Getwd()
		templateDir += "/templates"
	}
	templateDir += "/wheels"

	ID := c.Query("ID")
	if resysWheels, err := sysWheelsService.GetSysWheels(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		configJsonFile, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/config.json", templateDir, resysWheels.Template))
		response.OkWithDetailed(gin.H{
			"y": string(configJsonFile),
			"z": resysWheels.PageParams,
		}, "获取成功", c)
	}
}

func (sysWheelsApi *SysWheelsApi) UpdatePageParams(c *gin.Context) {
	var sysWheels wheel.SysWheels
	err := c.ShouldBindJSON(&sysWheels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysWheelsService.UpdatePageParams(sysWheels.ID, sysWheels.PageParams); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
