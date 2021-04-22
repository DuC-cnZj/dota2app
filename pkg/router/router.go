package router

import "github.com/gin-gonic/gin"

func Init(e *gin.Engine) {
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})
}