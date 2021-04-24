package router

import (
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/DuC-cnZj/dota2app/pkg/controllers"
	t "github.com/DuC-cnZj/dota2app/pkg/translator"
	"github.com/gin-gonic/gin"
)

const (
	JSONContentType = "application/json"
)

func Init(e *gin.Engine) {
	var cd = cors.DefaultConfig()
	cd.AllowAllOrigins = true
	cd.AddAllowHeaders("X-Requested-With", "Authorization", "Accept-Language")
	e.Use(cors.New(cd))

	authC := controllers.NewAuthController()
	authMiddleware, _ := authC.AuthMiddleware()

	e.NoRoute(func(ctx *gin.Context) {
		ctx.Data(http.StatusNotFound, JSONContentType, []byte(`{"code": 404, "message": "404 not found"}`))
	})

	e.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, JSONContentType, []byte(`{"success": "true"}`))
	})

	api := e.Group("/api", t.I18nMiddleware())
	{
		api.POST("/login", authMiddleware.LoginHandler)
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	up := controllers.NewUploadController()
	auth := api.Group("/", authMiddleware.MiddlewareFunc())
	{
		auth.POST("/userinfo", authC.Info)
		auth.POST("/update_userinfo", authC.UpdateInfo)
		auth.POST("/upload", up.Upload)
	}
}
