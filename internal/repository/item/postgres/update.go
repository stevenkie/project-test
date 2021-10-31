package postgres

import (
	"github.com/pkg/errors"
	"github.com/stevenkie/project-test/util"
)

//ReduceItemStock substract qty from table 'items'
func (ipg *itemPostgresRepo) ReduceItemStock(tx util.Transaction, itemID string, qty int32) error {
	_, err := tx.Exec(reduceItemStockByID, itemID, qty)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
