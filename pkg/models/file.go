package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	"github.com/dustin/go-humanize"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

//RelativePath
type File struct {
	ID int `json:"id" gorm:"primaryKey;"`

	// 上传的 oss 驱动
	Driver uint8 `json:"driver" gorm:"type:tinyint;comment:1=minio;"`

	// 全路径
	RelativePath string `json:"relative_path" gorm:"type:VARCHAR(255);"`

	// 文件大小
	Size int64 `json:"size"`

	// 类型: avatar,  file...
	Type string `json:"type" gorm:"type:VARCHAR(20);"`

	// 上传用户
	UserID int `json:"user_id"`

	// oss 返回的 obj 的整个 json
	Info string `json:"result" gorm:"type:text;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (f *File) GetSize() uint64 {
	return uint64(f.Size)
}

func (f *File) GetUploadType() contracts.UploadType {
	return f.Type
}

func (f *File) GetUserID() int {
	return f.UserID
}

func (f *File) GetRelativePath() string {
	return f.RelativePath
}

func (f *File) GetDriver() uint8 {
	return f.Driver
}

func (f *File) GetDriverName() string {
	return contracts.DriverNameMap[f.Driver]
}

func (f *File) GetFullPath() string {
	switch f.Driver {
	case contracts.DriverMinio:
		endpointUrl := utils.Storage().(contracts.WithMinio).MinioClient().EndpointURL().String()

		return fmt.Sprintf(
			"%s/%s",
			strings.TrimRight(endpointUrl, "/"),
			strings.TrimLeft(f.RelativePath, "/"))
	default:
		return f.RelativePath
	}
}

func (f *File) ToMinioObject() (*minio.ObjectInfo, error) {
	if f.Driver != contracts.DriverMinio {
		return nil, errors.New("file driver is not minio")
	}
	var info minio.ObjectInfo
	if err := json.Unmarshal([]byte(f.Info), &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func (f *File) ToHumanizeSize() string {
	return humanize.Bytes(uint64(f.Size))
}
