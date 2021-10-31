package cart

import cartModel "github.com/stevenkie/project-test/internal/model/cart"

// Usecase contract for cart interface
type Usecase interface {
	GetCart(userID string) (cartModel.Cart, error)
	AddItemToCart(input cartModel.AddItemToCart) error
	Checkout(userID string) (string, error)
}
