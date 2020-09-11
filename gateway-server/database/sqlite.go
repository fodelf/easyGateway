package service

import (
	"fmt"
	proxy "gateway/middleware/proxy"
	model "gateway/models"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"log"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "gateway.db")
	// defer DB.Close()
	if err = DB.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
	// 初始化汇总数据
	var (
		sumInfo      InterfaceEntity.SumInfo
		chartInfo    InterfaceEntity.ChartInfo
		serviceInfos []InterfaceEntity.ServiceInfo
	)
	if err := DB.Find(&sumInfo).Error; err != nil {
		sumInfo := InterfaceEntity.SumInfo{
			ServerSum:  0,
			WarningSum: 0,
			RequestSum: 0,
			FailSum:    0,
		}
		DB.Create(&sumInfo)
	}
	if err := DB.Find(&chartInfo).Error; err != nil {
		chartInfo := InterfaceEntity.ChartInfo{
			Time:    time.Now().Format("2006/01/02"),
			Total:   0,
			Success: 0,
			Fail:    0,
		}
		DB.Create(&chartInfo)
	}
	if err := DB.Find(&serviceInfos).Error; err != nil {
	} else {
		for i := 0; i < len(serviceInfos); i++ {
			var serviceInfo = structs.Map(serviceInfos[i])
			fmt.Println(serviceInfo)
			var SingleProxyConfig map[string]interface{} = map[string]interface{}{
				"serviceAddress": serviceInfo["ServiceAddress"],
				"servicePort":    serviceInfo["ServicePort"],
				"serviceRules":   []map[string]interface{}{},
			}
			var serviceRules = serviceInfo["ServiceRules"].(string)
			var rules = strings.Split(serviceRules, ",")
			fmt.Println(rules)
			for i := 0; i < len(rules); i++ {
				var ruleArray = strings.Split(rules[i], ";")
				// var oldPath = ""
				// var newPath = ""
				fmt.Println(ruleArray)
				// m, err := Utils.JsonToMap(ruleArray[1])
				// if err != nil {
				// 	fmt.Printf("Convert json to map failed with error: %+v\n", err)
				// } else {
				// 	var keys = Utils.GetKeys(m)
				// 	oldPath = keys[0]
				// 	newPath = m[oldPath]
				// }
				var rule = map[string]interface{}{
					"url": "",
					"pathReWrite": map[string]interface{}{
						"oldPath": "",
						"newPath": "",
					},
				}
				if len(ruleArray) == 3 {
					rule = map[string]interface{}{
						"url": ruleArray[0],
						"pathReWrite": map[string]interface{}{
							"oldPath": ruleArray[1],
							"newPath": ruleArray[2],
						},
					}
				} else {
					rule = map[string]interface{}{
						"url": ruleArray[0],
						"pathReWrite": map[string]interface{}{
							"oldPath": "",
							"newPath": "",
						},
					}
				}

				SingleProxyConfig["serviceRules"] = append(SingleProxyConfig["serviceRules"].([]map[string]interface{}), rule)
			}
			// var SingleProxyConfig map[string]interface{} = map[string]interface{}{
			// 	// "serviceAddress": "",
			// 	// "prot":           0,
			// 	// "serviceRules":   []map[string]interface{}{},
			// 	"serviceAddress": "172.23.0.187",
			// 	"prot":           9990,
			// 	"serviceRules": []map[string]interface{}{
			// 		{
			// 			"url": "/api/market/platform-api",
			// 			"pathReWrite": map[string]interface{}{
			// 				"oldPath": "/api/market/platform-api",
			// 				"newPath": "",
			// 			},
			// 		},
			// 	},
			// }
			proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		}
	}

	DB.DB().SetMaxIdleConns(1000)
	DB.DB().SetMaxOpenConns(5000)
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
	DB.LogMode(true)
}
