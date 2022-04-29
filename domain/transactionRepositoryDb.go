package domain

import (
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}
