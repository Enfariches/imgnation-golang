package s3

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	cfg "img/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageS3 struct {
	s3Client   *s3.Client
	bucketName string
}

func (s *StorageS3) Save(file multipart.File, key string) error {
	// Загружаем объект в S3
	const op = "storage.s3.Save"

	_, err := s.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("application/octet-stream"), // Укажите правильный тип контента
	})

	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return nil

}

func (s *StorageS3) Get(key string) ([]byte, error) {
	const op = "storage.s3.Get"

	object, err := s.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
	defer object.Body.Close()

	result, err := io.ReadAll(object.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func New(cfg_storage cfg.Storage) (*StorageS3, error) {
	const op = "storage.s3.New"

	os.Setenv("AWS_ACCESS_KEY_ID", cfg_storage.Aws_access_key_id)
	os.Setenv("AWS_SECRET_ACCESS_KEY", cfg_storage.Aws_secret_access_key)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg_storage.Region))

	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg_storage.Endpoint_url)
	})

	return &StorageS3{s3Client: client, bucketName: cfg_storage.BucketName}, nil
}
