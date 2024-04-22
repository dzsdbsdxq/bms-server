package request

import (
	"bms-server/model/common/request"
	"time"
)

type SysWheelPrizesSearch struct {
	ID             uint       `json:"id" form:"id"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
