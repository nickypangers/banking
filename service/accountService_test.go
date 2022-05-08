package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	realdomain "github.com/nickypangers/banking/domain"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/mocks/domain"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}

}

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

	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedNotFoundError("Unexpected database error"))

	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("failed while validating error for new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {

	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	// Act
	newAccount, appError := service.NewAccount(req)

	// Assert
	if appError != nil {
		t.Error("failed while testing the new account creation")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("failed while testing the new account creation")
	}

}
