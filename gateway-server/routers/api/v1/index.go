package v1

import (
	"fmt"
	service "gateway/database"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"net/http"
	Time "time"

	"gateway/pkg/e"
	Utils "gateway/utils"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/gin-gonic/gin"
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
	if err := service.DB.Find(&sumInfo).Error; err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, sumInfo)
	}
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
		charts      []InterfaceEntity.ChartInfo
		timeList    []string = []string{}
		totalList   []int64  = []int64{}
		successList []int64  = []int64{}
		failList    []int64  = []int64{}
	)
	if err := service.DB.Limit(7).Order("chart_id DESC").Find(&charts).Error; err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		for i, j := 0, len(charts)-1; i < j; i, j = i+1, j-1 {
			charts[i], charts[j] = charts[j], charts[i]
		}
		for i := 0; i < len(charts); i++ {
			fmt.Println(i)
			// t := reflect.TypeOf(charts[i])
			// immutable := reflect.ValueOf(charts[i])
			total := Utils.GetStructValue(charts[i], "Total")
			success := Utils.GetStructValue(charts[i], "Success")
			fail := Utils.GetStructValue(charts[i], "Fail")
			time := Utils.GetStructValueString(charts[i], "Time")
			totalList = append(totalList, total)
			successList = append(successList, success)
			failList = append(failList, fail)
			timeList = append(timeList, time)
		}
		// if len(charts) < 7 {
		// 	// var ln = len(charts) - 7
		// 	// for i := 0; i < ln; i++ {
		// 	// 	// typeOfChart := reflect.ValueOf(chart)
		// 	// 	// if catType, ok := typeOfChart.FieldByName("time"); ok {
		// 	// 	// 	// fmt.Println(catType.Tag.Get("Time"))
		// 	// 	// }
		// 	// }
		// 	for i := 0; i < len(charts); i++ {
		// 		fmt.Println(i)
		// 		// t := reflect.TypeOf(charts[i])
		// 		// immutable := reflect.ValueOf(charts[i])
		// 		total := Utils.GetStructValue(charts[i], "Total")
		// 		success := Utils.GetStructValue(charts[i], "Success")
		// 		fail := Utils.GetStructValue(charts[i], "Fail")
		// 		time := Utils.GetStructValueString(charts[i], "Time")
		// 		totalList = append(totalList, total)
		// 		successList = append(successList, success)
		// 		failList = append(failList, fail)
		// 		timeList = append(timeList, time)
		// 	}
		// } else {
		// 	for i := 0; i < len(charts); i++ {
		// 		var child = charts[:i]
		// 		fmt.Println(child)
		// 	}
		// }
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"timeList":    timeList,
			"totalList":   totalList,
			"successList": successList,
			"failList":    failList,
		})
	}
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
		total      int64
		fail       int64
		success    int64
		time       string
		percent    int64
		todayState string
	)
	if err := service.DB.Limit(1).Order("chart_id DESC").Find(&charts).Error; err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		total = Utils.GetStructValue(charts[0], "Total")
		success = Utils.GetStructValue(charts[0], "Success")
		fail = Utils.GetStructValue(charts[0], "Fail")
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
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"warningList": [7]string{
			"xx系统于2019年9月1号10时24分异常",
			"ss系统于2019年9月1号10时24分异常",
			"bb系统于2019年9月1号11时24分异常",
			"dd系统于2019年9月1号13时24分异常",
			"ss系统于2019年9月1号15时24分异常",
			"xx系统于2019年9月1号16时24分异常",
			"xx系统于2019年9月1号17时24分异常",
		},
	})
}
