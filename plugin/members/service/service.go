package service

import (
	"bms-server/global"
	"bms-server/model/members"
	"bms-server/plugin/members/model"
	"bms-server/utils"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type RegisterService struct{}

func (e *RegisterService) RegisterByUserNameAndPassword(req model.Register) (member *members.SysMembers, err error) {
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		// TODO resopne need
		return nil, err
	}
	var store = base64Captcha.DefaultMemStore

	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		return nil, errors.New("验证码错误")
	}
	// 判断两次密码是否正确
	if req.RePassword != req.Password {
		return nil, errors.New("两次密码不正确")
	}
	//.Model(&members.SysMembers{})
	err = global.GVA_DB.Where("username = ?", req.Username).First(&member).Error
	if err == nil {
		return nil, errors.New("用户名已注册")
	}
	platform := 1 //初始化 平台，1 = 系统注册
	enable := 1   //初始化用户状态
	score := 0    //初始化积分

	// 否则 附加uuid 密码hash加密 注册
	u := &members.SysMembers{
		Username: req.Username,
		NickName: "注册用户",
		Password: utils.BcryptHash(req.Password),
		Uuid:     uuid.Must(uuid.NewV4()).String(),
		Platform: &platform,
		Enable:   &enable,
		Score:    &score,
	}

	if err := global.GVA_DB.Create(&u).Error; err != nil {
		return nil, errors.New("注册失败")
	}

	return u, nil
	// 前面的代码 拿不到正确的 user，所以需要再次查询一次

}

func (e *RegisterService) LoginByUserNameAndPassword(req model.LoginRequest) (member *members.SysMembers, err error) {
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		return nil, err
	}

	var store = base64Captcha.DefaultMemStore

	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		return nil, errors.New("验证码错误")
	}
	err = global.GVA_DB.Where("username = ?", req.Username).First(&member).Error
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在!", zap.Error(err))
		return nil, errors.New("用户名或密码错误")
	}

	if ok := utils.BcryptCheck(req.Password, member.Password); !ok {
		global.GVA_LOG.Error("登陆失败! 密码错误!", zap.Error(err))
		return nil, errors.New("用户名或密码错误")
	}
	//正常用户登录标志
	if *member.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁!")
		return nil, errors.New("用户被禁止登录")
	}
	return member, nil

}
