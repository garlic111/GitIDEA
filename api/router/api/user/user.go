package user

import (
	"api/models"
	"api/pkg/e"
	"api/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	name := c.Query("name")
	data := make(map[string] interface{})
	data["avatar"] = models.GetUserID(name)
	util.ResponseWithJson(e.SUCCESS,data,c)
}

func GetUserMobile(c *gin.Context) {
	mobile := c.PostForm("mobile")
	util.ResponseWithJson(e.SUCCESS,gin.H{"mobile":mobile},c)
}