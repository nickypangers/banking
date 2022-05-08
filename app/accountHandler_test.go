package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/mocks/service"
)

var aRouter *mux.Router
var ah AccountHandler
var mockAccountService *service.MockAccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAccountService = service.NewMockAccountService(ctrl)
	ah = AccountHandler{mockAccountService}

	aRouter = mux.NewRouter()
	aRouter.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.GetAccount)
	return func() {
		aRouter = nil
		defer ctrl.Finish()
	}

}

func Test_get_account_should_return_account_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyAccount := dto.AccountResponse{
		AccountId:   "91234",
		CustomerId:  "1001",
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: "saving",
		Amount:      6000,
		Status:      "1",
	}

	mockAccountService.EXPECT().GetAccount("1001", "91234").Return(&dummyAccount, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers/1001/account/91234", nil)

	// Act
	recorder := httptest.NewRecorder()
	aRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
}

func Test_get_account_should_return_status_code_500_with_error_message(t *testing.T) {

	// Arrange
	teardown := setup(t)
	defer teardown()

	mockAccountService.EXPECT().GetAccount("1001", "91234").Return(nil, errs.NewNotFoundError("Account not found"))

	request, _ := http.NewRequest(http.MethodGet, "/customers/1001/account/91234", nil)

	// Act
	recorder := httptest.NewRecorder()
	aRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

}
