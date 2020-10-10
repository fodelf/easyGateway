package main

import (
	"fmt"
	service "gateway/database"

	// proxy "gateway/middleware/proxy"
	hystrix "gateway/hystrix"
	setting "gateway/pkg/setting"
	"gateway/routers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// dispatcher := proxy.NewDispatcher(2)
	// dispatcher.Run()
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	setting.Setup()
	service.ConnectDB()
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router := routers.InitRouter()
	router.Run(port)
	// go func() {
	// 	err := http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)
	// 	log.Fatal("sssssssssssssssssssssss")
	// 	log.Fatal(err)
	// }()
}
