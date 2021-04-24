package contracts

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type UploadType = string

type UploadDriver = uint8

type WithMinio interface {
	MinioClient() *minio.Client
	SetMinioClient(c *minio.Client)
}

type File interface {
	GetID() int
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
	Upload(file *multipart.FileHeader, name string, userID int) (f File, err error)
}
