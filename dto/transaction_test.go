package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdraw(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}
	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type should be deposit or withdraw" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		Amount: -100,
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Amount cannot be negative" {
		t.Error("Invalid message while testing amount")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing amount")
	}

}
