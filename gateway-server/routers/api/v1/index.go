package v1

import (
	"fmt"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"gateway/pkg/e"
	Utils "gateway/utils"
	"net/http"
	"time"
	Time "time"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}

// @Tags  首页模块
// @Summary 查询首页汇总信息
// @Description 查询首页汇总信息
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/index/sum [get]
func GetSum(c *gin.Context) {
	appG := app.Gin{C: c}
	var sumInfo InterfaceEntity.SumInfo
	// service.DB.Begin()
	// var tx = service.DB.Begin()
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
	if err := DB.Find(&sumInfo).Error; err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		// tx.Rollback()
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, sumInfo)
	}
	// tx.Commit()
	DB.Close()
}

// @Tags  首页模块
// @Summary 查询图表数据详情
// @Description 查询图表数据详情
// @Accept  json
// @Produce  json
// @Param id path string false "ID"  查询图表信息id
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/index/charts/{id} [get]
func GetCharts(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		charts          []InterfaceEntity.ChartInfo
		timeList        []string = []string{}
		totalList       []int    = []int{}
		successList     []int    = []int{}
		failList        []int    = []int{}
		tempTimeList    []string = []string{}
		tempTotalList   []int    = []int{}
		tempSuccessList []int    = []int{}
		tempFailList    []int    = []int{}
	)
	// var tx = service.DB.Begin()
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	id := c.Query("type")
	if err := DB.Limit(7).Order("chart_id DESC").Where("server_id = ?", id).Find(&charts).Error; err != nil {
		// tx.Rollback()
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		for i, j := 0, len(charts)-1; i < j; i, j = i+1, j-1 {
			charts[i], charts[j] = charts[j], charts[i]
		}
		fmt.Println(charts)
		for i := 0; i < len(charts); i++ {
			// t := reflect.TypeOf(charts[i])
			// immutable := reflect.ValueOf(charts[i])
			var chartInfo = structs.Map(charts[i])
			var total = chartInfo["Total"].(int)
			var success = chartInfo["Success"].(int)
			var fail = chartInfo["Fail"].(int)
			var time = chartInfo["Time"].(string)
			totalList = append(totalList, total)
			successList = append(successList, success)
			failList = append(failList, fail)
			timeList = append(timeList, time)
		}
		var dis = 7 - len(totalList)
		for k := 0; k < dis; k++ {
			var nTime = time.Now()
			var old = nTime.AddDate(0, 0, (k - dis - 1)).Format("2006/01/02")
			tempTimeList = append(tempTimeList, old)
			tempTotalList = append(tempTotalList, 0)
			tempSuccessList = append(tempSuccessList, 0)
			tempFailList = append(tempFailList, 0)
		}
		totalList = append(tempTotalList, totalList...)
		timeList = append(tempTimeList, timeList...)
		successList = append(tempSuccessList, successList...)
		failList = append(tempFailList, failList...)
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"timeList":    timeList,
			"totalList":   totalList,
			"successList": successList,
			"failList":    failList,
		})
	}
	// tx.Commit()
	DB.Close()
}

// @Tags  首页模块
// @Summary 查询今日实时数据
// @Description 查询今日实时数据
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/index/actualTime [get]
func GetActualTime(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		charts     []InterfaceEntity.ChartInfo
		total      int
		fail       int
		success    int
		time       string
		percent    int
		todayState string
	)
	// var tx = service.DB
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	// tx.Close()
	// defer tx.Commit()
	if err := DB.Where("server_id = ? AND time = ?", "all", Time.Now().Format("2006/01/02")).Find(&charts).Error; err != nil {
		// tx.Rollback()
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		// tx.Close()
		var chartInfo = structs.Map(charts[0])
		total = chartInfo["Total"].(int)
		success = chartInfo["Success"].(int)
		fail = chartInfo["Fail"].(int)
		time = Time.Now().Format("2006/01/02")
		if total != 0 {
			percent = (fail / total)
		} else {
			percent = 0
		}

		// percent := decimal.NewFromFloat(float64(fail)).Div(decimal.NewFromFloat(float64(total)))
		// fmt.Println(percent)
		if float64(percent) < 0.1 {
			todayState = "good"
		} else if float64(percent) < 0.6 {
			todayState = "normal"
		} else {
			todayState = "bad"
		}
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"realTime":   time,
			"todayState": todayState,
			"dataSum": map[string]interface{}{
				"total":   total,
				"success": success,
				"fail":    fail,
			},
		})
	}
	// tx.Commit()
	DB.Close()
}

// @Tags  首页模块
// @Summary 查询最近7条告警数据
// @Description 查询最近7条告警数据
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/index/warningList [get]
func GetWarningList(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		warnings   []InterfaceEntity.WarningInfo
		resultList []string = []string{}
	)
	// var tx = service.DB.Begin()
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	if err := DB.Limit(7).Order("warning_id DESC").Find(&warnings).Error; err != nil {
		// tx.Rollback()
		// appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"warningList": resultList,
		})
	} else {
		// tx.Commit()
		for i, j := 0, len(warnings)-1; i < j; i, j = i+1, j-1 {
			warnings[i], warnings[j] = warnings[j], warnings[i]
		}
		for i := 0; i < len(warnings); i++ {
			var warning = warnings[i]
			time := Utils.GetStructValueString(warning, "Time")
			system := Utils.GetStructValueString(warning, "System")
			var str = system + "系统于" + time + "异常"
			resultList = append(resultList, str)
		}
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"warningList": resultList,
		})
	}
	DB.Close()
	// defer tx.Close()
}
