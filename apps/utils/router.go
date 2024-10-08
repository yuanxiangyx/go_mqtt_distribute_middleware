package utils

import (
	"github.com/gin-gonic/gin"
	"mqtt_pro/apps/middleware"
	router2 "mqtt_pro/apps/router"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors)
	router.Use(middleware.GinLogger())
	router.Use(middleware.ExceptionMiddleware)

	ApiGroup := router.Group("/api")

	router2.InitTestRouter(ApiGroup)

	return router
}

func Register(r *gin.RouterGroup, httpMethods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	var routes gin.IRoutes
	for _, httpMethod := range httpMethods {
		routes = r.Handle(httpMethod, relativePath, handlers...)
	}
	return routes
}
