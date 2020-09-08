package main

import (
	"fmt"
	service "gateway/database"
	setting "gateway/pkg/setting"
	"gateway/routers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	setting.Setup()
	service.ConnectDB()
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router := routers.InitRouter()
	router.Run(port)
}
