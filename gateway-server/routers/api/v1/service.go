package v1

import (
	"fmt"
	service "gateway/database"
	InterfaceEntity "gateway/models/InterfaceEntity"
	Utils "gateway/utils"
	"net/http"
	"strconv"
	"strings"
	proxy "gateway/middleware/proxy"
	"gateway/pkg/e"
	// "io/ioutil"
	"reflect"
	"github.com/fatih/structs"
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
	}else{
		var serviceInfoMap = structs.Map(serviceInfo)
		var serviceRules = serviceInfoMap["ServiceRules"].(string)
		var rules = strings.Split(serviceRules, ",")
		var showRules = []map[string]interface{}{}
		for i := 0; i < len(rules); i++ {
			var ruleArray = strings.Split(rules[i], ";")
			var rule = map[string]interface{}{
				"url":ruleArray[0],
				"pathReWriteBefore":ruleArray[1],
				"pathReWriteUrl":ruleArray[2],
			}
			showRules = append(showRules,rule)
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
			"serverId":           serviceInfoMap["ServerId"],
			"serviceName":        serviceInfoMap["ServiceName"],
			"serviceType":           serviceInfoMap["ServiceType"],
			"serviceAddress":       serviceInfoMap["ServiceAddress"],
			"servicePort":          servicePort,
			"serviceLimit":         serviceLimit,
			"serviceBreak":         serviceBreak,
			"serviceRules":            serviceInfoMap["ServiceRules"],
			"useConsulId":          serviceInfoMap["UseConsulId"],
			"useConsulTag":         serviceInfoMap["UseConsulTag"],
			"useConsulCheckPath":    serviceInfoMap["UseConsulCheckPath"],
			"useConsulPort":         useConsulPort,
			"useConsulInterval":     useConsulInterval,
			"useConsulTimeout":      useConsulTimeout,
			"dingdingAccessToken":   serviceInfoMap["DingdingAccessToken"],
			"dingdingSecret":       serviceInfoMap["DingdingSecret"],
			"dingdingList":          serviceInfoMap["DingdingList"],
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
func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		// serviceInfo       InterfaceEntity.ServiceInfo = InterfaceEntity.ServiceInfo{}
		DingdingList      string                      = ""
		ServiceRules      string                      = ""
		deleteFlag        int                         = 0
		SingleProxyConfig map[string]interface{}      = map[string]interface{}{
			"serviceAddress": "",
			"servicePort":    80,
			"serviceRules":   []map[string]interface{}{},
		}
		servicePort int
		serviceLimit int
		serviceBreak int
		useConsulPort int
		useConsulInterval int
		useConsulTimeout int
		dingdingList []interface{}
		serviceRules []interface{}
	)
	var body = Utils.GetJsonBody(c)
	servicePort = int(body["servicePort"].(float64))
	serviceLimit, _ = strconv.Atoi(body["serviceLimit"].(string))
	serviceBreak, _ = strconv.Atoi(body["serviceBreak"].(string))
	useConsulPort, _ = strconv.Atoi(body["useConsulPort"].(string))
	useConsulInterval, _ = strconv.Atoi(body["useConsulInterval"].(string))
	useConsulTimeout, _ = strconv.Atoi(body["useConsulTimeout"].(string))
	if reflect.TypeOf(body["dingdingList"]) == nil{

	}else{
		dingdingList = body["dingdingList"].([]interface{})
	}
	if reflect.TypeOf(body["serviceRules"]) == nil{

	}else{
		serviceRules = body["serviceRules"].([]interface{})
	}
	for i := 0; i < len((dingdingList)); i++ {
		var dingding = dingdingList[i].(string)
		DingdingList = DingdingList + "," + dingding
	}
	for i := 0; i < len((serviceRules)); i++ {
		var service = serviceRules[i].(map[string]interface{})
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
	// tx := service.DB.Begin()
	if err := service.DB.Create(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Close()
		SingleProxyConfig["serviceAddress"] = body["serviceAddress"].(string)
		// servicePort, _ := strconv.Atoi(body["servicePort"].(string))
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
		DingdingList      string                      = ""
		ServiceRules      string                      = ""
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
	if err := service.DB.Model(&InterfaceEntity.ServiceInfo{}).Where("server_id =" ,body["ServerId"].(string)).Update(&serviceInfo).Error; err != nil {
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
type Service1 struct {
	ServerId     string  `json:"serverId"`
}
func DeleteService(c *gin.Context) {
	appG := app.Gin{C: c}
	// var (
	// 	serviceInfo InterfaceEntity.ServiceInfo
	// )
	// var serverId = c.Request.FormValue("serverId")
	// data, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println("xxxxxxxx")
	// fmt.Println(data)
	var ser Service1
	// fmt.Println(c.)
	// // If `GET`, only `Form` binding engine (`query`) used.
	// // 如果是Get，那么接收不到请求中的Post的数据？？
	// // 如果是Post, 首先判断 `content-type` 的类型 `JSON` or `XML`, 然后使用对应的绑定器获取数据.
	// // See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	// if c.ShouldBind(&ser) == nil {
	// 	fmt.Println(ser)
	// }
	// var body = Utils.GetJsonBody(c)
	c.ShouldBind(&ser)
	fmt.Println(ser.ServerId)
	// 数据物理删除
	// tx := service.DB.Begin()
	if err := service.DB.Where("server_id = ?",ser.ServerId).Delete(&InterfaceEntity.ServiceInfo{}).Error; err != nil {
		// fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
	// tx.Close()
}
