package user

import (
	"testing"

	"github.com/stevenkie/project-test/config"
	sessionRepo "github.com/stevenkie/project-test/internal/repository/session"
	userRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

func Test_userUC_ValidateSession(t *testing.T) {
	mockSessionRepo := &sessionRepo.RepositoryMock{
		GetTokenFunc: func(token string) (string, error) {
			if token == "not_found" {
				return "0", nil
			}
			return "1", nil
		},
	}

	type fields struct {
		config           *config.Config
		userPGRepo       userRepo.Repository
		sessionRedisRepo sessionRepo.Repository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValid bool
	}{
		{
			name: "positive test case",
			fields: fields{
				sessionRedisRepo: mockSessionRepo,
			},
			args: args{
				token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyMS0wNy0wM1QwNDo1MzoxMi43NDk5ODYxMDNaIiwicmFuZG9tIjo1NTc3MDA2NzkxOTQ3Nzc5NDEwfQ.Pb5ZRVjnbRuffwRjsfNvA7bMzsb6RzdC7lx7bSaWpVI",
			},
			wantValid: true,
		},
		{
			name: "positive not found",
			fields: fields{
				sessionRedisRepo: mockSessionRepo,
			},
			args: args{
				token: "Bearer not_found",
			},
			wantValid: false,
		},
		{
			name: "negative test case no token",
			fields: fields{
				sessionRedisRepo: mockSessionRepo,
			},
			args:      args{},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUC{
				config:           tt.fields.config,
				userPGRepo:       tt.fields.userPGRepo,
				sessionRedisRepo: tt.fields.sessionRedisRepo,
			}
			if gotValid := u.ValidateSession(tt.args.token); gotValid != tt.wantValid {
				t.Errorf("userUC.ValidateSession() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
