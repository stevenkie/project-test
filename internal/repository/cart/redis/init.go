package redis

import (
	"github.com/go-redis/redis/v8"
	cartRDRepo "github.com/stevenkie/project-test/internal/repository/cart"
)

type cartRedisRepo struct {
	redis *redis.Client
}

func InitSessionRedisRepo(redis *redis.Client) cartRDRepo.Repository {
	return &cartRedisRepo{
		redis: redis,
	}
}
