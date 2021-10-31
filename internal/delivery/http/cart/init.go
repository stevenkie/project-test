package cart

import (
	cartUsecase "github.com/stevenkie/project-test/internal/usecase/cart"
)

type HttpDelivery struct {
	cartUC cartUsecase.Usecase
}

func InitCartHttpDelivery(usecase cartUsecase.Usecase) *HttpDelivery {
	return &HttpDelivery{
		cartUC: usecase,
	}
}
