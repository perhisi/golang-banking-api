package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-banking-api/helper"
	"golang-banking-api/model/domain"

	"github.com/shopspring/decimal"
)

type AccountRepositoryImpl struct {
}

func NewAccountRepository() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{}
}

func (repository *AccountRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account {
	SQL := "insert into accounts(user_id, account_bank, balance, account_type) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, account.UserId, account.AccountBank, account.Balance.String(), account.AccountType)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	account.Id = int(id)
	return account
}

func (repository *AccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account {
	SQL := "update accounts set account_bank = ?, balance = ?, account_type = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, account.AccountBank, account.Balance.String(), account.AccountType, account.Id)
	helper.PanicIfError(err)

	return account
}

func (repository *AccountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, account domain.Account) {
	SQL := "delete from accounts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, account.Id)
	helper.PanicIfError(err)
}

func (repository *AccountRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, accountId int) (domain.Account, error) {
	SQL := "select id, user_id, account_bank, balance, account_type from accounts where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, accountId)
	helper.PanicIfError(err)
	defer rows.Close()

	account := domain.Account{}
	if rows.Next() {
		var balance string
		err := rows.Scan(&account.Id, &account.UserId, &account.AccountBank, &balance, &account.AccountType)
		helper.PanicIfError(err)
		account.Balance, err = decimal.NewFromString(balance)
		helper.PanicIfError(err)
		return account, nil
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return account, errors.New("account is not found")

}

func (repository *AccountRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Account {
	SQL := "select id, user_id, account_bank, balance, account_type from accounts"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var accounts []domain.Account
	for rows.Next() {
		account := domain.Account{}
		var balance string
		err := rows.Scan(&account.Id, &account.UserId, &account.AccountBank, &balance, &account.AccountType)
		helper.PanicIfError(err)
		account.Balance, err = decimal.NewFromString(balance)
		helper.PanicIfError(err)
		accounts = append(accounts, account)
	}
	err = rows.Err()
	helper.PanicIfError(err)
	return accounts
}

func (repository *AccountRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Account {
	SQL := "select id, user_id, account_bank, balance, account_type from accounts where user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var accounts []domain.Account
	for rows.Next() {
		account := domain.Account{}
		var balance string
		err := rows.Scan(&account.Id, &account.UserId, &account.AccountBank, &balance, &account.AccountType)
		helper.PanicIfError(err)
		account.Balance, err = decimal.NewFromString(balance)
		helper.PanicIfError(err)
		accounts = append(accounts, account)
	}
	err = rows.Err()
	helper.PanicIfError(err)
	return accounts
}
