package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/domain"
	service "github.com/nickypangers/banking/services"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":9000", router)

}
