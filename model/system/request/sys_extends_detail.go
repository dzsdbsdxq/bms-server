package request

import (
	"bms-server/model/common/request"
	"bms-server/model/system"
)

type SysExtendsDetailSearch struct {
	system.SysExtendsDetail
	request.PageInfo
}
