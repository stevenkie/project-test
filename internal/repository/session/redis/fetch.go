package redis

import "github.com/pkg/errors"

//GetToken from redis and get user_id
func (sr *sessionRedisRepo) GetToken(token string) (string, error) {
	userID, err := sr.redis.Get(sr.redis.Context(), token).Result()
	if err != nil {
		return userID, errors.WithStack(err)
	}
	return userID, nil
}
