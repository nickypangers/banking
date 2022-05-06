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
	return func() {
		cRouter = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := mockCustomerHandlerSetup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "John Doe", "New York", "110011", "2000-01-01", "1"},
		{"1001", "Rob Holding", "London", "12304", "1994-03-19", "1"},
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

func Test_should_return_status_code_500_with_error_message(t *testing.T) {

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
