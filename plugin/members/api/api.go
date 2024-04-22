package api

import (
	"bms-server/api/v1/members"
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/plugin/members/model"
	"bms-server/plugin/members/service"
	serv2 "bms-server/service"
	"bms-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PluginMemberApi struct{}

// Register
// @Tags Register
// @Summary 普通注册（手机号，密码）
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /members/register[post]
func (p *PluginMemberApi) Register(c *gin.Context) {
	var plug model.Register
	_ = c.ShouldBindJSON(&plug)
	if _, err := service.ServiceGroupApp.RegisterByUserNameAndPassword(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
		//var baseApi systemApi.BaseApi
		//baseApi.TokenNext(c, *res)
	}
}

// Login
// @Tags Login
// @Summary 普通登录（手机号，密码）
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /members/login[post]
func (p *PluginMemberApi) Login(c *gin.Context) {
	var plug model.LoginRequest
	_ = c.ShouldBindJSON(&plug)
	if res, err := service.ServiceGroupApp.LoginByUserNameAndPassword(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		// 登录成功，签发Token
		var baseApi members.SysMembersApi
		baseApi.TokenNext(c, *res)
	}
}

// GetMemberInfo
// @Tags GetMemberInfo
// @Summary 获取用户信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /members/memberInfo[post]
func (p *PluginMemberApi) GetMemberInfo(c *gin.Context) {
	memberInfo := utils.GetMemberInfo(c)
	response.OkWithDetailed(memberInfo, "success", c)
	serv2.ServiceGroupApp.WheelServiceGroup.InitAwardPool()
	//service.ServiceGroup
	//response.OkWithMessage("成功", c)
}
