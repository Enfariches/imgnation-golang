package redis

import (
	"context"
	"fmt"
	"img/internal/config"
	"img/internal/storage"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type StorageRedis struct {
	rCLient *redis.Client
	TTL     time.Duration
}

func (rdb *StorageRedis) CacheSave(file []byte, key string) error {
	const op = "storage.redis.CachedSave"

	err := rdb.rCLient.Set(ctx, key, file, rdb.TTL).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (rdb *StorageRedis) CacheGet(key string) ([]byte, error) {
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

func New(cfg config.Cache) (*StorageRedis, error) {
	const op = "storage.redis.New"

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &StorageRedis{rCLient: client, TTL: cfg.TTL}, nil
}
