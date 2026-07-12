package repository

import (
	"context"
	"database/sql"
	"golang-banking-api/model/domain"
)

type AccountRepository interface {
	Save(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account
	Update(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account
	Delete(ctx context.Context, tx *sql.Tx, account domain.Account)
	FindById(ctx context.Context, tx *sql.Tx, accountId int) (domain.Account, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Account
}
