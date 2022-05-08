package dto

import "testing"

func Test_amount_less_than_5000_should_return_error(t *testing.T) {

	// Arrange
	req := NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	// Act
	appError := req.Validate()

	// Assert
	if appError == nil && appError.Message != "To open a new account you need to deposit at least 5000" {
		t.Error("failed while testing the new account amount validation")
	}

}

func Test_account_type_not_saving_or_checking_should_return_error(t *testing.T) {

	// Arrange
	req := NewAccountRequest{
		CustomerId:  "100",
		AccountType: "something",
		Amount:      6000,
	}

	// Act
	appError := req.Validate()

	// Assert
	if appError == nil && appError.Message != "Account type should be checking or saving" {
		t.Error("failed while testing the new account type validation")
	}
}
