package router

import (
	"bms-server/router/example"
	"bms-server/router/integral"
	"bms-server/router/members"
	"bms-server/router/system"
	"bms-server/router/wheel"
	"bms-server/router/wheelAddress"
	"bms-server/router/wheelPrizes"
	"bms-server/router/wheelRecords"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	Members      members.RouterGroup
	Integral     integral.RouterGroup
	Wheel        wheel.RouterGroup
	WheelPrizes  wheelPrizes.RouterGroup
	WheelAddress wheelAddress.RouterGroup
	WheelRecords wheelRecords.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
