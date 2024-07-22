package service

import (
	"bms-server/global"
	"bms-server/model/system"
	"bms-server/plugin/sns/model/response"
	service2 "bms-server/service"
	"bms-server/utils"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/wumansgy/goEncrypt/aes"
	"time"
)

type SnsService struct{}

var sysSnsService = service2.ServiceGroupApp.SystemServiceGroup.SnsRecordService

func (e *SnsService) GetValidateCode(mobile string) (*response.SMSValidateCodeResponse, error) {
	var smsValidateCodeResponse response.SMSValidateCodeResponse
	start := time.Now()
	body := map[string]string{
		"appId":      "APP001",
		"actionCode": "5",
		"mobile":     mobile,
		"tenantId":   "ysz",
	}
	info, err := global.GVA_RESTY.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		SetFormData(body).Post("https://uiam.scms.sztv.com.cn:8800/SOAM/oauth/user/v3/getValidateCode")

	end := time.Since(start)
	_ = sysSnsService.CreateSysSnsRecord(system.SysSnsRecord{
		Mobile:       mobile,
		Method:       "getValidateCode",
		Path:         "/SOAM/oauth/user/v3/getValidateCode",
		Status:       info.StatusCode(),
		Latency:      end,
		ErrorMessage: "",
		Body:         utils.Interface2string(body),
		Resp:         string(info.Body()),
	})
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(info.Body(), &smsValidateCodeResponse)
	if err != nil {
		return nil, err
	}

	if smsValidateCodeResponse.ReturnCode != 0 {
		return nil, fmt.Errorf("请求错误，错误原因：%s", smsValidateCodeResponse.ReturnMessage)
	}
	return &smsValidateCodeResponse, nil
}

func (e *SnsService) SendSMS(mobile string, validateCode string) error {
	var smsSendResponse response.SMSSendResponse
	start := time.Now()
	body := map[string]string{
		"appId":        "APP001",
		"actionCode":   "5",
		"mobile":       mobile,
		"validateCode": validateCode,
		"tenantId":     "ysz",
	}

	info, err := global.GVA_RESTY.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		SetFormData(body).Post("https://uiam.scms.sztv.com.cn:8800/SOAM/oauth/user/v3/sendMessage")

	end := time.Since(start)
	_ = sysSnsService.CreateSysSnsRecord(system.SysSnsRecord{
		Mobile:       mobile,
		Method:       "sendMessage",
		Path:         "/SOAM/oauth/user/v3/sendMessage",
		Status:       info.StatusCode(),
		Latency:      end,
		ErrorMessage: "",
		Body:         utils.Interface2string(body),
		Resp:         string(info.Body()),
	})

	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(info.Body(), &smsSendResponse)
	if err != nil {
		return err
	}

	if smsSendResponse.ReturnCode != 0 {
		return fmt.Errorf("请求错误，错误原因：%s", smsSendResponse.ReturnMessage)
	}
	return nil
}
func (e *SnsService) CheckSMSCode(mobile string, code string) (string, string, error) {
	var smsCheckResponse response.SMSCheckResponse
	encrData, err := aes.AesCbcEncryptBase64([]byte(code), []byte("szmgszmgszmgszmg"), []byte("0102030405060708"))

	if err != nil {
		return "", "", err
	}

	start := time.Now()
	body := map[string]string{
		"appId":      "APP001",
		"authType":   "0",
		"mobile":     mobile,
		"mobileCode": encrData,
	}

	info, err := global.GVA_RESTY.R().
		SetResult(&smsCheckResponse).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		SetFormData(body).Post("https://uiam.scms.sztv.com.cn:8800/SOAM/oauth/user/v2_2/authorize")

	end := time.Since(start)
	_ = sysSnsService.CreateSysSnsRecord(system.SysSnsRecord{
		Mobile:       mobile,
		Method:       "authorize",
		Path:         "/SOAM/oauth/user/v2_2/authorize",
		Status:       info.StatusCode(),
		Latency:      end,
		ErrorMessage: "",
		Body:         utils.Interface2string(body),
		Resp:         string(info.Body()),
	})

	if err != nil {
		return "", "", err
	}

	if smsCheckResponse.ReturnCode != 0 {
		return "", "", fmt.Errorf("请求错误，错误原因：%s", smsCheckResponse.ReturnMessage)
	}
	return smsCheckResponse.Data.Uid, smsCheckResponse.Data.AccessToken, nil
}

func (e *SnsService) GetYszUserInfo(accessToken string) (*response.YszUserInfo, error) {
	var yszUserInfo response.YszUserInfo
	var yszUserInfoError response.YszUserInfoError
	start := time.Now()
	body := map[string]string{
		"access_token": accessToken,
	}
	info, err := global.GVA_RESTY.R().
		SetError(&yszUserInfoError).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		Post("https://uiam.scms.sztv.com.cn:8800/SOAM/oauth/user/getuserinfo?access_token=" + accessToken + "&appId=APP001")

	end := time.Since(start)

	if err != nil {
		return nil, err
	}
	if yszUserInfoError.Error == "invalid_token" {
		return nil, fmt.Errorf(yszUserInfoError.ErrorDescription)
	}
	err = jsoniter.Unmarshal(info.Body(), &yszUserInfo)
	if err != nil {
		return nil, err
	}
	_ = sysSnsService.CreateSysSnsRecord(system.SysSnsRecord{
		Mobile:       yszUserInfo.Data.Mobile,
		Method:       "getuserinfo",
		Path:         "/SOAM/oauth/user/getuserinfo",
		Status:       info.StatusCode(),
		Latency:      end,
		ErrorMessage: "",
		Body:         utils.Interface2string(body),
		Resp:         string(info.Body()),
	})
	return &yszUserInfo, nil
}

func (e *SnsService) GetYszMemberInfo(accessToken string) (*response.YszUserInfoNewResponse, error) {
	var yszUserInfo response.YszUserInfoNewResponse
	url := "https://memberrest.mp.sztv.com.cn/memberapi/api/user/detail?token=" + accessToken + "&appId=APP001&tenantid=ysz&uid=-1"
	start := time.Now()
	body := map[string]string{
		"token":    accessToken,
		"appId":    "APP001",
		"tenantid": "ysz",
		"uid":      "1",
	}
	info, err := global.GVA_RESTY.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		Post(url)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(info.Body(), &yszUserInfo)
	if err != nil {
		return nil, err
	}

	if yszUserInfo.Success == false {
		return nil, fmt.Errorf(yszUserInfo.Description)
	}
	end := time.Since(start)
	_ = sysSnsService.CreateSysSnsRecord(system.SysSnsRecord{
		Mobile:       yszUserInfo.Data.Mobile,
		Method:       "userdetail",
		Path:         "/memberapi/api/user/detail",
		Status:       info.StatusCode(),
		Latency:      end,
		ErrorMessage: "",
		Body:         utils.Interface2string(body),
		Resp:         string(info.Body()),
	})
	return &yszUserInfo, nil
}
