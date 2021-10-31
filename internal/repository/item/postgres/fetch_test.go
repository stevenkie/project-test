package postgres

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	itemModel "github.com/stevenkie/project-test/internal/model/item"
)

func Test_itemPostgresRepo_GetDB(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   *sqlx.DB
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipg := &itemPostgresRepo{
				db: tt.fields.db,
			}
			if got := ipg.GetDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemPostgresRepo.GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemPostgresRepo_GetItemByIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		itemIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []itemModel.Item
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "quantity", "price"}).
					AddRow(1, "abc", 10, 50)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: []itemModel.Item{
				{
					ID:       "1",
					Name:     "abc",
					Quantity: 10,
					Price:    50,
				},
			},
		},
		{
			name: "error",
			mock: func() {
				mock.ExpectQuery("(.*?)").WillReturnError(errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipg := &itemPostgresRepo{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := ipg.GetItemByIDs(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("itemPostgresRepo.GetItemByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemPostgresRepo.GetItemByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
