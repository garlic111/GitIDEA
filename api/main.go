package main

import (
	"api/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("启动中")
	setting.SetUp()
	log.Print(setting.ServerSetting.RunMode)
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "success",
		})
	})
	r.Run()
}
