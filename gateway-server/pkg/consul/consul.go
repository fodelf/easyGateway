package proxy

import (
	"log"
	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
	"net/http"
)
// consul 服务注册
func RegisterServer()  {
	fmt.Println("	fmt.Println(err)")
	// 创建consul客户端
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
			log.Fatal("consul client error : ", err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "c1"      // 服务节点的名称
	registration.Name = "wsp"      // 服务名称
	registration.Port = 3000              // 服务端口
	registration.Tags = []string{"wsp"} // tag，可以为空
	registration.Address = "127.0.0.1"      // 服务 IP

	// 健康检查 支持http及grpc 回调接口
	checkPort := 3000
	registration.Check = &consulapi.AgentServiceCheck{ // 健康检查
			HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/ping"),
			Timeout:                        "3s", // 超时时间
			Interval:                       "5s",  // 健康检查间隔
			// DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
			// GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	// 服务注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
			log.Fatal("register server error : ", err)
	}
	// KV get值
	data, _, _ := client.KV().Get("config", nil)
	fmt.Println(string(data.Value))
	// get string
	// str, err := client.Agent().Get("config").String()
	// if err != nil {
	// 	log.Fatal("consul client error : ", err)
	// }else{
	// 	log.Fatal("register server error : ", str)
	// }
}

// consul 健康检测
func heathCheck(c *gin.Context)  {
	c.JSON(http.StatusOK, "ok")
}