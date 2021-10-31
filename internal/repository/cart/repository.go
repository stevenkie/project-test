package cart

import cartModel "github.com/stevenkie/project-test/internal/model/cart"

// Repository repository interface for interacting with Cart Redis
type Repository interface {
	//GetCart get cart content per user_id
	GetCart(userID string) (cartModel.Cart, error)
	//SaveCart save cart content to redis
	EmptyCart(userID string) error
	//EmptyCart save empty cart content to cache
	SaveCart(userID string, cart cartModel.Cart) error
}
