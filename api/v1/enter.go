package v1

import (
	"bms-server/api/v1/example"
	"bms-server/api/v1/integral"
	"bms-server/api/v1/members"
	"bms-server/api/v1/system"
	"bms-server/api/v1/wheel"
	"bms-server/api/v1/wheelAddress"
	"bms-server/api/v1/wheelPrizes"
	"bms-server/api/v1/wheelRecords"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	MembersApiGroup      members.ApiGroup
	IntegralApiGroup     integral.ApiGroup
	WheelApiGroup        wheel.ApiGroup
	WheelPrizesApiGroup  wheelPrizes.ApiGroup
	WheelAddressApiGroup wheelAddress.ApiGroup
	WheelRecordsApiGroup wheelRecords.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
