package redis

import (
	"encoding/json"

	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
)

//SaveCart save cart content to redis
func (sr *cartRedisRepo) SaveCart(userID string, cart cartModel.Cart) error {
	converted, err := json.Marshal(cart)
	if err != nil {
		return errors.WithStack(err)
	}
	err = sr.redis.Set(sr.redis.Context(), cartkey+userID, string(converted), 0).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//EmptyCart save empty cart content to cache
func (sr *cartRedisRepo) EmptyCart(userID string) error {
	err := sr.redis.Del(sr.redis.Context(), cartkey+userID).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
