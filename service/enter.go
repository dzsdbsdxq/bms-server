package service

import (
	"bms-server/service/example"
	"bms-server/service/integral"
	"bms-server/service/members"
	"bms-server/service/system"
	"bms-server/service/wheel"
	"bms-server/service/wheelAddress"
	"bms-server/service/wheelPrizes"
	"bms-server/service/wheelRecords"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	MembersServiceGroup      members.ServiceGroup
	IntegralServiceGroup     integral.ServiceGroup
	WheelServiceGroup        wheel.ServiceGroup
	WheelPrizesServiceGroup  wheelPrizes.ServiceGroup
	WheelAddressServiceGroup wheelAddress.ServiceGroup
	WheelRecordsServiceGroup wheelRecords.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
