package members

import (
	"bms-server/global"
	"bms-server/model/common/response"
	"bms-server/model/integral"
	"bms-server/model/members"
	membersReq "bms-server/model/members/request"
	membersRes "bms-server/model/members/response"
	"bms-server/service"
	"bms-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type SysMembersApi struct {
}

// CreateSysMembers 创建sysMembers表
// @Tags SysMembers
// @Summary 创建sysMembers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body members.SysMembers true "创建sysMembers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysMembers/createSysMembers [post]
func (sysMembersApi *SysMembersApi) CreateSysMembers(c *gin.Context) {
	var sysMembers members.SysMembers
	err := c.ShouldBindJSON(&sysMembers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//系统注册platform=1
	platform := 1
	emailVerified := true
	sysMembers.Uuid = uuid.NewV4().String()
	sysMembers.Platform = &platform
	sysMembers.Project = "system"
	sysMembers.Openid = "sys_" + utils.GetUserIdByRandom()
	//系统增加的用户邮件默认验证
	sysMembers.EmailVerified = &emailVerified
	sysMembers.Password = utils.BcryptHash("123456")

	if err := sysMembersService.CreateSysMembers(&sysMembers); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysMembers 删除sysMembers表
// @Tags SysMembers
// @Summary 删除sysMembers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body members.SysMembers true "删除sysMembers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysMembers/deleteSysMembers [delete]
func (sysMembersApi *SysMembersApi) DeleteSysMembers(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysMembersService.DeleteSysMembers(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysMembersByIds 批量删除sysMembers表
// @Tags SysMembers
// @Summary 批量删除sysMembers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysMembers/deleteSysMembersByIds [delete]
func (sysMembersApi *SysMembersApi) DeleteSysMembersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysMembersService.DeleteSysMembersByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysMembers 更新sysMembers表
// @Tags SysMembers
// @Summary 更新sysMembers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body members.SysMembers true "更新sysMembers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysMembers/updateSysMembers [put]
func (sysMembersApi *SysMembersApi) UpdateSysMembers(c *gin.Context) {
	var sysMembers members.SysMembers
	err := c.ShouldBindJSON(&sysMembers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysMembersService.UpdateSysMembers(sysMembers); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysMembers 用id查询sysMembers表
// @Tags SysMembers
// @Summary 用id查询sysMembers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query members.SysMembers true "用id查询sysMembers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysMembers/findSysMembers [get]
func (sysMembersApi *SysMembersApi) FindSysMembers(c *gin.Context) {
	ID := c.Query("ID")
	if resysMembers, err := sysMembersService.GetSysMembers(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysMembers": resysMembers}, c)
	}
}

// GetSysMembersList 分页获取sysMembers表列表
// @Tags SysMembers
// @Summary 分页获取sysMembers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query membersReq.SysMembersSearch true "分页获取sysMembers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysMembers/getSysMembersList [get]
func (sysMembersApi *SysMembersApi) GetSysMembersList(c *gin.Context) {
	var pageInfo membersReq.SysMembersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysMembersService.GetSysMembersInfoList(pageInfo); err != nil {
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

// ResetPassword
// @Tags      SysMembers
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysMembers                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /sysMembers/resetPassword [post]
func (sysMembersApi *SysMembersApi) ResetPassword(c *gin.Context) {
	var sysMembers members.SysMembers
	err := c.ShouldBindJSON(&sysMembers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysMembersService.ResetPassword(sysMembers.ID)
	if err != nil {
		global.GVA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

// ScoreUpdate
// @Tags      SysMembers
// @Summary   积分变更
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysMembers                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "积分变更"
// @Router    /sysMembers/scoreUpdate [post]
func (sysMembersApi *SysMembersApi) ScoreUpdate(c *gin.Context) {
	var scoreReq membersReq.ScoreUpdateRequest
	err := c.ShouldBindJSON(&scoreReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(scoreReq.MemberId)
	fmt.Println(strconv.Itoa(int(scoreReq.MemberId)))
	// 判断用户是否存在
	sysMembers, err := sysMembersService.GetSysMembers(strconv.Itoa(int(scoreReq.MemberId)))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	projectId := 0
	uid := int(sysMembers.ID)
	//如果是积分减少，判断用户是否有足够的积分
	if scoreReq.OptionType == 2 && *sysMembers.Score < scoreReq.ScoreValue {
		response.FailWithMessage("积分不足", c)
		return
	}
	//往积分明细表增加记录
	sysIntegralService := service.ServiceGroupApp.IntegralServiceGroup.SysIntegralService
	err = sysIntegralService.CreateSysIntegral(&integral.SysIntegral{
		MemberId:  &uid,
		ProjectId: &projectId,
		ScoreType: "system",
		Aors:      &scoreReq.OptionType,
		Score:     &scoreReq.ScoreValue,
		Remark:    scoreReq.Desc,
	})
	if err != nil {
		response.FailWithMessage("变更失败", c)
		return
	}
	//修改用户表的总积分
	_ = sysMembersService.UpdateSysMemberFieldIncDec(sysMembers.ID, "score", scoreReq.OptionType == 1, scoreReq.ScoreValue)
	response.OkWithMessage("变更成功", c)
	//积分变更完成

	//if list, total, err := sysMembersService.GetSysMembersInfoList(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     list,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}
}

// TokenNext 登录以后签发jwt
func (sysMembersApi *SysMembersApi) TokenNext(c *gin.Context, member members.SysMembers) {
	j := &utils.JWTMember{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(membersReq.BaseMemberClaims{
		UUID:     member.Uuid,
		ID:       member.ID,
		NickName: member.NickName,
		Username: member.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	utils.SetMemberToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(membersRes.LoginResponse{
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
	return
}
