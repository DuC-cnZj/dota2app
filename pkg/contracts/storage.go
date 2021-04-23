package contracts

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type UploadType = string

type UploadDriver = uint8

const (
	TypeAvatar UploadType = "avatar"
)

const (
	// minio
	DriverMinio UploadDriver = iota + 1
)

var DriverNameMap map[UploadDriver]string = map[UploadDriver]string{
	DriverMinio: "minio",
}

type WithMinio interface {
	MinioClient() *minio.Client
	SetMinioClient(c *minio.Client)
}

type File interface {
	GetUploadType() UploadType
	GetUserID() int
	GetFullPath() string
	GetRelativePath() string
	GetSize() uint64
	GetDriver() uint8
	GetDriverName() string

	ToHumanizeSize() string
	ToMinioObject() (*minio.ObjectInfo, error)
}

type StorageInterface interface {
	Upload(file *multipart.FileHeader, name string, uploadType UploadType, userID int) (f File, err error)
}
