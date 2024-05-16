package request

import (
	"bms-server/model/common/request"
	"bms-server/model/system"
)

type SysSnsRecordSearch struct {
	system.SysSnsRecord
	request.PageInfo
}
