package service

import (
	"time"

	"github.com/nickypangers/banking-lib/errs"
	"github.com/nickypangers/banking/domain"
	"github.com/nickypangers/banking/dto"
)

//go:generate mockgen -destination=../mocks/service/mockAccountService.go -package=service github.com/nickypangers/banking/service AccountService
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetAccount(string, string) (*dto.AccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) GetAccount(customerId, accountId string) (*dto.AccountResponse, *errs.AppError) {
	account, err := s.repo.ById(accountId)
	if err != nil {
		return nil, err
	}
	response := account.ToAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}

}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.ById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient funds")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
