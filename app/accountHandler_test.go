package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/mocks/service"
)

var aRouter *mux.Router
var ah AccountHandler
var mockAccountService *service.MockAccountService

func mockAccountHandlerSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAccountService = service.NewMockAccountService(ctrl)
	ah = AccountHandler{mockAccountService}

	aRouter = mux.NewRouter()
	aRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ah.GetAccount)
	return func() {
		aRouter = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_account_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := mockAccountHandlerSetup(t)
	defer teardown()

	dummyAccount := dto.AccountResponse{"12345", "2004", "2000-02-02", "saving", 123.45, "1"}
	mockAccountService.EXPECT().GetAccount("", "").Return(&dummyAccount, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers/2004/accounts/12345", nil)

	// Act
	recorder := httptest.NewRecorder()
	aRouter.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

}
