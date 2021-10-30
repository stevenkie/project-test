package redis

import (
	"github.com/go-redis/redis/v8"
	sessionRDRepo "github.com/stevenkie/project-test/internal/repository/session"
)

type sessionRedisRepo struct {
	redis *redis.Client
}

func InitSessionRedisRepo(redis *redis.Client) sessionRDRepo.Repository {
	return &sessionRedisRepo{
		redis: redis,
	}
}
