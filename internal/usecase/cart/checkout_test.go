package cart

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	itemModel "github.com/stevenkie/project-test/internal/model/item"
	cartRepo "github.com/stevenkie/project-test/internal/repository/cart"
	itemRepo "github.com/stevenkie/project-test/internal/repository/item"
	"github.com/stevenkie/project-test/util"
)

func Test_cartUC_Checkout(t *testing.T) {
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockCartRepo := &cartRepo.RepositoryMock{
		GetCartFunc: func(userID string) (cartModel.Cart, error) {
			if userID == "err" {
				return cartModel.Cart{}, errors.New("test")
			}
			return cartModel.Cart{Items: []cartModel.CartItem{{ItemID: "1", Quantity: 1}, {ItemID: "2", Quantity: 1}}}, nil
		},
		EmptyCartFunc: func(userID string) error {
			return nil
		},
	}
	mockItemPGRepo := &itemRepo.RepositoryMock{
		GetDBFunc: func() *sqlx.DB {
			result := sqlx.NewDb(db, "sqlmock")
			mockSQL.ExpectBegin()
			return result
		},
		GetItemByIDsFunc: func(itemIDs []string) ([]itemModel.Item, error) {
			return []itemModel.Item{{ID: "1", Name: "test", Quantity: 1, Price: 10}, {ID: "2", Name: "test2", Quantity: 1, Price: 10}}, nil
		},
		ReduceItemStockFunc: func(tx util.Transaction, itemID string, qty int32) error {
			mockSQL.ExpectCommit()
			return nil
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
		want    string
		wantErr bool
	}{
		{
			name: "positive test case",
			fields: fields{
				itemPGRepo:    mockItemPGRepo,
				cartRedisRepo: mockCartRepo,
			},
			args: args{
				userID: "test",
			},
			want:    "Your checkout total is Rp.20",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &cartUC{
				itemPGRepo:    tt.fields.itemPGRepo,
				cartRedisRepo: tt.fields.cartRedisRepo,
			}
			got, err := cu.Checkout(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartUC.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cartUC.Checkout() = %v, want %v", got, tt.want)
			}
		})
	}
}
