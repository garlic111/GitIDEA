package util

import "github.com/gin-gonic/gin"
import "api/pkg/e"

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func ResponseWithJson(code int, data interface{}, c *gin.Context) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
