package v1

import (
	"fmt"
	service "gateway/database"
	proxy "gateway/middleware/proxy"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"gateway/pkg/e"
	Utils "gateway/utils"
	"net/http"
	"strconv"
	"strings"

	// "io/ioutil"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/fatih/structs"
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
	// var tx = service.DB.Begin()
	// tx.Close()
	if err := service.DB.Model(&serviceInfo).Count(&sum).Error; err != nil {
	}
	// tx.Close()
	// var tx1 = service.DB.Begin()
	if err := service.DB.Model(&serviceInfo).Where("service_type =?", "http").Count(&count).Error; err != nil {
	}
	// tx1.Close()
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
	var tx = service.DB.Begin()
	if err := tx.Find(&serverList).Where("delete_flag =?", 0).Error; err != nil {

	} else {
	}
	// tx.Close()
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
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		var serviceInfoMap = structs.Map(serviceInfo)
		var serviceRules = serviceInfoMap["ServiceRules"].(string)
		var rules = strings.Split(serviceRules, ",")
		var showRules = []map[string]interface{}{}
		for i := 0; i < len(rules); i++ {
			var ruleArray = strings.Split(rules[i], ";")
			var rule = map[string]interface{}{
				"url":               ruleArray[0],
				"pathReWriteBefore": ruleArray[1],
				"pathReWriteUrl":    ruleArray[2],
			}
			showRules = append(showRules, rule)
		}
		// fmt.Println(serviceInfo)
		serviceInfoMap["ServiceRules"] = showRules
		var servicePort = serviceInfoMap["ServicePort"]
		if serviceInfoMap["ServicePort"] == 0 {
			servicePort = ""
		}
		var serviceLimit = serviceInfoMap["ServiceLimit"]
		if serviceInfoMap["ServiceLimit"] == 0 {
			serviceLimit = ""
		}
		var serviceBreak = serviceInfoMap["ServiceBreak"]
		if serviceInfoMap["ServiceBreak"] == 0 {
			serviceBreak = ""
		}
		var useConsulPort = serviceInfoMap["UseConsulPort"]
		if serviceInfoMap["UseConsulPort"] == 0 {
			useConsulPort = ""
		}
		var useConsulInterval = serviceInfoMap["UseConsulInterval"]
		if serviceInfoMap["UseConsulInterval"] == 0 {
			useConsulInterval = ""
		}
		var useConsulTimeout = serviceInfoMap["UseConsulTimeout"]
		if serviceInfoMap["UseConsulTimeout"] == 0 {
			useConsulTimeout = ""
		}
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
			"serverId":            serviceInfoMap["ServerId"],
			"serviceName":         serviceInfoMap["ServiceName"],
			"serviceType":         serviceInfoMap["ServiceType"],
			"serviceAddress":      serviceInfoMap["ServiceAddress"],
			"servicePort":         servicePort,
			"serviceLimit":        serviceLimit,
			"serviceBreak":        serviceBreak,
			"serviceRules":        serviceInfoMap["ServiceRules"],
			"useConsulId":         serviceInfoMap["UseConsulId"],
			"useConsulTag":        serviceInfoMap["UseConsulTag"],
			"useConsulCheckPath":  serviceInfoMap["UseConsulCheckPath"],
			"useConsulPort":       useConsulPort,
			"useConsulInterval":   useConsulInterval,
			"useConsulTimeout":    useConsulTimeout,
			"dingdingAccessToken": serviceInfoMap["DingdingAccessToken"],
			"dingdingSecret":      serviceInfoMap["DingdingSecret"],
			"dingdingList":        serviceInfoMap["DingdingList"],
		})
	}

}

// @Tags  服务模块
// @Summary 新增服务
// @Description 新增服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/addService [post]
type ImportServiceBody struct {
	ServiceName         string                   `json:"serviceName"`
	ServiceType         string                   `json:"serviceType"`
	ServiceAddress      string                   `json:"serviceAddress"`
	ServicePort         int                      `json:"servicePort"`
	ServiceLimit        int                      `json:"serviceLimit"`
	ServiceBreak        int                      `json:"serviceBreak"`
	ServiceRules        []map[string]interface{} `json:"serviceRules"`
	UseConsulId         string                   `json:"useConsulId"`
	UseConsulTag        string                   `json:"useConsulTag"`
	UseConsulCheckPath  string                   `json:"useConsulCheckPath"`
	UseConsulPort       int                      `json:"useConsulPort"`
	UseConsulInterval   int                      `json:"useConsulInterval"`
	UseConsulTimeout    int                      `json:"useConsulTimeout"`
	DingdingAccessToken string                   `json:"dingdingAccessToken"`
	DingdingSecret      string                   `json:"dingdingSecret"`
	DingdingList        []string                 `json:"dingdingList"`
}

func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		// serviceInfo       InterfaceEntity.ServiceInfo = InterfaceEntity.ServiceInfo{}
		DingdingList      string
		ServiceRules      string                 = ""
		deleteFlag        int                    = 0
		SingleProxyConfig map[string]interface{} = map[string]interface{}{
			"serviceAddress": "",
			"servicePort":    80,
			"serviceRules":   []map[string]interface{}{},
		}
		servicePort       int
		serviceLimit      int
		serviceBreak      int
		useConsulPort     int
		useConsulInterval int
		useConsulTimeout  int
		dingdingList      []string
		serviceRules      []map[string]interface{}
		importServiceBody ImportServiceBody
	)
	c.ShouldBind(&importServiceBody)
	servicePort = importServiceBody.ServicePort
	serviceLimit = importServiceBody.ServiceLimit
	serviceBreak = importServiceBody.ServiceBreak
	useConsulPort = importServiceBody.UseConsulPort
	useConsulInterval = importServiceBody.UseConsulInterval
	useConsulTimeout = importServiceBody.UseConsulTimeout
	dingdingList = importServiceBody.DingdingList
	serviceRules = importServiceBody.ServiceRules
	for i := 0; i < len((dingdingList)); i++ {
		var dingding = dingdingList[i]
		DingdingList = DingdingList + "," + dingding
	}
	for i := 0; i < len((serviceRules)); i++ {
		var service = serviceRules[i]
		var oldPath = service["pathReWriteBefore"].(string)
		var newPath = service["pathReWriteUrl"].(string)
		if i == 0 {
			ServiceRules = service["url"].(string) + ";" + oldPath + ";" + newPath
		} else {
			ServiceRules = ServiceRules + "," + service["url"].(string) + ";" + oldPath + ";" + newPath
		}
		var rule = map[string]interface{}{
			"url": service["url"].(string),
			"pathReWrite": map[string]interface{}{
				"oldPath": oldPath,
				"newPath": newPath,
			},
		}
		SingleProxyConfig["serviceRules"] = append(SingleProxyConfig["serviceRules"].([]map[string]interface{}), rule)
	}
	serviceInfo := &InterfaceEntity.ServiceInfo{
		DeleteFlag:          deleteFlag,
		ServerId:            Utils.GenerateUUID(),
		ServiceName:         importServiceBody.ServiceName,
		ServiceType:         importServiceBody.ServiceType,
		ServiceAddress:      importServiceBody.ServiceAddress,
		ServicePort:         servicePort,
		ServiceLimit:        serviceLimit,
		ServiceBreak:        serviceBreak,
		ServiceRules:        ServiceRules,
		UseConsulId:         importServiceBody.UseConsulId,
		UseConsulTag:        importServiceBody.UseConsulTag,
		UseConsulCheckPath:  importServiceBody.UseConsulCheckPath,
		UseConsulPort:       useConsulPort,
		UseConsulInterval:   useConsulInterval,
		UseConsulTimeout:    useConsulTimeout,
		DingdingAccessToken: importServiceBody.DingdingAccessToken,
		DingdingSecret:      importServiceBody.DingdingSecret,
		DingdingList:        DingdingList,
	}
	fmt.Println(serviceInfo)
	// tx := service.DB.Begin()
	// service.DB.Open()
	if err := service.DB.Create(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Close()
		SingleProxyConfig["serviceAddress"] = importServiceBody.ServiceAddress
		SingleProxyConfig["servicePort"] = servicePort
		proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		// tx.Commit()
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
	var (
		// serviceInfo       InterfaceEntity.ServiceInfo = InterfaceEntity.ServiceInfo{}
		DingdingList string = ""
		ServiceRules string = ""
		// SingleProxyConfig map[string]interface{}      = map[string]interface{}{
		// 	"serviceAddress": "",
		// 	"servicePort":    80,
		// 	"serviceRules":   []map[string]interface{}{},
		// }
	)
	var body = Utils.GetJsonBody(c)
	servicePort, _ := strconv.Atoi(body["servicePort"].(string))
	serviceLimit, _ := strconv.Atoi(body["serviceLimit"].(string))
	serviceBreak, _ := strconv.Atoi(body["serviceBreak"].(string))
	useConsulPort, _ := strconv.Atoi(body["useConsulPort"].(string))
	useConsulInterval, _ := strconv.Atoi(body["useConsulInterval"].(string))
	useConsulTimeout, _ := strconv.Atoi(body["useConsulTimeout"].(string))
	var dingdingList = body["dingdingList"].([]interface{})
	for i := 0; i < len((dingdingList)); i++ {
		var dingding = dingdingList[i].(string)
		DingdingList = DingdingList + "," + dingding
	}
	var serviceRules = body["serviceRules"].([]interface{})
	for i := 0; i < len((serviceRules)); i++ {
		var service = serviceRules[i].(map[string]interface{})
		var oldPath = service["pathReWriteBefore"].(string)
		var newPath = service["pathReWriteUrl"].(string)
		if i == 0 {
			ServiceRules = service["url"].(string) + ";" + oldPath + ";" + newPath
		} else {
			ServiceRules = ServiceRules + "," + service["url"].(string) + ";" + oldPath + ";" + newPath
		}
		// var rule = map[string]interface{}{
		// 	"url": service["url"].(string),
		// 	"pathReWrite": map[string]interface{}{
		// 		"oldPath": oldPath,
		// 		"newPath": newPath,
		// 	},
		// }
		// SingleProxyConfig["serviceRules"] = append(SingleProxyConfig["serviceRules"].([]map[string]interface{}), rule)
	}
	serviceInfo := &InterfaceEntity.ServiceInfo{
		ServerId:            body["ServerId"].(string),
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
	fmt.Println(serviceInfo)
	service.DB.Close()
	if err := service.DB.Model(&InterfaceEntity.ServiceInfo{}).Where("server_id =", body["ServerId"].(string)).Update(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Close()
		// SingleProxyConfig["serviceAddress"] = body["serviceAddress"].(string)
		// servicePort, _ := strconv.Atoi(body["servicePort"].(string))
		// SingleProxyConfig["servicePort"] = servicePort
		// proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		// tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
}

// @Tags  服务模块
// @Summary 删除服务
// @Description 删除服务
// @Accept  json
// @Produce  json
// @Param serviceId path string true "ServiceId"  删除服务的id
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/deleteService [post]
type DeleteServiceBody struct {
	ServerId string `json:"serverId"`
}

func DeleteService(c *gin.Context) {
	appG := app.Gin{C: c}
	var ser DeleteServiceBody
	c.ShouldBind(&ser)
	// 数据物理删除
	// tx := service.DB.Begin()
	if err := service.DB.Where("server_id = ?", ser.ServerId).Delete(&InterfaceEntity.ServiceInfo{}).Error; err != nil {
		// fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
	// tx.Close()
}
