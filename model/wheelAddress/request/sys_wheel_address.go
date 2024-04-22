package request

import (
	"bms-server/model/common/request"
	"time"
	
)

type SysWheelAddressSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
    request.PageInfo
}
