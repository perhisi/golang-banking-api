package service

import (
	"context"
	"database/sql"
	"golang-banking-api/exception"
	"golang-banking-api/helper"
	"golang-banking-api/model/domain"
	"golang-banking-api/model/web"
	"golang-banking-api/repository"

	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	AccountRepository     repository.AccountRepository
	UserRepository        repository.UserRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, accountRepository repository.AccountRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		AccountRepository:     accountRepository,
		UserRepository:        userRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, request web.TransactionCreateRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := helper.GetUserRole(ctx)
	if role != domain.RoleAdmin {
		account, err := service.AccountRepository.FindById(ctx, tx, request.FromAccountId)
		if err != nil || account.UserId != helper.GetUserId(ctx) {
			panic(exception.NewNotFoundError("from account is not found or does not belong to the user"))
		}
	}

	fromAccount, err := service.AccountRepository.FindById(ctx, tx, request.FromAccountId)
	if err != nil {
		panic(exception.NewNotFoundError("from account is not found"))
	}

	toAccount, err := service.AccountRepository.FindById(ctx, tx, request.ToAccountId)
	if err != nil {
		panic(exception.NewNotFoundError("to account is not found"))
	}

	amount := request.Amount

	switch domain.TransactionType(request.Type) {
	case domain.Transfer:
		fromAccount.Balance = fromAccount.Balance.Sub(amount)
		toAccount.Balance = toAccount.Balance.Add(amount)
	case domain.Deposit:
		toAccount.Balance = toAccount.Balance.Add(amount)
	case domain.Withdrawal:
		fromAccount.Balance = fromAccount.Balance.Sub(amount)
	}

	service.AccountRepository.Update(ctx, tx, fromAccount)
	service.AccountRepository.Update(ctx, tx, toAccount)

	transaction := domain.Transaction{
		FromAccountId: request.FromAccountId,
		ToAccountId:   request.ToAccountId,
		Amount:        request.Amount,
		Type:          domain.TransactionType(request.Type),
		Description:   request.Description,
	}

	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Update(ctx context.Context, transactionId int, request web.TransactionUpdateRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	transaction.FromAccountId = request.FromAccountId
	transaction.ToAccountId = request.ToAccountId
	transaction.Amount = request.Amount
	transaction.Type = request.Type
	transaction.Description = request.Description

	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, transactionId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TransactionRepository.Delete(ctx, tx, transaction)
}

func (service *TransactionServiceImpl) FindById(ctx context.Context, transactionId int) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	role := helper.GetUserRole(ctx)
	if role != domain.RoleAdmin {
		fromAccount, errFrom := service.AccountRepository.FindById(ctx, tx, transaction.FromAccountId)
		toAccount, errTo := service.AccountRepository.FindById(ctx, tx, transaction.ToAccountId)
		userId := helper.GetUserId(ctx)
		if (errFrom != nil || fromAccount.UserId != userId) && (errTo != nil || toAccount.UserId != userId) {
			panic(exception.NewNotFoundError("transaction is not found"))
		}
	}

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindAll(ctx context.Context) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := service.TransactionRepository.FindAll(ctx, tx)

	return helper.ToTransactionResponses(transactions)
}

func (service *TransactionServiceImpl) FindByAccountId(ctx context.Context, accountId int) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := helper.GetUserRole(ctx)
	var transactions []domain.Transaction
	if role == domain.RoleAdmin {
		transactions = service.TransactionRepository.FindByAccountId(ctx, tx, accountId)
	} else {
		userId := helper.GetUserId(ctx)
		userAccounts := service.AccountRepository.FindByUserId(ctx, tx, userId)
		userAccountIds := make(map[int]bool)
		for _, acc := range userAccounts {
			userAccountIds[acc.Id] = true
		}
		if !userAccountIds[accountId] {
			panic(exception.NewNotFoundError("account is not found"))
		}
		transactions = service.TransactionRepository.FindByAccountId(ctx, tx, accountId)
	}

	return helper.ToTransactionResponses(transactions)
}

func (service *TransactionServiceImpl) isAccountOwner(ctx context.Context, account domain.Account) bool {
	return helper.GetUserRole(ctx) == domain.RoleAdmin || account.UserId == helper.GetUserId(ctx)
}

func (service *TransactionServiceImpl) Deposit(ctx context.Context, request web.TransactionDepositRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindById(ctx, tx, request.AccountId)
	if err != nil {
		panic(exception.NewNotFoundError("account is not found"))
	}
	if !service.isAccountOwner(ctx, account) {
		panic(exception.NewNotFoundError("account is not found"))
	}

	account.Balance = account.Balance.Add(request.Amount)
	service.AccountRepository.Update(ctx, tx, account)

	transaction := domain.Transaction{
		ToAccountId: request.AccountId,
		Amount:      request.Amount,
		Type:        domain.Deposit,
		Description: request.Description,
	}
	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Withdraw(ctx context.Context, request web.TransactionWithdrawRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindById(ctx, tx, request.AccountId)
	if err != nil {
		panic(exception.NewNotFoundError("account is not found"))
	}
	if !service.isAccountOwner(ctx, account) {
		panic(exception.NewNotFoundError("account is not found"))
	}
	if account.Balance.LessThan(request.Amount) {
		panic(exception.NewBadRequestError("insufficient balance"))
	}

	account.Balance = account.Balance.Sub(request.Amount)
	service.AccountRepository.Update(ctx, tx, account)

	transaction := domain.Transaction{
		FromAccountId: request.AccountId,
		Amount:        request.Amount,
		Type:          domain.Withdrawal,
		Description:   request.Description,
	}
	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Transfer(ctx context.Context, request web.TransactionTransferRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fromAccount, err := service.AccountRepository.FindById(ctx, tx, request.FromAccountId)
	if err != nil {
		panic(exception.NewNotFoundError("from account is not found"))
	}
	toAccount, err := service.AccountRepository.FindById(ctx, tx, request.ToAccountId)
	if err != nil {
		panic(exception.NewNotFoundError("to account is not found"))
	}

	if !service.isAccountOwner(ctx, fromAccount) {
		panic(exception.NewNotFoundError("from account is not found"))
	}
	if fromAccount.Balance.LessThan(request.Amount) {
		panic(exception.NewBadRequestError("insufficient balance"))
	}

	fromAccount.Balance = fromAccount.Balance.Sub(request.Amount)
	toAccount.Balance = toAccount.Balance.Add(request.Amount)
	service.AccountRepository.Update(ctx, tx, fromAccount)
	service.AccountRepository.Update(ctx, tx, toAccount)

	transaction := domain.Transaction{
		FromAccountId: request.FromAccountId,
		ToAccountId:   request.ToAccountId,
		Amount:        request.Amount,
		Type:          domain.Transfer,
		Description:   request.Description,
	}
	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindMyTransactions(ctx context.Context, userId int) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	accounts := service.AccountRepository.FindByUserId(ctx, tx, userId)
	seen := make(map[int]bool)
	var transactions []domain.Transaction
	for _, account := range accounts {
		accountTransactions := service.TransactionRepository.FindByAccountId(ctx, tx, account.Id)
		for _, t := range accountTransactions {
			if seen[t.Id] {
				continue
			}
			seen[t.Id] = true
			transactions = append(transactions, t)
		}
	}

	return helper.ToTransactionResponses(transactions)
}

func (service *TransactionServiceImpl) FindMyTransactionById(ctx context.Context, userId int, transactionId int) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if helper.GetUserRole(ctx) != domain.RoleAdmin {
		accessible := false
		if transaction.FromAccountId != 0 {
			if acc, err := service.AccountRepository.FindById(ctx, tx, transaction.FromAccountId); err == nil && acc.UserId == userId {
				accessible = true
			}
		}
		if !accessible && transaction.ToAccountId != 0 {
			if acc, err := service.AccountRepository.FindById(ctx, tx, transaction.ToAccountId); err == nil && acc.UserId == userId {
				accessible = true
			}
		}
		if !accessible {
			panic(exception.NewNotFoundError("transaction is not found"))
		}
	}

	return helper.ToTransactionResponse(transaction)
}
