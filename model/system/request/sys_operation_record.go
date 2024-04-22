package request

import (
	"bms-server/model/common/request"
	"bms-server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
