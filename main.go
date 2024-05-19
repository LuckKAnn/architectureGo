package main

import (
	"fmt"
	"ginDemo/models"
	"ginDemo/pkg/logging"
	"ginDemo/pkg/setting"
	"ginDemo/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	fmt.Println(endPoint)
	server := endless.NewServer(endPoint, routers.InitGinServer())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	//r.Run() // listen and serve on 0.0.0.0:8080
}
