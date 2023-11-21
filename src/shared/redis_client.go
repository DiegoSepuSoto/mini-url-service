package shared

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) RedisStringResult
}

type RedisStringResult interface {
	Result() (string, error)
}

type redisClient struct {
	redisClient *redis.Client
}

func (c *redisClient) Get(ctx context.Context, key string) RedisStringResult {
	return c.redisClient.Get(ctx, key)
}

func CreateRedisClient() *redisClient {
	return &redisClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_HOST"),
		}),
	}
}
