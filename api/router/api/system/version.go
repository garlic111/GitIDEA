package system

import (
	"api/pkg/e"
	"api/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func GetAppVersion(c *gin.Context) {
	util.ResponseWithJson(e.SUCCESS, gin.H{"version": "v1.0.0"}, c)
}

func GetTime(c *gin.Context) {
	util.ResponseWithJson(e.SUCCESS, gin.H{
		"time": time.Now().Unix(),
	}, c)
}
