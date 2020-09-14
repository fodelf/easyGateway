package service

import (
	"fmt"
	proxy "gateway/middleware/proxy"
	model "gateway/models"
	InterfaceEntity "gateway/models/InterfaceEntity"
	Utils "gateway/utils"
	"log"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "database_file.sqlite?cache=shared&mode=rwc")
	
	if err = DB.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
	// 初始化汇总数据
	var (
		sumInfo      InterfaceEntity.SumInfo
		chartInfo    InterfaceEntity.ChartInfo
		serviceInfos []InterfaceEntity.ServiceInfo
		consulInfo   InterfaceEntity.ConsulInfo
		rabbitMQInfo InterfaceEntity.RabbitMQInfo
	)
	if err := DB.Find(&sumInfo).Error; err != nil {
		sumInfo = InterfaceEntity.SumInfo{
			ServerSum:  0,
			WarningSum: 0,
			RequestSum: 0,
			FailSum:    0,
		}
		DB.Create(&sumInfo)
	}
	if err := DB.Find(&consulInfo).Error; err != nil {
		consulInfo = InterfaceEntity.ConsulInfo{
			ConsulId:      Utils.GenerateUUID(),
			ConsulAddress: "",
			Type:          "consul",
		}
		DB.Create(&consulInfo)
	}
	if err := DB.Find(&rabbitMQInfo).Error; err != nil {
		rabbitMQInfo := InterfaceEntity.RabbitMQInfo{
			RabbitMQId:       Utils.GenerateUUID(),
			RabbitMQAddress:  "",
			RabbitMQUserName: "",
			RabbitMQPassword: "",
			Type:             "rabbitMq",
		}
		DB.Create(&rabbitMQInfo)
	}
	if err := DB.Find(&sumInfo).Error; err != nil {
		sumInfo = InterfaceEntity.SumInfo{
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
			proxy.ProxyConfig = append(proxy.ProxyConfig, SingleProxyConfig)
		}
	}

	// DB.DB().SetMaxIdleConns(1000)
	// DB.DB().SetMaxOpenConns(5000)
	defer DB.Close()
	// DB.SingularTable(true)
	// DB.DB().SetMaxIdleConns(10)
	// DB.DB().SetMaxOpenConns(100)
	// DB.DB().SetConnMaxLifetime(30 * time.Minute)
	DB.LogMode(true)
}
