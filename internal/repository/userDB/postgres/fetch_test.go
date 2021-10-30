package postgres

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	userModel "github.com/stevenkie/project-test/internal/model/user"
)

func Test_userPostgresRepo_GetDB(t *testing.T) {
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
			upg := &userPostgresRepo{
				db: tt.fields.db,
			}
			if got := upg.GetDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userPostgresRepo.GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userPostgresRepo_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    userModel.User
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "address", "is_active"}).
					AddRow(1, "abc", "aaa", "bbb", true)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: userModel.User{
				ID:       "1",
				Email:    "abc",
				Password: "aaa",
				Address:  "bbb",
				IsActive: true,
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
			upg := &userPostgresRepo{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := upg.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userPostgresRepo.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userPostgresRepo.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userPostgresRepo_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    userModel.User
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "address", "is_active"}).
					AddRow(1, "abc", "aaa", "bbb", true)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: userModel.User{
				ID:       "1",
				Email:    "abc",
				Password: "aaa",
				Address:  "bbb",
				IsActive: true,
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
			upg := &userPostgresRepo{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := upg.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userPostgresRepo.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userPostgresRepo.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
