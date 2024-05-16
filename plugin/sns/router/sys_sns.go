package router

import (
	"bms-server/plugin/sns/api"
	"github.com/gin-gonic/gin"
)

type SnsRouter struct{}

func (s *SnsRouter) InitSnsRouter(Router *gin.RouterGroup) {
	SnsApi := api.ApiGroupApp.SnsApi
	{
		Router.POST("getValidateCode", SnsApi.GetValidateCode) // 获取验证码校验code
		Router.POST("getEncode", SnsApi.GetEncode)
		Router.POST("code", SnsApi.SendSMS)
		//Router.POST("userinfo", SnsApi.GetUserInfo)
		// 登录登出刷新token无需鉴权
		Router.POST("login", SnsApi.LoginHandler)
		Router.POST("codeEn", SnsApi.GetCodeEn)
		//Router.POST("loginByMobile", SnsApi.LoginHandler2)
	}
}
