package api

import (
	"bms-server/model/common/response"
	"bms-server/plugin/sns/config"
	"bms-server/plugin/sns/model/request"
	response2 "bms-server/plugin/sns/model/response"
	"bms-server/plugin/sns/service"
	"bms-server/plugin/sns/utils"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/aes"
	"sync"
)

type SnsApi struct{}

func (s *SnsApi) GetValidateCode(c *gin.Context) {
	var req request.SMSValidateCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 密码通过RSA解密
	decodeData, err := utils.RSADecrypt([]byte(req.Mobile), []byte(config.RSAPrivateBytes))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	valDetail, err := service.ServiceGroupApp.GetValidateCode(string(decodeData))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(valDetail.Data, "操作成功", c)
}

// SendSMS 发送短信
func (s *SnsApi) SendSMS(c *gin.Context) {
	var req request.SMSSendRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 密码通过RSA解密
	decodeData, err := utils.RSADecrypt([]byte(req.Mobile), []byte(config.RSAPrivateBytes))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.SendSMS(string(decodeData), req.Code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("短信发送成功", c)
}

func (s *SnsApi) LoginHandler(c *gin.Context) {
	var req request.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 密码通过RSA解密
	code, err := utils.RSADecrypt([]byte(req.Code), []byte(config.RSAPrivateBytes))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	////验证手机验证码是否正确
	openId, accessToken, err := service.ServiceGroupApp.CheckSMSCode(req.Mobile, string(code))

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//var yszInfo *response2.YszUserInfo
	var yszInfoNew *response2.YszUserInfoNewResponse
	var wg sync.WaitGroup
	wg.Add(1)
	//go func(accessToken string) {
	//	defer wg.Done()
	//	yszInfo, err = service.ServiceGroupApp.GetYszUserInfo(accessToken)
	//}(accessToken)
	go func(accessToken string) {
		defer wg.Done()
		yszInfoNew, err = service.ServiceGroupApp.GetYszMemberInfo(accessToken)
	}(accessToken)
	wg.Wait()

	//if yszInfo == nil {
	//	response.FailWithMessage("获取用户错误[1]", c)
	//	return
	//}
	if yszInfoNew == nil {
		response.FailWithMessage("获取用户错误[2]", c)
		return
	}
	returnData := &response2.YseUserInfoNewResponse{
		ID:                 yszInfoNew.Data.Userid,
		CreatedAt:          yszInfoNew.Data.Regtime,
		UpdatedAt:          "",
		DeletedAt:          "",
		NickName:           yszInfoNew.Data.Nickname,
		Openid:             openId,
		Avatar:             yszInfoNew.Data.Avatar,
		IsDel:              0,
		Platform:           1,
		City:               "",
		Province:           "",
		Country:            "",
		Mobile:             yszInfoNew.Data.Mobile,
		EszUid:             openId,
		Partment:           "",
		LastLogin:          yszInfoNew.Data.LoginTime,
		Email:              "",
		MemberUid:          yszInfoNew.Data.Userid,
		IsSign:             yszInfoNew.Data.IsSign,
		UserType:           yszInfoNew.Data.UserType,
		SignFlag:           yszInfoNew.Data.SignFlag,
		Redites:            yszInfoNew.Data.Redites,
		ReditesAddOfOneDay: yszInfoNew.Data.ReditesAddOfOneDay,
		GroupId:            yszInfoNew.Data.Groupid,
		MemberToken:        yszInfoNew.Data.Token,
		MessageUserId:      yszInfoNew.Data.MessageUserId,
	}

	response.OkWithDetailed(
		gin.H{
			"UserInfo": returnData,
			"Token":    accessToken,
		},
		"登录成功", c)
	return
}

func (s *SnsApi) GetEncode(c *gin.Context) {
	var req request.SMSValidateCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	encrypt, err := utils.RSAEncrypt([]byte(req.Mobile), []byte(config.RSAPublicBytes))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(encrypt, c)
}

func (s *SnsApi) GetCodeEn(c *gin.Context) {
	var req request.SMSCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	encrData, err := aes.AesCbcEncryptBase64([]byte(req.Code), []byte("szmgszmgszmgszmg"), []byte("0102030405060708"))

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(encrData, c)
}
