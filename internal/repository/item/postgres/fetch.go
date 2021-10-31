package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	itemModel "github.com/stevenkie/project-test/internal/model/item"
	"github.com/stevenkie/project-test/util"
)

// GetDB function for getting database for util WithTransactionFunction
func (ipg *itemPostgresRepo) GetDB() *sqlx.DB {
	return ipg.db
}

//GetItemByIDs from db and return array of items
func (ipg *itemPostgresRepo) GetItemByIDs(itemIDs []string) ([]itemModel.Item, error) {
	var result []itemModel.Item
	err := ipg.db.Select(&result, getItemsByIDs, pq.Array(util.Distinct(itemIDs)))
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}
