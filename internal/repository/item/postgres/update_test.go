package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stevenkie/project-test/util"
)

func Test_itemPostgresRepo_ReduceItemStock(t *testing.T) {
	mocktx := &util.TransactionMock{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			return nil, nil
		},
	}

	mocktxNegative := &util.TransactionMock{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			return nil, errors.New("error here")
		},
	}

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		tx     util.Transaction
		itemID string
		qty    int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive test case",
			args: args{
				tx: mocktx,
			},
		},
		{
			name: "negative test case",
			args: args{
				tx: mocktxNegative,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipg := &itemPostgresRepo{
				db: tt.fields.db,
			}
			if err := ipg.ReduceItemStock(tt.args.tx, tt.args.itemID, tt.args.qty); (err != nil) != tt.wantErr {
				t.Errorf("itemPostgresRepo.ReduceItemStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
