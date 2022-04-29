package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/dto"
	service "github.com/nickypangers/banking/services"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	if customerId == "" {
		writeResponse(w, http.StatusBadRequest, "Customer id is required")
		return
	}

	if accountId == "" {
		writeResponse(w, http.StatusBadRequest, "Account id is required")
		return
	}

	account, err := ah.service.GetAccount(customerId, accountId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, account)
}

func (ah *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	account, appError := ah.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, account)
}
