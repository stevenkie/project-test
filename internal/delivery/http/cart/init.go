package cart

import (
	cartUsecase "github.com/stevenkie/project-test/internal/usecase/cart"
	userUsecase "github.com/stevenkie/project-test/internal/usecase/user"
)

type HttpDelivery struct {
	cartUC cartUsecase.Usecase
	userUC userUsecase.Usecase
}

func InitCartHttpDelivery(usecase cartUsecase.Usecase, userUsecase userUsecase.Usecase) *HttpDelivery {
	return &HttpDelivery{
		cartUC: usecase,
		userUC: userUsecase,
	}
}
