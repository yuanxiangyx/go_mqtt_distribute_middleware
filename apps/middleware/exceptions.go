package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Api struct {
	Code    int
	Message string
}

func ExceptionMiddleware(c *gin.Context) {
	//异常处理
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case *Api:
				zap.S().Error(t.Message)
				c.JSON(t.Code, gin.H{
					"message": t.Message,
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "服务器内部异常",
				})
			}
			c.Abort()
		}
	}()
	c.Next()
}
