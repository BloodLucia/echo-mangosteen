package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type cache struct {
	*redis.Client
}

// GetString implements Cache.
func (rdb *cache) GetString(ctx context.Context, key string) (val string, has bool, err error) {
	val, err = rdb.Client.Get(ctx, key).Result()
	if err != nil {
		return "", false, err
	}

	return val, true, nil
}

type Cache interface {
	GetString(ctx context.Context, key string) (val string, has bool, err error)
}

func NewCahce() (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &cache{Client: client}, nil
}
