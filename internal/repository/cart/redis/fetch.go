package redis

import (
	"encoding/json"

	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
)

//GetCart get cart content per user_id
func (sr *cartRedisRepo) GetCart(userID string) (cartModel.Cart, error) {
	var result cartModel.Cart
	cartItem, err := sr.redis.Get(sr.redis.Context(), cartkey+userID).Result()
	if err != nil {
		return result, errors.WithStack(err)
	}
	if len(cartItem) == 0 {
		return result, nil
	}
	json.Unmarshal([]byte(cartItem), &result)
	return result, nil
}
