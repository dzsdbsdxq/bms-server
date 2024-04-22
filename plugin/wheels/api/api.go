package api

import (
	"bms-server/global"
	wheelPrizesReq "bms-server/model/wheelPrizes/request"
	"bms-server/service"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"text/template"
)

type WheelApi struct{}

var sysWheelsService = service.ServiceGroupApp.WheelServiceGroup.SysWheelsService
var sysWheelsPrizeService = service.ServiceGroupApp.WheelPrizesServiceGroup.SysWheelPrizesService

func (w *WheelApi) IndexPage(c *gin.Context) {

	//获取路径上的UUID
	id := c.Query("id")

	//通过ID去查询数据是否存在此条记录
	wheels, err := sysWheelsService.GetSysWheelsByUUID(id)
	fmt.Println(wheels, err)
	if err != nil {
		return
	}
	//获取奖项
	prizes, _, err := sysWheelsPrizeService.GetSysWheelPrizesInfoList(wheelPrizesReq.SysWheelPrizesSearch{
		ID: wheels.ID,
	})
	if err != nil {
		return
	}
	tmplPath := fmt.Sprintf("%s/wheels/%s/index.html", global.GVA_CONFIG.System.WheelsTemplateDir, wheels.Template)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var pageParams = make(map[string]interface{}, 0)
	err = jsoniter.Unmarshal([]byte(wheels.PageParams), &pageParams)
	if err != nil {
		return
	}
	fmt.Println(pageParams)
	fmt.Println("Lucky:", pageParams["lucky"])
	luckConfig, err := jsoniter.Marshal(pageParams["lucky"])
	if err != nil {
		return
	}
	err = tmpl.Execute(c.Writer, gin.H{
		"wheels":     wheels,
		"luckConfig": string(luckConfig),
		"page":       pageParams,
		"prizes":     prizes,
	})
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	//c.HTML(http.StatusOK, wheels.Template+".html", gin.H{
	//	"wheels":   wheels,
	//	"pageEdit": wheels.PageParams,
	//	"prizes":   prizes,
	//})
}
