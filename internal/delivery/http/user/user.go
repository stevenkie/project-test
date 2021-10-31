package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	httpModel "github.com/stevenkie/project-test/internal/model/http"
	userModel "github.com/stevenkie/project-test/internal/model/user"
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

func (hd *HttpDelivery) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(ErrorUserIDMustBeProvided))
		return
	}
	result, err := hd.userUC.GetUserByID(userID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}

func (hd *HttpDelivery) InsertUser(w http.ResponseWriter, r *http.Request) {
	var p userModel.InsertUser
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = hd.userUC.InsertUser(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, httpModel.GeneralSuccessMessage)
}

func (hd *HttpDelivery) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var p userModel.UpdateUser
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(ErrorUserIDMustBeProvided))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	p.ID = userID
	err = hd.userUC.UpdateUser(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, httpModel.GeneralSuccessMessage)
}

func (hd *HttpDelivery) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(ErrorUserIDMustBeProvided))
		return
	}
	err := hd.userUC.DeleteUser(userID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, httpModel.GeneralSuccessMessage)
}
