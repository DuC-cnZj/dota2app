package contracts

import (
	"mime/multipart"
)

type StorageInterface interface {
	Upload(file *multipart.FileHeader, name string) (url string, err error)
}
