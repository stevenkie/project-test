package cart

import (
	"errors"
	"reflect"
	"testing"

	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	cartRepo "github.com/stevenkie/project-test/internal/repository/cart"
	itemRepo "github.com/stevenkie/project-test/internal/repository/item"
)

func Test_cartUC_GetCart(t *testing.T) {
	mockCartRepo := &cartRepo.RepositoryMock{
		GetCartFunc: func(userID string) (cartModel.Cart, error) {
			if userID == "err" {
				return cartModel.Cart{}, errors.New("test")
			}
			return cartModel.Cart{Items: []cartModel.CartItem{{ItemID: "1", Quantity: 1}}}, nil
		},
	}

	type fields struct {
		itemPGRepo    itemRepo.Repository
		cartRedisRepo cartRepo.Repository
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cartModel.Cart
		wantErr bool
	}{
		{
			name: "posiive test case",
			fields: fields{
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				userID: "test",
			},
			want: cartModel.Cart{Items: []cartModel.CartItem{{ItemID: "1", Quantity: 1}}},
		},
		{
			name: "negative test case",
			fields: fields{
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				userID: "err",
			},
			want:    cartModel.Cart{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &cartUC{
				itemPGRepo:    tt.fields.itemPGRepo,
				cartRedisRepo: tt.fields.cartRedisRepo,
			}
			got, err := cu.GetCart(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartUC.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartUC.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartUC_AddItemToCart(t *testing.T) {
	mockCartRepo := &cartRepo.RepositoryMock{
		GetCartFunc: func(userID string) (cartModel.Cart, error) {
			if userID == "err" {
				return cartModel.Cart{}, errors.New("test")
			}
			return cartModel.Cart{Items: []cartModel.CartItem{{ItemID: "1", Quantity: 1}}}, nil
		},
		SaveCartFunc: func(userID string, cart cartModel.Cart) error {
			if userID == "err1" {
				return errors.New("test")
			}
			return nil
		},
	}

	type fields struct {
		itemPGRepo    itemRepo.Repository
		cartRedisRepo cartRepo.Repository
	}
	type args struct {
		input cartModel.AddItemToCart
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive test case",
			fields: fields{
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				input: cartModel.AddItemToCart{
					UserID:   "1",
					ItemID:   "1",
					Quantity: 1,
				},
			},
		},
		{
			name: "negative test case",
			fields: fields{
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				input: cartModel.AddItemToCart{
					UserID:   "err",
					ItemID:   "1",
					Quantity: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "negative test case error",
			fields: fields{
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				input: cartModel.AddItemToCart{
					UserID:   "err1",
					ItemID:   "1",
					Quantity: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &cartUC{
				itemPGRepo:    tt.fields.itemPGRepo,
				cartRedisRepo: tt.fields.cartRedisRepo,
			}
			if err := cu.AddItemToCart(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("cartUC.AddItemToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
