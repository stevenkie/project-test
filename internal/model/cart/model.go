package cart

import "github.com/go-playground/validator"

var validate *validator.Validate

type Identifier struct {
	ID string `json:"id" validate:"required"`
}

func (p Identifier) Validate() error {
	return validate.Struct(p)
}

type CartItem struct {
	ItemID   string
	Quantity int32
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type AddItemToCart struct {
	UserID   string `json:"id"`
	ItemID   string `json:"item_id"`
	Quantity int32  `json:"quantity"`
}

func (p Cart) GetCartItemIDs() []string {
	result := make([]string, 0)
	for _, v := range p.Items {
		result = append(result, v.ItemID)
	}
	return result
}

func (p Cart) GetCartItemQuantity(itemID string) int32 {
	for _, v := range p.Items {
		if v.ItemID == itemID {
			return v.Quantity
		}
	}
	return 0
}

func (p Cart) AddToCart(input AddItemToCart) Cart {
	found := false
	for i, v := range p.Items {
		if v.ItemID == input.ItemID {
			p.Items[i].Quantity += input.Quantity
			found = true
		}
	}
	if !found {
		p.Items = append(p.Items, CartItem{
			ItemID:   input.ItemID,
			Quantity: input.Quantity,
		})
	}
	for i, v := range p.Items {
		if v.Quantity <= 0 {
			p.Items = removeCartItem(p.Items, i)
		}
	}
	return p
}

func removeCartItem(slice []CartItem, s int) []CartItem {
	return append(slice[:s], slice[s+1:]...)
}

func (p Cart) GetQtyMap() map[string]int32 {
	result := make(map[string]int32)
	for _, cI := range p.Items {
		result[cI.ItemID] += cI.Quantity
	}

	return result
}
