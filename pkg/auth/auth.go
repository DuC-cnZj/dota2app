package auth

import (
	"github.com/DuC-cnZj/dota2app/pkg/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func User(ctx *gin.Context) *models.User {
	if u, exists := ctx.Get(jwt.IdentityKey); exists {
		return u.(*models.User)
	}

	return nil
}
