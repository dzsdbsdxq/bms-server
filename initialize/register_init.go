package initialize

import (
	_ "bms-server/source/example"
	_ "bms-server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
