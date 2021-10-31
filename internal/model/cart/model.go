package cart

type CartItem struct {
	ItemID   string
	Quantity int32
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type AddItemToCart struct {
	UserID   string
	ItemID   string
	Quantity int32
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
