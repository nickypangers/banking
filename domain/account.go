package domain

import (
	"time"

	"github.com/nickypangers/banking-lib/errs"
	"github.com/nickypangers/banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) statusAsText() string {
	statusAsText := "active"
	if a.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (a Account) ToAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.statusAsText(),
	}
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/nickypangers/banking/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ById(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
