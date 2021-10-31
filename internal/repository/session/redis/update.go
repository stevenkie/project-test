package redis

import (
	"github.com/pkg/errors"
)

//SetUserToken to redis, this cache is used for token auth
func (sr *sessionRedisRepo) SetUserToken(userID string, token string) error {
	err := sr.redis.Set(sr.redis.Context(), token, userID, DefaultCacheExpiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
