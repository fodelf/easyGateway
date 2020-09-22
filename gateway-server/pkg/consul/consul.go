package pkg

import (
	"encoding/json"
	"fmt"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"log"
	"net/http"
	"strconv"

	// "github.com/CatchZeng/dingtalk"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"

	// "github.com/hashicorp/consul/watch"
	"github.com/jinzhu/gorm"
)

func ConsulKVTest() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "172.16.242.129:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// KV, put值
	values := "test"
	key := "go-consul-test/172.16.242.129:8100"
	client.KV().Put(&consulapi.KVPair{Key: key, Flags: 0, Value: []byte(values)}, nil)

	// KV get值
	data, _, _ := client.KV().Get(key, nil)
	fmt.Println(string(data.Value))

	// KV list
	datas, _, _ := client.KV().List("go", nil)
	for _, value := range datas {
		fmt.Println(value)
	}
	keys, _, _ := client.KV().Keys("go", "", nil)
	fmt.Println(keys)
}

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

// 取消consul注册的服务
func ConsulDeRegister(useConsulId string) {
	var (
		consulInfo InterfaceEntity.ConsulInfo
	)
	// 创建连接consul服务配置
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
	client.Agent().ServiceDeregister(useConsulId)
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
	watchConfig := make(map[string]interface{})

	watchConfig["type"] = "service"
	watchConfig["service"] = "test"
	watchConfig["handler_type"] = "script"
	watchPlan, _ := watch.Parse(watchConfig)
	// util.CheckError(err)
	watchPlan.Handler = func(lastIndex uint64, result interface{}) {
		services := result.([]*api.ServiceEntry)
		str, _ := json.Marshal(services)
		// util.CheckError(err)
		fmt.Println(string(str))
		// fmt.Println(result)
	}
	if err := watchPlan.Run("127.0.0.1:8500"); err != nil {
		log.Fatalf("start watch error, error message: %s", err.Error())
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
