package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/nickypangers/banking/services"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customers)

}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
