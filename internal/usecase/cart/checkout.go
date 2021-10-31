package cart

import (
	"fmt"

	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	itemModel "github.com/stevenkie/project-test/internal/model/item"
	"github.com/stevenkie/project-test/util"
)

//Checkout checkout current cart content
func (cu *cartUC) Checkout(userID string) (string, error) {
	// get all necessary data
	cart, err := cu.cartRedisRepo.GetCart(userID)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(cart.Items) == 0 {
		return "", errors.WithStack(errors.New(EmptyCartMessage))
	}
	allItems, err := cu.itemPGRepo.GetItemByIDs(cart.GetCartItemIDs())
	if err != nil {
		return "", errors.WithStack(err)
	}

	// validation
	err = validateItemAvailability(cart, allItems)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// reduce stock and empty cart
	err = util.WithTransaction(cu.itemPGRepo.GetDB(), func(tx util.Transaction) error {
		for itemID, qty := range cart.GetQtyMap() {
			err = cu.itemPGRepo.ReduceItemStock(tx, itemID, qty)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = cu.cartRedisRepo.EmptyCart(userID)
	if err != nil {
		return "", errors.WithStack(err)
	}

	total := calculateTotal(cart, allItems)

	return fmt.Sprintf(CheckoutMessage, total), nil
}

func validateItemAvailability(cart cartModel.Cart, items []itemModel.Item) error {
	qtyMap := cart.GetQtyMap()
	for itemID, cartQty := range qtyMap {
		for _, item := range items {
			if itemID == item.ID && cartQty > item.Quantity {
				return errors.WithStack(fmt.Errorf(OutOfStockMessage, item.Name, item.Quantity))
			}
		}
	}
	return nil
}

func calculateTotal(cart cartModel.Cart, items []itemModel.Item) int32 {
	qtyMap := cart.GetQtyMap()
	var total int32
	for itemID, cartQty := range qtyMap {
		for _, item := range items {
			if itemID == item.ID {
				subTotal := cartQty * item.Price
				total += subTotal
			}
		}
	}
	return total
}
