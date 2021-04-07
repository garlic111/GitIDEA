package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "api/pkg/e"

type Response struct {
	code int
	msg  string
	data interface{}
}

func ResponseWithJson(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
