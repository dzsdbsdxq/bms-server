package api

import (
	"bms-server/global"
	"bms-server/model/common/response"
	wheelPrizesReq "bms-server/model/wheelPrizes/request"
	"bms-server/model/wheelRecords"
	"bms-server/plugin/wheels/model"
	"bms-server/service"
	"bms-server/service/wheel"
	"bms-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type WheelApi struct{}

var sysWheelsService = service.ServiceGroupApp.WheelServiceGroup.SysWheelsService
var sysWheelsPrizeService = service.ServiceGroupApp.WheelPrizesServiceGroup.SysWheelPrizesService
var sysWheelsRecordsService = service.ServiceGroupApp.WheelRecordsServiceGroup.SysWheelRecordsService

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

// WheelWin Win 抽奖接口
func (w *WheelApi) WheelWin(c *gin.Context) {
	var wwr model.WheelWinRequest
	err := c.ShouldBindJSON(&wwr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 查询该活动信息
	wheels, err := sysWheelsService.GetSysWheels(wwr.WheelId)
	if err != nil {
		response.FailWithMessage("活动不存在", c)
		return
	}
	//判断活动是否下架
	if *wheels.Status == 1 {
		response.FailWithMessage("活动不存在", c)
		return
	}
	//判断活动时间是否开启
	if time.Now().Before(*wheels.StartTime) {
		response.FailWithMessage("活动未开始", c)
		return
	}
	//判断活动时间是否结束
	if time.Now().After(*wheels.EndTime) {
		response.FailWithMessage("活动已结束", c)
		return
	}

	memberId := int(utils.GetMemberID(c))

	records := sysWheelsRecordsService.GetSysWheelRecordsInfoListBy(int(wheels.ID), memberId)

	//判断用户每日可抽次数
	if *wheels.DayGetNum > 0 {
		if records != nil && len(records) >= *wheels.DayGetNum {
			response.FailWithMessage("今日抽奖已达上限", c)
			return
		}
	}
	allowWinPrize := true
	isHits := 0
	// 判断用户是否中奖
	for _, record := range records {
		if *record.IsHit == 1 {
			isHits++
		}
	}
	//整场活动同一用户可中奖次数
	if isHits >= *wheels.GetTotalNum {
		//如果中奖已达上限，返回未抽奖
		allowWinPrize = false
	}
	var aw *wheel.AwardBatchService

	//AlgType = 1 系统自带派奖算法，2=根据创建抽奖概率
	if *wheels.AlgType == 1 {
		aw = wheel.WinPrize(int(wheels.ID), allowWinPrize)
	} else {

	}

	if aw != nil {
		// 保存用户中奖纪录
		_ = sysWheelsRecordsService.CreateSysWheelRecords(&wheelRecords.SysWheelRecords{
			MemberId:     &memberId,
			Name:         aw.PrizeName,
			WheelId:      &aw.WheelId,
			WheelPrizeId: &aw.PrizeId,
			Type:         &aw.PrizeType,
			IsHit:        &aw.IsHit,
		})
		//奖品已抽数量增加
		_ = sysWheelsPrizeService.UpdateSysWheelPrizesField("use_nums", strconv.Itoa(aw.PrizeId))
	}
	response.OkWithDetailed(aw, "success", c)
}
