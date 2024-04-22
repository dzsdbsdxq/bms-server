// 自动生成模板SysWheels
package wheel

import (
	"bms-server/global"
	"time"
	
)

// sysWheels表 结构体  SysWheels
type SysWheels struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:大转盘名称;size:255;"`  //大转盘名称 
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:唯一UUID;size:191;"`  //唯一UUID 
      NeedScore  *int `json:"needScore" form:"needScore" gorm:"column:need_score;comment:所需积分;size:10;"`  //所需积分 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;"`  //状态 
      Template  string `json:"template" form:"template" gorm:"column:template;comment:模版目录名，空表示自定义模版;size:255;"`  //模版目录名，空表示自定义模版 
      DayGetNum  *int `json:"dayGetNum" form:"dayGetNum" gorm:"column:day_get_num;comment:用户每日可抽次数，0表示不限次数;size:10;"`  //用户每日可抽次数，0表示不限次数 
      AlgType  *int `json:"algType" form:"algType" gorm:"column:alg_type;comment:抽奖算法（1=随机概率，根据奖品数量随机，2=根据自定义奖品概率）;size:10;"`  //抽奖算法（1=随机概率，根据奖品数量随机，2=根据自定义奖品概率） 
      ShareTitle  string `json:"shareTitle" form:"shareTitle" gorm:"column:share_title;comment:分享标题;size:255;"`  //分享标题 
      ShareDesc  string `json:"shareDesc" form:"shareDesc" gorm:"column:share_desc;comment:分享标题;size:255;"`  //分享标题 
      ShareImg  string `json:"shareImg" form:"shareImg" gorm:"column:share_img;comment:分享标题;size:255;"`  //分享标题 
      PageParams  string `json:"pageParams" form:"pageParams" gorm:"column:page_params;comment:页面元素;"`  //页面元素 
      StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;"`  //开始时间 
      EndTime  *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;"`  //结束时间 
      VoiceUrl  string `json:"voiceUrl" form:"voiceUrl" gorm:"column:voice_url;comment:音频地址;size:255;"`  //音频地址 
      ThumbUrl  string `json:"thumbUrl" form:"thumbUrl" gorm:"column:thumb_url;comment:转盘缩略图;size:255;"`  //转盘缩略图 
}


// TableName sysWheels表 SysWheels自定义表名 sys_wheels
func (SysWheels) TableName() string {
  return "sys_wheels"
}

