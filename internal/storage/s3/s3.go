package s3

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageS3 struct {
	db *s3.Client
}

func (s *StorageS3) Save(log *slog.Logger, file multipart.File) error {
	// Загружаем объект в S3
	const op = "storage.s3.Save"
	_, err := s.db.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("imgination"),
		Key:         aws.String("test"), //Имя под которым будет в S3
		Body:        file,
		ContentType: aws.String("application/octet-stream"), // Укажите правильный тип контента
	})

	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}
	
	return nil

}

func New(log *slog.Logger) (*StorageS3, error) {
	const op = "storage.s3.New"
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)

	return &StorageS3{db: client}, nil
}
