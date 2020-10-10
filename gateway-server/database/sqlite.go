package service

import (
	"fmt"
	proxy "gateway/middleware/proxy"
	model "gateway/models"
	InterfaceEntity "gateway/models/InterfaceEntity"
	pkg "gateway/pkg/consul"
	Utils "gateway/utils"
	"log"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")

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
	} else {
		if err := DB.First(&consulInfo).Error; err != nil {

		} else {
			if consulInfo.ConsulAddress != "" {
				go pkg.InitWatch()
			}
		}
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
	if err := DB.Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), "all").First(&chartInfo).Error; err != nil {
		chartInfo := InterfaceEntity.ChartInfo{
			Time:     time.Now().Format("2006/01/02"),
			Total:    0,
			Success:  0,
			Fail:     0,
			ServerId: "all",
		}
		DB.Create(&chartInfo)
	}
	if err := DB.Find(&serviceInfos).Error; err != nil {
	} else {
		for i := 0; i < len(serviceInfos); i++ {
			var serviceInfo = serviceInfos[i]
			// 很奇怪
			var serviceBreak = serviceInfo.ServiceBreak
			var serviceLimit = serviceInfo.ServiceLimit
			var generateUUID = serviceInfo.ServerId
			if serviceBreak > 0 && serviceLimit > 0 {
				fmt.Printf("SSS")
				hystrix.ConfigureCommand(generateUUID, hystrix.CommandConfig{
					//多长时间 超时
					Timeout: 500000,
					//最大并发数
					MaxConcurrentRequests: 10,
					//错误百分比,错误率达到这个数值开启熔断
					ErrorPercentThreshold: 90,
					// //当熔断器被打开后，SleepWindow的时间就是控制过多久后去尝试服务是否可用了(毫秒)
					SleepWindow: 10,
					// //最小请求数，只有到达这个数量后才判断是否开启熔断
					RequestVolumeThreshold: 1,
				})
			} else if serviceBreak > 0 && serviceLimit == 0 {
				hystrix.ConfigureCommand(generateUUID, hystrix.CommandConfig{
					ErrorPercentThreshold:  serviceBreak,
					RequestVolumeThreshold: 1000000000000,
				})
			} else if serviceBreak == 0 && serviceLimit > 0 {
				hystrix.ConfigureCommand(generateUUID, hystrix.CommandConfig{
					MaxConcurrentRequests:  serviceLimit,
					RequestVolumeThreshold: 1000000000000,
				})
			}
			if err := DB.Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), serviceInfo.ServerId).First(&chartInfo).Error; err != nil {
				chartInfo := InterfaceEntity.ChartInfo{
					Time:     time.Now().Format("2006/01/02"),
					Total:    0,
					Success:  0,
					Fail:     0,
					ServerId: serviceInfo.ServerId,
				}
				DB.Create(&chartInfo)
			}
			var SingleProxyConfig map[string]interface{} = map[string]interface{}{
				"serverId":       serviceInfo.ServerId,
				"serviceAddress": serviceInfo.ServiceAddress,
				"servicePort":    serviceInfo.ServicePort,
				"serviceRules":   []map[string]interface{}{},
			}
			var serviceRules = serviceInfo.ServiceRules
			var rules = strings.Split(serviceRules, ",")
			for i := 0; i < len(rules); i++ {
				var ruleArray = strings.Split(rules[i], ";")
				var rule = map[string]interface{}{
					"url": ruleArray[0],
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
				} else if len(ruleArray) == 2 {
					rule = map[string]interface{}{
						"url": ruleArray[0],
						"pathReWrite": map[string]interface{}{
							"oldPath": ruleArray[1],
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
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// DB.DB().SetConnMaxLifetime(30 * time.Minute)
	DB.LogMode(true)
}
