package redis

import (
	"testing"

	"github.com/go-redis/redis/v8"
	cartRDRepo "github.com/stevenkie/project-test/internal/repository/cart"
)

func TestInitCartRedisRepo(t *testing.T) {
	type args struct {
		redis *redis.Client
	}
	tests := []struct {
		name string
		args args
		want cartRDRepo.Repository
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCartRedisRepo(tt.args.redis)
		})
	}
}
