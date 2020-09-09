package v1

import (
	service "gateway/database"
	InterfaceEntity "gateway/models/InterfaceEntity"
	Utils "gateway/utils"
	"net/http"
	"strconv"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Tags  服务模块
// @Summary 服务数据汇总
// @Description 查询服务数据汇总
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceSum [get]
func GetServerSum(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		serviceInfo InterfaceEntity.ServiceInfo
		sum         int                      = 0
		count       int                      = 0
		serverList  []map[string]interface{} = []map[string]interface{}{}
	)
	tx := service.DB.Begin()
	if err := tx.Model(&serviceInfo).Count(&sum).Error; err != nil {
	}
	if err := tx.Model(&serviceInfo).Where("service_type =?", "http").Count(&count).Error; err != nil {
	}
	var serviceInterface = map[string]interface{}{
		"label": "http",
		"count": count,
		"value": "http",
	}
	serverList = append(serverList, serviceInterface)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverList": serverList,
		"sum":        sum,
	})
}

type Rule struct {
	url         string
	pathReWrite string
}

// @Tags  服务模块
// @Summary 服务列表
// @Description 查询服务列表
// @Accept  json
// @Produce  json
// @Param type path string false "Type"  查询服务类型
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceList [get]
func GetServerList(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		serverList []InterfaceEntity.ServiceInfo = []InterfaceEntity.ServiceInfo{}
	)
	if err := service.DB.Find(&serverList).Error; err != nil {

	} else {
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverList": serverList,
	})
}

// @Tags  服务模块
// @Summary 服务详情
// @Description 查询服务详情
// @Accept  json
// @Produce  json
// @Param type path string false "Type"  查询服务类型
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceDetail/{id} [get]
func GetServerDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverId":       "xxxxxx",
		"serviceName":    "XX",
		"serviceType":    "http",
		"serviceAddress": "127.0.0.1",
		"servicePort":    "300",
		"serviceLimt":    200,
		"serviceBreak":   30,
		"serviceRules": [1]map[string]interface{}{
			{
				"url":         "/api",
				"pathReWrite": `{/api:''}`,
			},
		},
		"useConsulId":         "xxxxx",
		"useConsulTag":        "ss",
		"useConsulCheckPath":  "/checkHealth",
		"useConsulPort":       9990,
		"useConsulInterval":   10,
		"useConsulTimeout":    3,
		"dingdingAccessToken": "xxxxxxxxxx",
		"dingdingSercet":      "xxxxxxxxxx",
		"dingdingList":        []int{18651892475},
	})
}

// @Tags  服务模块
// @Summary 新增服务
// @Description 新增服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/addService [post]
//汇总实体类
func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		serviceInfo  InterfaceEntity.ServiceInfo = InterfaceEntity.ServiceInfo{}
		DingdingList string                      = ""
		ServiceRules string                      = ""
	)
	var body = Utils.GetJsonBody(c)
	var dingdingList = body["dingdingList"].([]interface{})
	for i := 0; i < len((dingdingList)); i++ {
		var dingding = dingdingList[i].(string)
		DingdingList = DingdingList + "," + dingding
	}
	var serviceRules = body["serviceRules"].([]interface{})
	for i := 0; i < len((serviceRules)); i++ {
		var service = serviceRules[i].(map[string]interface{})
		ServiceRules = ServiceRules + "," + service["url"].(string) + "@" + service["pathReWrite"].(string)
	}
	servicePort, _ := strconv.Atoi(body["servicePort"].(string))
	serviceLimit, _ := strconv.Atoi(body["serviceLimit"].(string))
	serviceBreak, _ := strconv.Atoi(body["serviceBreak"].(string))
	useConsulPort, _ := strconv.Atoi(body["useConsulPort"].(string))
	useConsulInterval, _ := strconv.Atoi(body["useConsulInterval"].(string))
	useConsulTimeout, _ := strconv.Atoi(body["useConsulTimeout"].(string))
	serviceInfo = InterfaceEntity.ServiceInfo{
		ServerId:            Utils.GenerateUUID(),
		ServiceName:         body["serviceName"].(string),
		ServiceType:         body["serviceType"].(string),
		ServiceAddress:      body["serviceAddress"].(string),
		ServicePort:         servicePort,
		ServiceLimit:        serviceLimit,
		ServiceBreak:        serviceBreak,
		ServiceRules:        ServiceRules,
		UseConsulId:         body["useConsulId"].(string),
		UseConsulTag:        body["useConsulTag"].(string),
		UseConsulCheckPath:  body["useConsulCheckPath"].(string),
		UseConsulPort:       useConsulPort,
		UseConsulInterval:   useConsulInterval,
		UseConsulTimeout:    useConsulTimeout,
		DingdingAccessToken: body["dingdingAccessToken"].(string),
		DingdingSecret:      body["dingdingSecret"].(string),
		DingdingList:        DingdingList,
	}
	if err := service.DB.Create(&serviceInfo).Error; err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
}

// @Tags  服务模块
// @Summary 编辑服务
// @Description 编辑服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/editService [post]
func EditService(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
}

// @Tags  服务模块
// @Summary 删除服务
// @Description 删除服务
// @Accept  json
// @Produce  json
// @Param serviceId path string true "ServiceId"  删除服务的id
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/deleteService [post]
func DeleteService(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
}
