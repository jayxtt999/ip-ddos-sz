package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/routers"
	"log"
	"net/http"
)

func init()  {
	setting.Setup()
	logging.Setup()
}

func main()  {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}




