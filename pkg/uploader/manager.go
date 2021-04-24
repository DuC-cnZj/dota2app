package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/models"
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

func (m *Manager) Upload(file *multipart.FileHeader, name string, userID int) (contracts.File, error) {
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer open.Close()
	object, err := m.MinioClient().PutObject(context.Background(), utils.Config().MinioBucket, name, open, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return nil, err
	}

	marshal, _ := json.Marshal(&object)
	f := &models.File{
		Driver: models.DriverMinio,
		Path:   fmt.Sprintf("%s/%s/%s", m.MinioClient().EndpointURL().String(), object.Bucket, object.Key),
		UserID: userID,
		Info:   string(marshal),
		Size:   file.Size,
	}
	utils.DB().Create(f)

	return f, nil
}
