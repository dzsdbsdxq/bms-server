package request

import (
	"bms-server/model/common/request"
	"time"
)

type SysMembersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Username       string     `json:"username" form:"username"`
	Phone          string     `json:"phone" form:"phone"`

	request.PageInfo
}

type ScoreUpdateRequest struct {
	MemberId   uint   `json:"memberId" form:"memberId"`
	OptionType int    `json:"optionType" form:"optionType"`
	ScoreValue int    `json:"scoreValue" form:"scoreValue"`
	ScoreType  string `json:"scoreType" form:"scoreType"`
	Title      string `json:"title" form:"title"`
	Desc       string `json:"desc" form:"desc"`
}
