package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stevenkie/project-test/config"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	sessionRepo "github.com/stevenkie/project-test/internal/repository/session"
	userRepo "github.com/stevenkie/project-test/internal/repository/userdb"
	"github.com/stevenkie/project-test/util"
)

func Test_userUC_GetUserByID(t *testing.T) {
	mockUserPGRepo := &userRepo.RepositoryMock{
		GetUserByIDFunc: func(userID string) (userModel.User, error) {
			if userID == "2" {
				return userModel.User{}, errors.New("test")
			}
			return userModel.User{ID: "1"}, nil
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    userModel.User
		wantErr bool
	}{
		{
			name: "positive test case",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				userID: "1",
			},
			want: userModel.User{
				ID: "1",
			},
			wantErr: false,
		},
		{
			name: "negative test case",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				userID: "2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUC{
				config:           tt.fields.config,
				userPGRepo:       tt.fields.userPGRepo,
				sessionRedisRepo: tt.fields.sessionRedisRepo,
			}
			got, err := u.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUC.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUC.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUC_InsertUser(t *testing.T) {
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockUserPGRepo := &userRepo.RepositoryMock{
		GetDBFunc: func() *sqlx.DB {
			result := sqlx.NewDb(db, "sqlmock")
			mockSQL.ExpectBegin()
			return result
		},
		InsertUserFunc: func(tx util.Transaction, data userModel.InsertUser) error {
			if data.Email == "testing@gmail.com" {
				mockSQL.ExpectRollback()
				return errors.New("test")
			}
			mockSQL.ExpectCommit()
			return nil
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
	}
	type args struct {
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
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.InsertUser{
					Email:    "aaa@gmail.com",
					Password: "abcdefghij",
					Address:  "aaa bbb"},
			},
			wantErr: false,
		},
		{
			name: "negative test case verify",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.InsertUser{},
			},
			wantErr: true,
		},
		{
			name: "negative test case error",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.InsertUser{
					Email:    "testing@gmail.com",
					Password: "abcdefghij",
					Address:  "aaa bbb",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUC{
				config:           tt.fields.config,
				userPGRepo:       tt.fields.userPGRepo,
				sessionRedisRepo: tt.fields.sessionRedisRepo,
			}
			if err := u.InsertUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("userUC.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUC_UpdateUser(t *testing.T) {
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockUserPGRepo := &userRepo.RepositoryMock{
		GetDBFunc: func() *sqlx.DB {
			result := sqlx.NewDb(db, "sqlmock")
			mockSQL.ExpectBegin()
			return result
		},
		UpdateUserFunc: func(tx util.Transaction, data userModel.UpdateUser) error {
			if data.Email == "testing@gmail.com" {
				mockSQL.ExpectRollback()
				return errors.New("test")
			}
			mockSQL.ExpectCommit()
			return nil
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
	}
	type args struct {
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
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.UpdateUser{
					ID:      "1",
					Email:   "aaa@gmail.com",
					Address: "aaa bbb"},
			},
			wantErr: false,
		},
		{
			name: "negative test case verify",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.UpdateUser{},
			},
			wantErr: true,
		},
		{
			name: "negative test case error",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				data: userModel.UpdateUser{
					ID:      "1",
					Email:   "testing@gmail.com",
					Address: "aaa bbb",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUC{
				config:           tt.fields.config,
				userPGRepo:       tt.fields.userPGRepo,
				sessionRedisRepo: tt.fields.sessionRedisRepo,
			}
			if err := u.UpdateUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("userUC.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUC_DeleteUser(t *testing.T) {
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockUserPGRepo := &userRepo.RepositoryMock{
		GetDBFunc: func() *sqlx.DB {
			result := sqlx.NewDb(db, "sqlmock")
			mockSQL.ExpectBegin()
			return result
		},
		DeleteUserFunc: func(tx util.Transaction, userID string) error {
			if userID == "2" {
				mockSQL.ExpectRollback()
				return errors.New("test")
			}
			mockSQL.ExpectCommit()
			return nil
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
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
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				userID: "1",
			},
			wantErr: false,
		},
		{
			name: "negative test case error",
			fields: fields{
				userPGRepo: mockUserPGRepo,
			},
			args: args{
				userID: "2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUC{
				config:           tt.fields.config,
				userPGRepo:       tt.fields.userPGRepo,
				sessionRedisRepo: tt.fields.sessionRedisRepo,
			}
			if err := u.DeleteUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("userUC.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
