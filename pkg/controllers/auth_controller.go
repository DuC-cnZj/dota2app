package controllers

import (
	"net/http"
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

type UserInfo struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	AvatarId          int    `json:"avatar_id"`
	Intro             string `json:"intro"`
	BackgroundImage   string `json:"background_image"`
	BackgroundImageId int    `json:"background_image_id"`
}

func (au *AuthController) Info(ctx *gin.Context) {
	var user *models.User = auth.UserPreload(ctx, "Avatar", "BackgroundImage")

	response.Success(ctx, 200, &UserInfo{
		Id:                user.ID,
		Email:             user.Email,
		Name:              user.Name,
		Avatar:            user.Avatar.GetFullPath(),
		AvatarId:          user.Avatar.ID,
		Intro:             user.Intro,
		BackgroundImage:   user.BackgroundImage.GetFullPath(),
		BackgroundImageId: user.BackgroundImage.ID,
	})
}

type UpdateInput struct {
	Name              string `json:"name" binding:"required"`
	Intro             string `json:"intro"`
	AvatarID          int    `json:"avatar_id" binding:"required"`
	BackgroundImageID int    `json:"background_image_id"`
}

func (au *AuthController) UpdateInfo(ctx *gin.Context) {
	var (
		input UpdateInput
		user  *models.User = auth.UserPreload(ctx, "Avatar", "BackgroundImage")
	)

	if err := ctx.ShouldBind(&input); err != nil {
		response.Error(ctx, 422, err)
		return
	}

	if err := utils.DB().Model(user).Updates(map[string]interface{}{
		"name":  input.Name,
		"intro": input.Intro,
	}).Error; err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	var bg models.File
	if input.BackgroundImageID != user.BackgroundImage.ID {
		if input.BackgroundImageID == 0 {
			utils.DB().Model(user).Association("BackgroundImage").Clear()
		} else if utils.DB().Where("`id` = ?", input.BackgroundImageID).First(&bg).Error == nil {
			utils.DB().Model(user).Association("BackgroundImage").Append(&bg)
		}
	}
	var avatar models.File
	if input.AvatarID != user.Avatar.ID && utils.DB().Where("`id` = ?", input.AvatarID).First(&avatar).Error == nil {
		utils.DB().Model(user).Association("Avatar").Append(&avatar)
	}

	response.Success(ctx, http.StatusOK, &UserInfo{
		Id:                user.ID,
		Email:             user.Email,
		Name:              user.Name,
		Avatar:            user.Avatar.GetFullPath(),
		AvatarId:          user.Avatar.ID,
		Intro:             user.Intro,
		BackgroundImage:   user.BackgroundImage.GetFullPath(),
		BackgroundImageId: user.BackgroundImage.ID,
	})
}

func (au *AuthController) GetHistoryAvatars(ctx *gin.Context) {
	var paginate Pagination
	user := &models.User{ID: auth.ID(ctx)}
	if err := ctx.ShouldBind(&paginate); err != nil {
		response.Error(ctx, 422, "")
		return
	}

	res, total := user.HistoryAvatarsWithPaginate(&paginate.Page, &paginate.PageSize)

	response.Pagination(ctx, http.StatusOK, res, paginate.Page, paginate.PageSize, total)
}

func (au *AuthController) GetHistoryBackgroundImages(ctx *gin.Context) {
	var paginate Pagination
	user := &models.User{ID: auth.ID(ctx)}
	if err := ctx.ShouldBind(&paginate); err != nil {
		response.Error(ctx, 422, "")
		return
	}

	res, total := user.HistoryAvatarsWithPaginate(&paginate.Page, &paginate.PageSize)

	response.Pagination(ctx, http.StatusOK, res, paginate.Page, paginate.PageSize, total)
}

func (au *AuthController) AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(utils.Config().AppSecret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)

			return int(claims["id"].(float64))
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{"id": v.ID}
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
			// TODO 权限认证
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
