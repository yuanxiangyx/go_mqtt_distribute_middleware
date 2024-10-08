package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}
