package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/mocks/service"
)

var cRouter *mux.Router
var ch CustomerHandler
var mockCustomerService *service.MockCustomerService

func mockCustomerHandlerSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandler{mockCustomerService}

	cRouter = mux.NewRouter()
	cRouter.HandleFunc("/customers", ch.getAllCustomers)
	cRouter.HandleFunc("/customers/{account_id:[0-9]+}", ch.getCustomer)
	return func() {
		cRouter = nil
		defer ctrl.Finish()
	}
}

func Test_get_all_customers_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := mockCustomerHandlerSetup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "John Doe", City: "New York", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1001", Name: "Rob Holding", City: "London", Zipcode: "12304", DateofBirth: "1994-03-19", Status: "1"},
	}
	mockCustomerService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	cRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
}

func Test_get_all_customers_should_return_status_code_500_with_error_message(t *testing.T) {

	// Arrange
	teardown := mockCustomerHandlerSetup(t)
	defer teardown()

	mockCustomerService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedNotFoundError("some data"))

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	cRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func Test_get_customer_should_return_status_code_200(t *testing.T) {

	// Arrange
	teardown := mockCustomerHandlerSetup(t)
	defer teardown()

	dummyCustomer := &dto.CustomerResponse{Id: "1001", Name: "John Doe", City: "New York", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"}
	mockCustomerService.EXPECT().GetCustomer("").Return(dummyCustomer, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers/1001", nil)

	// Act
	recorder := httptest.NewRecorder()
	cRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

}
