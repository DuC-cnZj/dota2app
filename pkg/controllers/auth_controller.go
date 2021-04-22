package controllers

import (
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/auth"
	"github.com/DuC-cnZj/dota2app/pkg/derrors"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/models"
	"github.com/DuC-cnZj/dota2app/pkg/response"
	t "github.com/DuC-cnZj/dota2app/pkg/translator"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (au *AuthController) authenticate(ctx *gin.Context, input *LoginForm) (*models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = utils.DB().Where("`email` = ?", input.Email).First(&user).Error; err != nil {
		return nil, t.TransError(derrors.UserNotFound, t.GetLocale(ctx))
	}

	if err := utils.PasswordCheck(user.Password, input.Password); err != nil {
		return nil, t.TransError(derrors.PasswordError, t.GetLocale(ctx))
	}

	return &user, nil
}

func (au *AuthController) Info(ctx *gin.Context) {
	user := auth.User(ctx)

	response.Success(ctx, 200, gin.H{
		"id":     user.ID,
		"email":  user.Email,
		"name":   user.Name,
		"avatar": user.Avatar,
	})
}

func (au *AuthController) AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "sso",
		Key:        []byte(utils.Config().AppSecret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}

			return jwt.MapClaims{}
		},
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var (
				input LoginForm
				err   error
			)

			if err = ctx.ShouldBind(&input); err != nil {
				dlog.Error(err)
				return "", err
			}

			return au.authenticate(ctx, &input)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(ctx *gin.Context, code int, message string) {
			response.Error(ctx, code, message)
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			return t.Trans(e.Error(), t.GetLocale(c))
		},
	})
}
