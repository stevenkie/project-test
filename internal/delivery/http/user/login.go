package user

import (
	"encoding/json"
	"net/http"

	loginModel "github.com/stevenkie/project-test/internal/model/login"
)

func (hd *HttpDelivery) Login(w http.ResponseWriter, r *http.Request) {
	var p loginModel.Login
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	token, err := hd.userUC.Login(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, token)
}
