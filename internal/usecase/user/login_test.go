package user

import (
	"errors"
	"testing"

	"github.com/stevenkie/project-test/config"
	loginModel "github.com/stevenkie/project-test/internal/model/login"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	sessionRepo "github.com/stevenkie/project-test/internal/repository/session"
	userRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

func Test_userUC_Login(t *testing.T) {
	mockUserPGRepo := &userRepo.RepositoryMock{
		GetUserByEmailFunc: func(email string) (userModel.User, error) {
			if email == "test2" {
				return userModel.User{
					Password: "testpassword2",
				}, nil
			}
			if email == "error" {
				return userModel.User{}, errors.New("error")
			}
			if email == "tester2" {
				return userModel.User{
					ID:       "2",
					Email:    "tester2",
					Password: "$2a$10$LjIQStxh0z/7o1JcGJrJ6eAQXvmEDYPGqS9TBdXxwhQuK/VTrXs7a",
				}, nil
			}
			return userModel.User{
				Email:    "tester",
				Password: "$2a$10$LjIQStxh0z/7o1JcGJrJ6eAQXvmEDYPGqS9TBdXxwhQuK/VTrXs7a",
			}, nil
		},
	}

	mockSessionRepo := &sessionRepo.RepositoryMock{
		SetUserTokenFunc: func(userID, token string) error {
			if userID == "2" {
				return errors.New("test")
			}
			return nil
		},
	}

	mCfg := &config.Config{
		Server: config.ServerConfig{
			Secret: "secret_sample",
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
	}
	type args struct {
		data loginModel.Login
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
				data: loginModel.Login{
					Email:    "tester",
					Password: "testpassword",
				},
			},
			fields: fields{
				config:           mCfg,
				userPGRepo:       mockUserPGRepo,
				sessionRedisRepo: mockSessionRepo,
			},
			wantErr: false,
		},
		{
			name: "negative test case db error",
			args: args{
				data: loginModel.Login{
					Email:    "error",
					Password: "testpassword",
				},
			},
			fields: fields{
				config:           mCfg,
				userPGRepo:       mockUserPGRepo,
				sessionRedisRepo: mockSessionRepo,
			},
			wantErr: true,
		},
		{
			name: "negative test case wrong pass",
			args: args{
				data: loginModel.Login{
					Email:    "test2",
					Password: "testpassword",
				},
			},
			fields: fields{
				config:           mCfg,
				userPGRepo:       mockUserPGRepo,
				sessionRedisRepo: mockSessionRepo,
			},
			wantErr: true,
		},
		{
			name: "negative redis missing",
			args: args{
				data: loginModel.Login{
					Email:    "tester2",
					Password: "testpassword",
				},
			},
			fields: fields{
				config:           mCfg,
				userPGRepo:       mockUserPGRepo,
				sessionRedisRepo: mockSessionRepo,
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
			_, err := u.Login(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUC.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
