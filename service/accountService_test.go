package service

import (
	"testing"

	"github.com/nickypangers/banking/dto"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {

	// Arrange
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	// Act
	_, appError := service.NewAccount(request)

	// Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	
}
