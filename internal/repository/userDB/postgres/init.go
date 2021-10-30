package postgres

import (
	"github.com/jmoiron/sqlx"
)

//userPostgresRepo  struct for postgres repository
type userPostgresRepo struct {
	db *sqlx.DB
}
