// 自动生成模板SysIntegral
package integral

import (
	"bms-server/global"
	"bms-server/model/members"
)

// sysIntegral表 结构体  SysIntegral
type SysIntegral struct {
	global.GVA_MODEL
	MemberId   *int                `json:"memberId" form:"memberId" gorm:"column:member_id;comment:会员id;size:19;foreignKey:MemberId"` //会员id
	ProjectId  *int                `json:"projectId" form:"projectId" gorm:"column:project_id;comment:适用规则ID;size:10;"`               //适用规则ID
	ScoreType  string              `json:"scoreType" form:"scoreType" gorm:"column:score_type;comment:积分类型;size:11;"`                 //积分类型
	Aors       *int                `json:"aors" form:"aors" gorm:"column:aors;comment:积分增或减,1=add,2=del;"`                            //积分增或减,1=add,2=del
	Score      *int                `json:"score" form:"score" gorm:"column:score;comment:分值;size:10;"`                                //分值
	Remark     string              `json:"remark" form:"remark" gorm:"column:remark;comment:备注说明;size:255;"`                          //备注说明
	SysMembers *members.SysMembers `json:"member" gorm:"foreignKey:MemberId"`
}

// TableName sysIntegral表 SysIntegral自定义表名 sys_integral
func (SysIntegral) TableName() string {
	return "sys_integral"
}
