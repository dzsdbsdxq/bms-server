package example

import (
	"bms-server/global"
	"bms-server/model/common/request"
	"bms-server/model/common/response"
	"github.com/gin-gonic/gin"
)

type FlushCdnApi struct{}

func (e *FlushCdnApi) CreateFlushCdn(c *gin.Context) {
	var flushCdn request.FlushCdnRequest
	err := c.ShouldBindJSON(&flushCdn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	url := "https://testing.mp.sztv.com.cn/flushp?url=" + flushCdn.Url
	info, _ := global.GVA_RESTY.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
		Get(url)
	if string(info.Body()) != "success" {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("刷新成功", c)
}
