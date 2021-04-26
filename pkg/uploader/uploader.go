package uploader

import (
	"context"
	"fmt"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const ReadonlyPolicy = `
{
    "Version":"2012-10-17",
    "Statement":[
        {
            "Effect":"Allow",
            "Principal":{
                "AWS":[
                    "*"
                ]
            },
            "Action":[
                "s3:GetBucketLocation"
            ],
            "Resource":[
                "arn:aws:s3:::%s"
            ]
        },
        {
            "Effect":"Allow",
            "Principal":{
                "AWS":[
                    "*"
                ]
            },
            "Action":[
                "s3:ListBucket"
            ],
            "Resource":[
                "arn:aws:s3:::%s"
            ],
            "Condition":{
                "StringEquals":{
                    "s3:prefix":[
                        "*.*"
                    ]
                }
            }
        },
        {
            "Effect":"Allow",
            "Principal":{
                "AWS":[
                    "*"
                ]
            },
            "Action":[
                "s3:GetObject"
            ],
            "Resource":[
                "arn:aws:s3:::%s/*.**"
            ]
        }
    ]
}
`

func Init(app contracts.ApplicationInterface) (contracts.StorageInterface, error) {
	minioClient, err := minio.New(app.Config().MinioEndpoint, &minio.Options{Creds: credentials.NewStaticV4(app.Config().MinioAccessKey, app.Config().MinioAccessSecret, "")})
	bucketName := app.Config().MinioBucket
	if err != nil {
		return nil, err
	}
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}

		if err := minioClient.SetBucketPolicy(context.Background(), bucketName, fmt.Sprintf(ReadonlyPolicy, bucketName, bucketName, bucketName)); err != nil {
			return nil, err
		}
	}

	return NewManager(minioClient), nil
}
