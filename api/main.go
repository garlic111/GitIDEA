package main

import (
	"api/models"
	"api/pkg/logging"
	"api/pkg/setting"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello, api 正在启动中...")

	router := router.InitRouter()

	fmt.Println("hello,mybranch")

	// 初始化设置
	setting.SetUp()
	logging.SetUp()
	models.SetUp()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:      router,
		ReadTimeout:  setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout, MaxHeaderBytes: 1 << 20}

	s.ListenAndServe()
}
