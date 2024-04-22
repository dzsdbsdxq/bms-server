package response

import "bms-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
