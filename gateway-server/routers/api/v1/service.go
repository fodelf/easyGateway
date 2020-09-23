package v1

import (
	"fmt"
	proxy "gateway/middleware/proxy"
	InterfaceEntity "gateway/models/InterfaceEntity"
	Pkg "gateway/pkg/consul"
	"gateway/pkg/e"
	Utils "gateway/utils"
	"net/http"
	"strings"
	"time"

	// "io/ioutil"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	// var tx = service.DB.Begin()
	// tx.Close()
	if err := DB.Model(&serviceInfo).Count(&sum).Error; err != nil {
	}
	// tx.Close()
	// var tx1 = service.DB.Begin()
	if err := DB.Model(&serviceInfo).Where("service_type =?", "http").Count(&count).Error; err != nil {
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
	DB.Close()
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
	// var tx = service.DB.Begin()
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	if err := DB.Find(&serverList).Where("delete_flag =?", 0).Error; err != nil {

	} else {
	}
	// tx.Close()
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverList": serverList,
	})

	DB.Close()
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
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	if err := DB.Find(&serviceInfo).Where("service_id =?", serverId).Error; err != nil {
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		var serviceInfoMap = structs.Map(serviceInfo)
		var serviceRules = serviceInfoMap["ServiceRules"].(string)
		var rules = strings.Split(serviceRules, ",")
		var showRules = []map[string]interface{}{}
		for i := 0; i < len(rules); i++ {
			var ruleArray = strings.Split(rules[i], ";")
			var rule = map[string]interface{}{
				"url":               "",
				"pathReWriteBefore": "",
				"pathReWriteUrl":    "",
			}
			if len(ruleArray) == 3 {
				rule = map[string]interface{}{
					"url":               ruleArray[0],
					"pathReWriteBefore": ruleArray[1],
					"pathReWriteUrl":    ruleArray[2],
				}
			} else if len(ruleArray) == 2 {
				rule = map[string]interface{}{
					"url":               ruleArray[0],
					"pathReWriteBefore": ruleArray[1],
					"pathReWriteUrl":    "",
				}
			} else if len(ruleArray) == 1 {
				rule = map[string]interface{}{
					"url":               ruleArray[0],
					"pathReWriteBefore": "",
					"pathReWriteUrl":    "",
				}
			}
			showRules = append(showRules, rule)
		}
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
		var dingdingList = []string{}
		if serviceInfoMap["DingdingList"].(string) != "" {
			dingdingList = strings.Split(serviceInfoMap["DingdingList"].(string), ",")
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
			"dingdingList":        dingdingList,
		})
		DB.Close()
	}
	// fmt.Println(serviceInfo
}

// @Tags  服务模块
// @Summary 新增服务
// @Description 新增服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/addService [post]
// type ImportServiceBody struct {
// 	ServerId            string                   `json:"serverId"`
// 	ServiceName         string                   `json:"serviceName"`
// 	ServiceType         string                   `json:"serviceType"`
// 	ServiceAddress      string                   `json:"serviceAddress"`
// 	ServicePort         int                      `json:"servicePort"`
// 	ServiceLimit        int                      `json:"serviceLimit"`
// 	ServiceBreak        int                      `json:"serviceBreak"`
// 	ServiceRules        []map[string]interface{} `json:"serviceRules"`
// 	UseConsulId         string                   `json:"useConsulId"`
// 	UseConsulTag        string                   `json:"useConsulTag"`
// 	UseConsulCheckPath  string                   `json:"useConsulCheckPath"`
// 	UseConsulPort       int                      `json:"useConsulPort"`
// 	UseConsulInterval   int                      `json:"useConsulInterval"`
// 	UseConsulTimeout    int                      `json:"useConsulTimeout"`
// 	DingdingAccessToken string                   `json:"dingdingAccessToken"`
// 	DingdingSecret      string                   `json:"dingdingSecret"`
// 	DingdingList        []string                 `json:"dingdingList"`
// }

func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		serviceInfoCount  InterfaceEntity.ServiceInfo
		DingdingList      string                 = ""
		ServiceRules      string                 = ""
		deleteFlag        int                    = 0
		SingleProxyConfig map[string]interface{} = map[string]interface{}{
			"serviceAddress": "",
			"servicePort":    80,
			"serviceRules":   []map[string]interface{}{},
			"serverId":       "",
			"total":          0,
			"success":        0,
			"fail":           0,
		}
		servicePort       int
		serviceLimit      int
		serviceBreak      int
		useConsulPort     int
		useConsulInterval int
		useConsulTimeout  int
		dingdingList      []string
		serviceRules      []map[string]interface{}
		importServiceBody InterfaceEntity.ImportServiceBody
		sum               int
		sumInfo           InterfaceEntity.SumInfo
		consulInfo        InterfaceEntity.ConsulInfo
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
		if i == 0 {
		} else {
			DingdingList = DingdingList + "," + dingding
		}
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
	var generateUUID = Utils.GenerateUUID()
	serviceInfo := &InterfaceEntity.ServiceInfo{
		DeleteFlag:          deleteFlag,
		ServerId:            generateUUID,
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
		WarnTime:            time.Now().Format("2006/01/02 15:04:05"),
	}
	// fmt.Println(serviceInfo)
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
	if err := DB.First(&consulInfo).Error; err != nil {
	}
	var consulInfoObj = structs.Map(consulInfo)
	if consulInfoObj["ConsulAddress"].(string) != "" {
		Pkg.RegisterServer(importServiceBody)
	}
	if err := DB.Create(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		// tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Close()
		SingleProxyConfig["serviceAddress"] = importServiceBody.ServiceAddress
		SingleProxyConfig["servicePort"] = servicePort
		SingleProxyConfig["serverId"] = generateUUID
		proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		// tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
	if err := DB.Model(&serviceInfoCount).Count(&sum).Error; err != nil {
	}
	if err := DB.First(&sumInfo).Update("server_sum", sum).Error; err != nil {
		// DB.Rollback()
	} else {
		// DB.Commit()
	}
	chartInfo := InterfaceEntity.ChartInfo{
		Time:     time.Now().Format("2006/01/02"),
		Total:    0,
		Success:  0,
		Fail:     0,
		ServerId: generateUUID,
	}
	DB.Create(&chartInfo)
	DB.Close()
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
		DingdingList      string = ""
		ServiceRules      string = ""
		importServiceBody InterfaceEntity.ImportServiceBody
		// SingleProxyConfig map[string]interface{}      = map[string]interface{}{
		// 	"serviceAddress": "",
		// 	"servicePort":    80,
		// 	"serviceRules":   []map[string]interface{}{},
		// }
		servicePort       int
		serviceLimit      int
		serviceBreak      int
		useConsulPort     int
		useConsulInterval int
		useConsulTimeout  int
	)
	c.ShouldBind(&importServiceBody)
	servicePort = importServiceBody.ServicePort
	serviceLimit = importServiceBody.ServiceLimit
	serviceBreak = importServiceBody.ServiceBreak
	useConsulPort = importServiceBody.UseConsulPort
	useConsulInterval = importServiceBody.UseConsulInterval
	useConsulTimeout = importServiceBody.UseConsulTimeout
	var dingdingList = importServiceBody.DingdingList
	for i := 0; i < len((dingdingList)); i++ {
		var dingding = dingdingList[i]
		if i == 0 {
			DingdingList = dingding
		} else {
			DingdingList = DingdingList + "," + dingding
		}
	}
	var serviceRules = importServiceBody.ServiceRules
	for i := 0; i < len((serviceRules)); i++ {
		var service = serviceRules[i]
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
		ServerId:            importServiceBody.ServerId,
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
	fmt.Println(importServiceBody.ServerId)
	// service.DB.Close()
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
	tx := DB.Begin()
	if err := tx.Model(&InterfaceEntity.ServiceInfo{}).Where("server_id = ?", importServiceBody.ServerId).Update(&serviceInfo).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		// tx.Close()
		// SingleProxyConfig["serviceAddress"] = body["serviceAddress"].(string)
		// servicePort, _ := strconv.Atoi(body["servicePort"].(string))
		// SingleProxyConfig["servicePort"] = servicePort
		// proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		tx.Commit()
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
	var (
		sumInfo    InterfaceEntity.SumInfo
		ser        DeleteServiceBody
		consulInfo InterfaceEntity.ConsulInfo
	)
	c.ShouldBind(&ser)
	// 数据物理删除
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
	if err := DB.First(&consulInfo).Error; err != nil {
	}
	var consulInfoObj = structs.Map(consulInfo)
	if consulInfoObj["ConsulAddress"].(string) != "" {
		Pkg.ConsulDeRegister(ser.ServerId)
	}
	if err := DB.First(&sumInfo).Update("server_sum", gorm.Expr("server_sum - ?", 1)).Error; err != nil {
	}
	if err := DB.Where("server_id = ?", ser.ServerId).Delete(&InterfaceEntity.ServiceInfo{}).Error; err != nil {
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
	DB.Close()
	// tx.Close()
}
