package service

type ServiceGroup struct {
	ExtendsService
	ExtendsDetailService
}

var ServiceGroupApp = new(ServiceGroup)
