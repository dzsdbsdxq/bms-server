package api

import (
	"bms-server/api/v1/members"
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/plugin/members/model"
	"bms-server/plugin/members/service"
	"bms-server/service/wheel"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
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
	// memberInfo := utils.GetMemberInfo(c)
	// response.OkWithDetailed(memberInfo, "success", c)
	awards := make([]wheel.AwardBatchService, 0)
	awards = append(awards, wheel.AwardBatchService{
		WheelId:      1,
		PrizeId:      1,
		PrizeType:    0,
		IsHit:        1,
		Probability:  100,
		Uuid:         "u-1",
		PrizeName:    "p-1",
		TotalBalance: 2,
		TotalAmount:  2,
		UpdateTime:   0,
		StartTime:    1713888257,
		EndTime:      1713974657,
	}, wheel.AwardBatchService{
		WheelId:      1,
		PrizeId:      2,
		PrizeType:    0,
		Probability:  1000,
		IsHit:        1,
		Uuid:         "u-2",
		PrizeName:    "p-2",
		TotalBalance: 1,
		TotalAmount:  1,
		UpdateTime:   0,
		StartTime:    1713888257,
		EndTime:      1713974657,
	}, wheel.AwardBatchService{
		WheelId:      1,
		PrizeId:      3,
		PrizeType:    0,
		IsHit:        0,
		Probability:  8900,
		Uuid:         "u-3",
		PrizeName:    "p-3",
		TotalBalance: 5,
		TotalAmount:  5,
		UpdateTime:   0,
		StartTime:    1713888257,
		EndTime:      1713974657,
	})

	var wg sync.WaitGroup
	for _, aw := range awards {
		aw := aw
		wg.Add(1)
		go func(aw wheel.AwardBatchService) {
			wheel.InitAwardPool(&aw)
			wg.Done()
		}(aw)
	}
	wg.Wait()

	response.OkWithMessage("成功", c)
}

func (p *PluginMemberApi) GetAwardInfo(c *gin.Context) {
	aw := wheel.WinPrize(3, true, 2)
	//aw := wheel.WinPrizeCustom(1, true)
	response.OkWithDetailed(aw, "success", c)
}
