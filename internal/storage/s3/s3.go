package s3

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageS3 struct {
	s3Client *s3.Client
}

func (s *StorageS3) Save(log *slog.Logger, file multipart.File, key string) error {
	// Загружаем объект в S3
	const op = "storage.s3.Save"

	_, err := s.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("imgnation"),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("application/octet-stream"), // Укажите правильный тип контента
	})

	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return nil

}

func (s *StorageS3) Get(key string) (io.ReadCloser, error) {
	const op = "storage.s3.Get"

	result, err := s.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("imgnation"),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return result.Body, nil
}

func New() (*StorageS3, error) {
	const op = "storage.s3.New"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	client := s3.NewFromConfig(cfg)

	return &StorageS3{s3Client: client}, nil
}
