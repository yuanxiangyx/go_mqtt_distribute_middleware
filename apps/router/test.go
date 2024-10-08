package router

import (
	"github.com/gin-gonic/gin"
	"mqtt_pro/apps/api"
)

func InitTestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		Register(TestRouter, []string{"GET"}, "/", api.TestApi)
	}
}
