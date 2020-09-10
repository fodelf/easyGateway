package v1

import (
	"fmt"
	service "gateway/database"
	InterfaceEntity "gateway/models/InterfaceEntity"
	Utils "gateway/utils"
	"net/http"
	"strconv"

	"gateway/pkg/e"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
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
	if err := service.DB.Find(&serverList).Where("delete_flag =?", 0).Error; err != nil {

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
	var (
		serviceInfo InterfaceEntity.ServiceInfo
	)
	var serverId = c.Query("serverId")
	if err := service.DB.Find(&serviceInfo).Where("service_id =?", serverId).Error; err != nil {
	}
	fmt.Println(serviceInfo)
	appG.Response(http.StatusOK, e.SUCCESS, serviceInfo)
}

// @Tags  服务模块
// @Summary 新增服务
// @Description 新增服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/addService [post]
func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		serviceInfo  InterfaceEntity.ServiceInfo = InterfaceEntity.ServiceInfo{}
		DingdingList string                      = ""
		ServiceRules string                      = ""
		deleteFlag   int                         = 0
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
		ServiceRules = ServiceRules + "," + service["url"].(string) + ":" + service["pathReWrite"].(string)
	}
	servicePort, _ := strconv.Atoi(body["servicePort"].(string))
	serviceLimit, _ := strconv.Atoi(body["serviceLimit"].(string))
	serviceBreak, _ := strconv.Atoi(body["serviceBreak"].(string))
	useConsulPort, _ := strconv.Atoi(body["useConsulPort"].(string))
	useConsulInterval, _ := strconv.Atoi(body["useConsulInterval"].(string))
	useConsulTimeout, _ := strconv.Atoi(body["useConsulTimeout"].(string))
	serviceInfo = InterfaceEntity.ServiceInfo{
		DeleteFlag:          deleteFlag,
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
	tx := service.DB.Begin()
	if err := tx.Create(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
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
	var (
		serviceInfo InterfaceEntity.ServiceInfo
	)
	var body = Utils.GetJsonBody(c)
	fmt.Println(body["serverId"].(string))
	// 数据物理删除
	if err := service.DB.Where("server_id = ?", body["serverId"].(string)).Delete(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
}
