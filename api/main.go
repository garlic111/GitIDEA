package main

import (
	"log"
	"api/pkg/setting"
)

func main() {
	log.Println("启动中")
	setting.SetUp()
	log.Print(setting.ServerSetting.HttpPort)
}
