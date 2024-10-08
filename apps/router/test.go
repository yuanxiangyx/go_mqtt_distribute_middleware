package router

import (
	"github.com/gin-gonic/gin"
	"mqtt_pro/apps/api"
	utils2 "mqtt_pro/apps/utils"
)

func InitTestRouter(Router *gin.RouterGroup) {
	//用户相关的路由
	UserRouter := Router.Group("test")
	{
		// 用户认证登陆
		utils2.Register(UserRouter, []string{"POST"}, "login", api.TestApi)
	}
}
