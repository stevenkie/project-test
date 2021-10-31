package cart

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	httpModel "github.com/stevenkie/project-test/internal/model/http"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	log.Printf("error = %+v", err)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(httpModel.BuildAPIResponseError(statusCode, errors.Cause(err)))
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(httpModel.BuildAPIResponseSuccess(statusCode, data))
}

func (hd *HttpDelivery) GetCartByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(ErrorUserIDMustBeProvided))
		return
	}
	result, err := hd.cartUC.GetCart(userID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}

func (hd *HttpDelivery) AddCart(w http.ResponseWriter, r *http.Request) {
	var p cartModel.AddItemToCart
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = hd.cartUC.AddItemToCart(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, httpModel.GeneralSuccessMessage)
}
