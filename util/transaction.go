package util

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// TxFn holds the function that will be executed
type TxFn func(Transaction) error

// WithTransaction creates a new transaction and handles rollback/commit based on result (error)
func WithTransaction(db *sqlx.DB, fn TxFn) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil { // recover from panic
			tx.Rollback()
			panic(p)
		} else if err != nil { // rollback because error happens
			tx.Rollback()
		} else { // no error, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
