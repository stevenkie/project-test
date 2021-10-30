package postgres

import "github.com/jmoiron/sqlx"

// GetDB function for getting database for util WithTransactionFunction
func (upg *userPostgresRepo) GetDB() *sqlx.DB {
	return upg.db
}
