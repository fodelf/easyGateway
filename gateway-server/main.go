package main

import (
	"fmt"
	setting "gateway/pkg/setting"
	"gateway/routers"

	_ "github.com/mattn/go-sqlite3"
)

//  "time"

func main() {
	setting.Setup()
	// ConnectDB()
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router := routers.InitRouter()
	router.Run(port)
}
