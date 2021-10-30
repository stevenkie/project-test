package user

import (
	"testing"

	"github.com/stevenkie/project-test/config"
	sessionRepo "github.com/stevenkie/project-test/internal/repository/session"
	userRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

func TestInitUserUsecase(t *testing.T) {
	type args struct {
		cfg   *config.Config
		db    userRepo.Repository
		redis sessionRepo.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		InitUserUsecase(tt.args.cfg, tt.args.db, tt.args.redis)
	}
}
