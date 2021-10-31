package postgres

import (
	"github.com/jmoiron/sqlx"
	itemPGRepo "github.com/stevenkie/project-test/internal/repository/item"
)

//userPostgresRepo  struct for postgres repository
type itemPostgresRepo struct {
	db *sqlx.DB
}

// InitItemPGRepo initialize new repository for userPG
func InitItemPGRepo(db *sqlx.DB) itemPGRepo.Repository {
	return &itemPostgresRepo{
		db: db,
	}
}
