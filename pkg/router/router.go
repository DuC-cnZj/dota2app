package router

import (
	"net/http"

	"github.com/DuC-cnZj/dota2app/pkg/controllers"
	t "github.com/DuC-cnZj/dota2app/pkg/translator"
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	authC := controllers.NewAuthController()
	authMiddleware, _ := authC.AuthMiddleware()

	e.NoRoute(func(ctx *gin.Context) {
		ctx.Data(http.StatusNotFound, "application/json", []byte(`{"code": 404, "message": "404 not found"}`))
	})

	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	api := e.Group("/api", t.I18nMiddleware())
	{
		api.POST("/login", authMiddleware.LoginHandler)
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}
