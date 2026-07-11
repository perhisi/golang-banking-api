package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-banking-api/helper"
	"golang-banking-api/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "insert into transaction(from_account_id, to_account_id, amount, type, description) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transaction.FromAccountId, transaction.ToAccountId, transaction.Amount, transaction.Type, transaction.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	transaction.Id = int(id)
	return transaction
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "update transaction set from_account_id = ?, to_account_id = ?, amount = ?, type = ?, description = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, transaction.FromAccountId, transaction.ToAccountId, transaction.Amount, transaction.Type, transaction.Description, transaction.Id)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) {
	SQL := "delete from transaction where id = ?"
	_, err := tx.ExecContext(ctx, SQL, transaction.Id)
	helper.PanicIfError(err)
}

func (repository *TransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error) {
	SQL := "select id, from_account_id, to_account_id, amount, type, description from transaction where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(&transaction.Id, &transaction.FromAccountId, &transaction.ToAccountId, &transaction.Amount, &transaction.Type, &transaction.Description)
		helper.PanicIfError(err)
		return transaction, nil
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return transaction, errors.New("transaction is not found")
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction {
	SQL := "select id, from_account_id, to_account_id, amount, type, description from transaction"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.FromAccountId, &transaction.ToAccountId, &transaction.Amount, &transaction.Type, &transaction.Description)
		helper.PanicIfError(err)
		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return transactions
}
