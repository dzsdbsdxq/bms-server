package example

import "bms-server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
	FlushCdnApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
