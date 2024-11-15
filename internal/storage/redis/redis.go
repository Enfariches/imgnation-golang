package redis

import (
	"context"
	"fmt"
	"img/internal/storage"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type StorageRedis struct {
	rCLient *redis.Client
}

func (rdb *StorageRedis) CachedSave(file []byte, key string) error {
	const op = "storage.redis.CachedSave"

	err := rdb.rCLient.Set(ctx, key, file, time.Minute).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (rdb *StorageRedis) CachedGet(key string) ([]byte, error) {
	const op = "storage.redis.CachedGet"

	result, err := rdb.rCLient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, storage.ErrRedisNotFoundOctet
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func New() (*StorageRedis, error) {
	const op = "storage.redis.New"
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &StorageRedis{rCLient: client}, nil
}
