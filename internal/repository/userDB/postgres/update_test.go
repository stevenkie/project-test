package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	"github.com/stevenkie/project-test/util"
)

func Test_userPostgresRepo_InsertUser(t *testing.T) {
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
		tx   util.Transaction
		data userModel.InsertUser
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
			upg := &userPostgresRepo{
				db: tt.fields.db,
			}
			if err := upg.InsertUser(tt.args.tx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("userPostgresRepo.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userPostgresRepo_UpdateUser(t *testing.T) {
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
		tx   util.Transaction
		data userModel.UpdateUser
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
			upg := &userPostgresRepo{
				db: tt.fields.db,
			}
			if err := upg.UpdateUser(tt.args.tx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("userPostgresRepo.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userPostgresRepo_DeleteUser(t *testing.T) {
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
			upg := &userPostgresRepo{
				db: tt.fields.db,
			}
			if err := upg.DeleteUser(tt.args.tx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("userPostgresRepo.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
