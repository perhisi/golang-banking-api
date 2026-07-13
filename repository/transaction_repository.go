package repository

import (
	"context"
	"database/sql"
	"golang-banking-api/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction)
	FindById(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
	FindByAccountId(ctx context.Context, tx *sql.Tx, accountId int) []domain.Transaction
}
