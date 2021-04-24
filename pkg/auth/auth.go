package auth

import (
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/models"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func User(ctx *gin.Context) *models.User {
	var user models.User
	if u, exists := ctx.Get(jwt.IdentityKey); exists {
		if err := utils.DB().Where("`id` = ?", u.(int)).First(&user).Error; err != nil {
			dlog.Error(err)
			return nil
		}

		return &user
	}

	return nil
}

func ID(ctx *gin.Context) int {
	if id, exists := ctx.Get(jwt.IdentityKey); exists {
		return id.(int)
	}

	return 0
}

func UserPreload(ctx *gin.Context, preloads ...string) *models.User {
	var user models.User
	dlog.Debug("UserPreload: ", preloads)
	if u, exists := ctx.Get(jwt.IdentityKey); exists {
		db := utils.DB()
		for _, s := range preloads {
			db = db.Preload(s)
		}
		if err := db.Where("`id` = ?", u.(int)).First(&user).Error; err != nil {
			dlog.Error(err)
			return nil
		}

		return &user
	}

	return nil
}
