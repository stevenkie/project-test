package cart

import (
	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
)

//AddItemToCart add single item to cart
func (cu *cartUC) AddItemToCart(input cartModel.AddItemToCart) error {
	// get current data
	cart, err := cu.cartRedisRepo.GetCart(input.UserID)
	if err != nil {
		return errors.WithStack(err)
	}
	cart = cart.AddToCart(input)
	// save cart
	err = cu.cartRedisRepo.SaveCart(input.UserID, cart)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//GetCart get single data of user
func (cu *cartUC) GetCart(userID string) (cartModel.Cart, error) {
	// get current data
	cart, err := cu.cartRedisRepo.GetCart(userID)
	if err != nil {
		return cartModel.Cart{}, errors.WithStack(err)
	}
	return cart, nil
}

//EmptyCart add empty cart
func (cu *cartUC) EmptyCart(userID string) error {
	err := cu.cartRedisRepo.EmptyCart(userID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
