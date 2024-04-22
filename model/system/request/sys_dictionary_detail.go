package request

import (
	"bms-server/model/common/request"
	"bms-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
