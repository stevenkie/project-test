package item

import (
	"github.com/jmoiron/sqlx"
	itemModel "github.com/stevenkie/project-test/internal/model/item"
	"github.com/stevenkie/project-test/util"
)

// Repository repository interface for interacting with item DB
type Repository interface {
	//GetItemByIDs from db and return array of items
	GetItemByIDs(itemIDs []string) ([]itemModel.Item, error)
	//ReduceItemStock substract qty from table 'items'
	ReduceItemStock(tx util.Transaction, itemID string, qty int32) error
	//GetDB for used withing util.WithTransaction
	GetDB() *sqlx.DB
}
