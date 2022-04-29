package service

import (
	"time"

	"github.com/nickypangers/banking/domain"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
)

type TransactionService interface {
	Deposit(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
	// Withdraw(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) Deposit(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: "deposit",
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.repo.Save(t)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToNewTransactionResponseDto()
	return &response, nil

}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
