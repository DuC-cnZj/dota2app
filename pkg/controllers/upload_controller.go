package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"

	"github.com/DuC-cnZj/dota2app/pkg/auth"
	"github.com/DuC-cnZj/dota2app/pkg/derrors"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/response"
	t "github.com/DuC-cnZj/dota2app/pkg/translator"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (*UploadController) Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	if file == nil {
		response.Error(ctx, 422, t.TransError(derrors.FileRequired, t.GetLocale(ctx)))
		return
	}

	if !validateContentType(file.Header.Get("Content-Type")) {
		response.Error(ctx, 422, t.TransError(derrors.FileMustBeImage, t.GetLocale(ctx)))
		return
	}

	f, err := utils.Storage().Upload(file, generateFileName(ctx, file.Filename), contracts.TypeAvatar, auth.User(ctx).ID)
	if err != nil {
		dlog.Error(err)
		response.Error(ctx, 500, err)
		return
	}

	response.Success(ctx, 201, gin.H{
		"path":          f.GetFullPath(),
		"size":          f.GetSize(),
		"humanize_size": f.ToHumanizeSize(),
		"relative_path": f.GetRelativePath(),
	})
}

func generateFileName(ctx *gin.Context, oldName string) string {
	user := auth.User(ctx)

	return fmt.Sprintf("%s-%s-%s%s", time.Now().Format("20060102"), user.Name, utils.RandomString(10), filepath.Ext(oldName))
}

func validateContentType(contentType string) bool {
	var validated bool

	for _, ct := range []string{"image/gif", "image/jpeg", "image/png"} {
		if contentType == ct {
			validated = true
			break
		}
	}

	return validated
}
