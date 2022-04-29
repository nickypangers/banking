package service

import (
	"github.com/nickypangers/banking/domain"
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
)

type TransactionService interface {
	Deposit(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
	Withdraw(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}
