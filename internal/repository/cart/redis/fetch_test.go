package redis

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
)

func Test_cartRedisRepo_GetCart(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectGet(cartkey + "test").SetVal(`{"items":[{"ItemID":"1","Quantity":1}]}`)

	dbNF, mock := redismock.NewClientMock()
	mock.ExpectGet(cartkey + "test").SetVal(``)

	dbFail, mockFail := redismock.NewClientMock()
	mockFail.ExpectGet(cartkey + "test").SetErr(errors.New("test"))

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
		want    cartModel.Cart
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
			want: cartModel.Cart{
				Items: []cartModel.CartItem{
					{
						ItemID:   "1",
						Quantity: 1,
					},
				},
			},
		},
		{
			name: "negative test case not found",
			fields: fields{
				redis: dbNF,
			},
			args: args{
				userID: "test",
			},
			want: cartModel.Cart{},
		},
		{
			name: "negative test case error",
			fields: fields{
				redis: dbFail,
			},
			args: args{
				userID: "test",
			},
			want:    cartModel.Cart{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &cartRedisRepo{
				redis: tt.fields.redis,
			}
			got, err := sr.GetCart(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartRedisRepo.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartRedisRepo.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
