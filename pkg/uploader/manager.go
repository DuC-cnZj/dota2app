package uploader

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/DuC-cnZj/dota2app/pkg/utils"
	"github.com/minio/minio-go/v7"
)

type Manager struct {
	minioClient *minio.Client
}

func NewManager(minioClient *minio.Client) *Manager {
	return &Manager{minioClient: minioClient}
}

func (m *Manager) MinioClient() *minio.Client {
	return m.minioClient
}

func (m *Manager) SetMinioClient(c *minio.Client) {
	m.minioClient = c
}

func (m *Manager) Upload(file *multipart.FileHeader, name string) (url string, err error) {
	open, err := file.Open()
	if err != nil {
		return "", err
	}
	defer open.Close()
	object, err := m.MinioClient().PutObject(context.Background(), utils.Config().MinioBucket, name, open, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", m.minioClient.EndpointURL().String(), object.Bucket, object.Key), nil
}
