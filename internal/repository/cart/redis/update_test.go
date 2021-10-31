package redis

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
)

func Test_cartRedisRepo_SaveCart(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectSet(cartkey+"test", `{"items":[{"ItemID":"1","Quantity":1}]}`, 0).SetVal("succes")

	dbFail, mockFail := redismock.NewClientMock()
	mockFail.ExpectSet("test", "test", 0).SetErr(errors.New("test"))

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		userID string
		cart   cartModel.Cart
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
				redis: db,
			},
			args: args{
				userID: "test",
				cart: cartModel.Cart{
					Items: []cartModel.CartItem{
						{
							ItemID:   "1",
							Quantity: 1,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "positive test case",
			fields: fields{
				redis: dbFail,
			},
			args: args{
				userID: "test",
				cart:   cartModel.Cart{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &cartRedisRepo{
				redis: tt.fields.redis,
			}
			if err := sr.SaveCart(tt.args.userID, tt.args.cart); (err != nil) != tt.wantErr {
				t.Errorf("cartRedisRepo.SaveCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cartRedisRepo_EmptyCart(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectDel(cartkey + "test").SetVal(1)

	dbFail, mock := redismock.NewClientMock()
	mock.ExpectDel(cartkey + "test").SetErr(errors.New("test"))

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		userID string
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
				redis: db,
			},
			args: args{
				userID: "test",
			},
		},
		{
			name: "negative test case",
			fields: fields{
				redis: dbFail,
			},
			args: args{
				userID: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &cartRedisRepo{
				redis: tt.fields.redis,
			}
			if err := sr.EmptyCart(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("cartRedisRepo.EmptyCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
