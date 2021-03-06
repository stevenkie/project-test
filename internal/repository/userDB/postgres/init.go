package postgres

import (
	"github.com/jmoiron/sqlx"
	userPGRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

//userPostgresRepo  struct for postgres repository
type userPostgresRepo struct {
	db *sqlx.DB
}

// InitUserPGRepo initialize new repository for userPG
func InitUserPGRepo(db *sqlx.DB) userPGRepo.Repository {
	return &userPostgresRepo{
		db: db,
	}
}
