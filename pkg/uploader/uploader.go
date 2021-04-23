package uploader

import (
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Init(app contracts.ApplicationInterface) (contracts.StorageInterface, error) {
	minioClient, err := minio.New(app.Config().MinioEndpoint, &minio.Options{Creds: credentials.NewStaticV4(app.Config().MinioAccessKey, app.Config().MinioAccessSecret, "")})

	if err != nil {
		return nil, err
	}

	return NewManager(minioClient), nil
}
