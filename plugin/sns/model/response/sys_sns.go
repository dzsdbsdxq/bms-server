package response

type SMSValidateCodeResponse struct {
	ReturnCode    uint   `json:"returnCode"`
	RequestID     string `json:"requestID"`
	ReturnMessage string `json:"returnMessage"`
	Data          string `json:"data"`
}

type SMSSendResponse struct {
	ReturnCode    uint   `json:"returnCode"`
	RequestID     string `json:"requestID"`
	ReturnMessage string `json:"returnMessage"`
}
type SMSCheckResponse struct {
	ReturnCode uint   `json:"returnCode"`
	IsNew      string `json:"isNew"`
	Data       struct {
		Uid          string `json:"uid"`
		Scope        string `json:"scope"`
		ExpiresIn    uint   `json:"expires_in"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	} `json:"data"`
	ReturnMessage string `json:"returnMessage"`
}

type YszUserInfo struct {
	ReturnCode    uint   `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	Data          struct {
		Uid             string `json:"uid"`
		Sex             string `json:"sex"`
		PostCode        string `json:"postCode"`
		ProfileImageUrl string `json:"profileImageUrl"`
		Country         string `json:"country"`
		City            string `json:"city"`
		Address         string `json:"address"`
		NickName        string `json:"nickName"`
		Email           string `json:"email"`
		Name            string `json:"name"`
		Province        string `json:"province"`
		District        string `json:"district"`
		Mobile          string `json:"mobile"`
	}
}
type YszUserInfoNewResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Data        struct {
		Userid                uint          `json:"userid"`
		Username              string        `json:"username"`
		Intro                 string        `json:"intro"`
		HavePassword          uint          `json:"havePassword"`
		Nickname              string        `json:"nickname"`
		Mobile                string        `json:"mobile"`
		Avatar                string        `json:"avatar"`
		BackgroundUrl         string        `json:"backgroundUrl"`
		MergerName            string        `json:"mergerName"`
		AreaId                string        `json:"areaId"`
		AreaCityId            string        `json:"areaCityId"`
		AreaProvinceId        string        `json:"areaProvinceId"`
		Userauth              string        `json:"userauth"`
		Groupid               interface{}   `json:"groupid"`
		Token                 string        `json:"token"`
		Emptype               string        `json:"emptype"`
		Sex                   int           `json:"sex"`
		Tag                   string        `json:"tag"`
		Birthday              string        `json:"birthday"`
		Studyintegral         string        `json:"studyintegral"`
		GroupName             uint          `json:"groupName"`
		GroupLogo             string        `json:"groupLogo"`
		GroupIntegralEnd      uint          `json:"groupIntegralEnd"`
		GroupIntegralStart    uint          `json:"groupIntegralStart"`
		Redites               uint          `json:"redites"`
		UserLevel             string        `json:"userLevel"`
		RealName              string        `json:"realName"`
		UnionId               string        `json:"unionId"`
		InviterCount          uint          `json:"inviterCount"`
		Partystatus           uint          `json:"partystatus"`
		SignFlag              bool          `json:"signFlag"`
		ReditesAddOfOneDay    uint          `json:"reditesAddOfOneDay"`
		InviterinvitationCode string        `json:"inviterinvitationCode"`
		InvitationCode        string        `json:"invitationCode"`
		LoginTime             string        `json:"loginTime"`
		MessageUserId         string        `json:"messageUserId"`
		Regtime               string        `json:"regtime"`
		UnreadCount           uint          `json:"unreadCount"`
		MemberOther           []interface{} `json:"member_other"`
		UserType              uint          `json:"userType"`
		AuthorInfo            string        `json:"authorInfo"`
		IdName                string        `json:"idName"`
		IdCard                string        `json:"idCard"`
		MUid                  uint          `json:"m_uid"`
		IsSign                string        `json:"isSign"`
	} `json:"data"`
	Success bool `json:"success"`
}

// YseUserInfoResponse 返回用户信息
type YseUserInfoResponse struct {
	MemberUid          uint        `json:"memberUid"`
	IsSign             string      `json:"isSign"`
	UserType           uint        `json:"userType"`
	SignFlag           bool        `json:"signFlag"`
	Redites            uint        `json:"redites"`
	ReditesAddOfOneDay uint        `json:"reditesAddOfOneDay"`
	GroupId            interface{} `json:"groupid"`
	MemberToken        string      `json:"memberToken"`
	MessageUserId      string      `json:"messageUserId"`
}
type YseUserInfoNewResponse struct {
	ID                 uint
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
	NickName           string      `json:"nick_name"`
	Openid             string      `json:"openid"`
	Avatar             string      `json:"avatar"`
	IsDel              uint        `json:"is_del"`
	Platform           uint        `json:"platform"`
	City               string      `json:"city"`
	Province           string      `json:"province"`
	Country            string      `json:"country"`
	Mobile             string      `json:"mobile"`
	EszUid             string      `json:"esz_uid"`
	Partment           string      `json:"partment"`
	LastLogin          string      `json:"last_login"`
	Email              string      `json:"email"`
	MemberUid          uint        `json:"memberUid"`
	IsSign             string      `json:"isSign"`
	UserType           uint        `json:"userType"`
	SignFlag           bool        `json:"signFlag"`
	Redites            uint        `json:"redites"`
	ReditesAddOfOneDay uint        `json:"reditesAddOfOneDay"`
	GroupId            interface{} `json:"groupid"`
	MemberToken        string      `json:"memberToken"`
	MessageUserId      string      `json:"messageUserId"`
}

type YszUserInfoError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
