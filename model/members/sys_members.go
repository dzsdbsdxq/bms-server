// 自动生成模板SysMembers
package members

import (
	"bms-server/global"
	"time"
)

// sysMembers表 结构体  SysMembers
type SysMembers struct {
	global.GVA_MODEL
	Uuid          string     `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:191;"`                    //用户UUID
	Username      string     `json:"username" form:"username" gorm:"column:username;comment:用户名;size:191;"`           //用户名
	Password      string     `json:"-" gorm:"column:password;comment:用户密码;size:191;"`                                 //用户密码
	NickName      string     `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`         //用户昵称
	Avatar        string     `json:"avatar" form:"avatar" gorm:"column:avatar;comment:用户头像;size:255;"`                //用户头像
	Gender        *int       `json:"gender" form:"gender" gorm:"column:gender;comment:性别(1=男，2=女);"`                  //性别(1=男，2=女)
	Age           *int       `json:"age" form:"age" gorm:"column:age;comment:年龄;"`                                    //年龄
	Birthday      *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;"`                     //生日
	BgImage       string     `json:"bgImage" form:"bgImage" gorm:"column:bg_image;comment:背景图;"`                      //背景图
	Openid        string     `json:"openid" form:"openid" gorm:"column:openid;comment:openid;size:255;"`              //openid
	Platform      *int       `json:"platform" form:"platform" gorm:"column:platform;comment:用户注册来源;size:10;"`         //用户注册来源
	Project       string     `json:"project" form:"project" gorm:"column:project;comment:用户项目来源;size:191;"`           //用户项目来源
	Score         *int       `json:"score" form:"score" gorm:"column:score;comment:积分;size:10;"`                      //积分
	Phone         string     `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`                  //用户手机号
	Email         string     `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`                   //用户邮箱
	EmailVerified *bool      `json:"emailVerified" form:"emailVerified" gorm:"column:email_verified;comment:邮箱是否验证;"` //邮箱是否验证
	Enable        *int       `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否被冻结 1正常 2冻结;size:19;"`      //用户是否被冻结 1正常 2冻结
}

// TableName sysMembers表 SysMembers自定义表名 sys_members
func (SysMembers) TableName() string {
	return "sys_members"
}
