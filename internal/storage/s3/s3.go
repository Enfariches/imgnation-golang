package s3

import (
	"context"
	"img/internal/lib/logger/sl"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func New(log *slog.Logger) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("Error load cfg for S3", sl.Error(err))
		return err
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)

	// Запрашиваем список бакетов
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Error("Error get ListBuckets", sl.Error(err))
		return err
	}

	for _, bucket := range result.Buckets {
		log.Info("bucket=%s creation time=%s", aws.ToString(bucket.Name), bucket.CreationDate.Local().Format("2006-01-02 15:04:05 Monday"))
	}
	return nil
}
