package response

import "bms-server/model/system"

type ExtendDetailResponse struct {
	system.SysExtendsDetail
	Jsx interface{} `json:"jsx"`
}
