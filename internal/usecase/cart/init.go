package cart

import (
	cartRepo "github.com/stevenkie/project-test/internal/repository/cart"
	itemRepo "github.com/stevenkie/project-test/internal/repository/item"
)

type cartUC struct {
	itemPGRepo    itemRepo.Repository
	cartRedisRepo cartRepo.Repository
}

// InitCartUsecase for cart
func InitCartUsecase(db itemRepo.Repository, redis cartRepo.Repository) Usecase {
	return &cartUC{
		itemPGRepo:    db,
		cartRedisRepo: redis,
	}
}
