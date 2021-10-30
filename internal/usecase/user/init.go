package user

import (
	"github.com/stevenkie/project-test/config"
	sessionRepo "github.com/stevenkie/project-test/internal/repository/session"
	userRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

type userUC struct {
	config           *config.Config
	userPGRepo       userRepo.Repository
	sessionRedisRepo sessionRepo.Repository
}

// NewUserUsecase for
func InitUserUsecase(cfg *config.Config, db userRepo.Repository, redis sessionRepo.Repository) Usecase {
	return &userUC{
		config:           cfg,
		userPGRepo:       db,
		sessionRedisRepo: redis,
	}
}
