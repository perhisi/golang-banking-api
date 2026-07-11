package repository

import (
	"context"
	"database/sql"
	"golang-banking-api/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, Transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, Transaction domain.Transaction)
	FindById(ctx context.Context, tx *sql.Tx, TransactionId int) (domain.Transaction, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
}