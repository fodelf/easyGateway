package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
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
	// name := c.Query("name")
	// state := -1
	// if arg := c.Query("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// }

	// tagService := tag_service.Tag{
	// 	Name:     name,
	// 	State:    state,
	// 	PageNum:  util.GetPage(c),
	// 	PageSize: setting.AppSetting.PageSize,
	// }
	// tags, err := tagService.GetAll()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	// 	return
	// }

	// count, err := tagService.Count()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	// 	return
	// }

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverSum": 100,
		"warningSum": 100,
		"requestSum": 100,
		"failSum": 100,
	})
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
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"timeList": [7]string{"2010/9/1","2010/9/2","2010/9/3","2010/9/4","2010/9/5","2010/9/6","2010/9/7"},
		"totalList": [7]int{100,100,100,100,100,100,100},
		"successList": [7]int{98,98,98,98,98,98,98},
		"failList": [7]int{2,2,2,2,2,2,2},
	})
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
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"realTime": 1599387778612,
		"todayState": "good",
		"dataSum": map[string]interface{}{
			"total":10000,
			"success":9990,
			"fail":10,
		},
	})
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