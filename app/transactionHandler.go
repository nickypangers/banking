package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/dto"
	service "github.com/nickypangers/banking/services"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th *TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
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

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	request.AccountId = accountId
	transaction, appError := th.service.Deposit(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, transaction)

}
