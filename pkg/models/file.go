package models

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/dustin/go-humanize"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

const (
	TypeAvatar          contracts.UploadType = "avatar"
	TypeBackgroundImage contracts.UploadType = "background_image"
)

const (
	// minio
	DriverMinio contracts.UploadDriver = iota + 1
)

var DriverNameMap map[contracts.UploadDriver]string = map[contracts.UploadDriver]string{
	DriverMinio: "minio",
}

type File struct {
	ID int `json:"id" gorm:"primaryKey;"`

	// 上传的 oss 驱动
	Driver uint8 `json:"driver" gorm:"type:tinyint;not null;comment:1=minio;"`

	// 全路径
	Path string `json:"path" gorm:"type:VARCHAR(255);not null;"`

	// 文件大小
	Size int64 `json:"size" gorm:"not null;default:0;"`

	// 上传用户
	UserID int `json:"user_id" gorm:"not null;"`

	FileableID   int
	FileableType string `gorm:"not null;"`

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
	return f.FileableType
}

func (f *File) GetUserID() int {
	return f.UserID
}

func (f *File) GetRelativePath() string {
	parse, err := url.Parse(f.Path)
	if err != nil {
		return f.Path
	}

	return parse.Path
}

func (f *File) GetDriver() uint8 {
	return f.Driver
}

func (f *File) GetDriverName() string {
	return DriverNameMap[f.Driver]
}

func (f *File) GetFullPath() string {
	return f.Path
}

func (f *File) GetID() int {
	return f.ID
}

func (f *File) ToMinioObject() (*minio.ObjectInfo, error) {
	if f.Driver != DriverMinio {
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
