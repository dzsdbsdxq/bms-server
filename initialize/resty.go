package initialize

import (
	"bms-server/global"
	"github.com/go-resty/resty/v2"
	"time"
)

func InitRestyClient() {
	global.GVA_RESTY = resty.New()
	global.GVA_RESTY.
		SetTimeout(time.Second*20).
		SetRetryCount(3).
		SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36").
		SetHeader("Content-Type", "application/json")
	//SetHeader("origin", "https://www.sztv.com.cn")
}
