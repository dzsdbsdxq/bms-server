package members

import "bms-server/service"

type ApiGroup struct {
	SysMembersApi
}

var (
	sysMembersService = service.ServiceGroupApp.MembersServiceGroup.SysMembersService
)
