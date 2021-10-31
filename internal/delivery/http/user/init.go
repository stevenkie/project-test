package user

import (
	userUsecase "github.com/stevenkie/project-test/internal/usecase/user"
)

type HttpDelivery struct {
	userUC userUsecase.Usecase
}

func InitUserHttpDelivery(usecase userUsecase.Usecase) *HttpDelivery {
	return &HttpDelivery{
		userUC: usecase,
	}
}
