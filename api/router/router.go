package router

import (
	"api/middleware/cors"
	"api/middleware/myjwt"
	"api/pkg/e"
	"api/pkg/setting"
	system "api/router/api/system"
	user "api/router/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	var authMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AllUserAuthorizator{})
	r.POST("/login", authMiddleware.LoginHandler)

	//auth
	r.NoRoute(func(c *gin.Context) {
		code := e.PAGE_NOT_FOUND
		c.JSON(404, gin.H{"code": code, "message": e.GetMsg(code), "data": ""})
	})

	// 系统信息 路由分组
	apiSys := r.Group("/api/system")
	{
		apiSys.GET("version", system.GetAppVersion)
		apiSys.GET("time", system.GetTime)
	}

	apiUser := r.Group("api/user")
	{
		apiUser.GET("info", user.GetUserInfo)
		apiUser.GET("userMobile",user.GetUserMobile)
	}

	return r
}
