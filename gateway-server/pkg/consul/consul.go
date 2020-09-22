package pkg

import (
	"fmt"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"log"
	"net/http"
	"strconv"

	// "github.com/CatchZeng/dingtalk"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/jinzhu/gorm"
)

// var Client consulapi

// // 实例化客户端
// func initClient() {
// 	config := consulapi.DefaultConfig()
// 	config.Address = "127.0.0.1:8500"
// 	client, err := consulapi.NewClient(config)

// 	if err != nil {
// 		log.Fatal("consul client error : ", err)
// 	} else {
// 		Client = client
// 	}
// }
type ImportServiceBody struct {
	ServerId            string                   `json:"serverId"`
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

// consul 服务注册
func RegisterServer(importServiceBody InterfaceEntity.ImportServiceBody) {
	// client1 := dingtalk.NewClient(accessToken, secret)
	// msg := dingtalk.NewTextMessage().SetContent("测试文本").SetAt([]string{"+86-18651892475"}, false)
	// client1.Send(msg)
	var (
		consulInfo InterfaceEntity.ConsulInfo
	)
	DB, err := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	if err := DB.First(&consulInfo).Error; err != nil {
	}
	DB.Close()
	var consulInfoObj = structs.Map(consulInfo)
	config := consulapi.DefaultConfig()
	port := strconv.Itoa(consulInfoObj["ConsulPort"].(int))
	config.Address = consulInfoObj["ConsulAddress"].(string) + ":" + port
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = importServiceBody.UseConsulId              // 服务节点的名称
	registration.Name = importServiceBody.ServiceName            // 服务名称
	registration.Port = 3000                                     // 服务端口
	registration.Tags = []string{importServiceBody.UseConsulTag} // tag，可以为空
	registration.Address = importServiceBody.ServiceAddress      // 服务 IP

	// 健康检查 支持http及grpc 回调接口
	checkPort := 3000
	registration.Check = &consulapi.AgentServiceCheck{ // 健康检查
		HTTP:     fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, importServiceBody.UseConsulCheckPath),
		Timeout:  "3s",  // 超时时间
		Interval: "30s", // 健康检查间隔
		// DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
		// GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	// 服务注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
	// KV get值
	// data, _, _ := client.KV().Get("config", nil)
	// fmt.Println(string(data.Value))
	// get string
	// str, err := client.Agent().Get("config").String()
	// if err != nil {
	// 	log.Fatal("consul client error : ", err)
	// }else{
	// 	log.Fatal("register server error : ", str)
	// }
}

// consul 健康检测
func heathCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
