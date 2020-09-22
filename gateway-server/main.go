package main

import (
	"fmt"
	service "gateway/database"
	setting "gateway/pkg/setting"
	"gateway/routers"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"
)

// func startWatch() {
// 	watchConfig := make(map[string]interface{})

// 	watchConfig["type"] = "service"
// 	watchConfig["service"] = "test"
// 	watchConfig["handler_type"] = "script"
// 	watchPlan, _ := watch.Parse(watchConfig)
// 	// util.CheckError(err)
// 	watchPlan.Handler = func(lastIndex uint64, result interface{}) {
// 		services := result.([]*api.ServiceEntry)
// 		str, _ := json.Marshal(services)
// 		// util.CheckError(err)
// 		fmt.Println(string(str))
// 		// fmt.Println(result)
// 	}
// 	if err := watchPlan.Run("127.0.0.1:8500"); err != nil {
// 		log.Fatalf("start watch error, error message: %s", err.Error())
// 	}
// }
func waitToUnRegistService() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	// if consulClient == nil {
	// 	return
	// }

	// if err := consulClient.Agent().ServiceDeregister(*serverID); err != nil {
	// 	log.Fatal(err)
	// }
	os.Exit(0)
}
func main() {
	// go waitToUnRegistService()
	// startWatch()
	setting.Setup()
	service.ConnectDB()
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router := routers.InitRouter()
	router.Run(port)
}
