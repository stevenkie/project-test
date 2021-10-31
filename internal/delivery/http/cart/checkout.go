package cart

import (
	"encoding/json"
	"net/http"

	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	httpModel "github.com/stevenkie/project-test/internal/model/http"

	"github.com/pkg/errors"
)

func (hd *HttpDelivery) CheckoutCarts(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get(httpModel.HeaderAuth)
	validSession := hd.userUC.ValidateSession(authorizationHeader)
	if !validSession {
		writeErrorResponse(w, http.StatusForbidden, errors.New(ErrorForbidden))
		return
	}
	var p cartModel.Identifier
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if p.ID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(ErrorUserIDMustBeProvided))
		return
	}
	result, err := hd.cartUC.Checkout(p.ID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}
