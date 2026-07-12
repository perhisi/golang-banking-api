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

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	UserRepository    repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAccountService(accountRepository repository.AccountRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AccountService {
	return &AccountServiceImpl{
		AccountRepository: accountRepository,
		UserRepository:    userRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AccountServiceImpl) Create(ctx context.Context, request web.AccountCreateRequest) web.AccountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	account := domain.Account{
		UserId:      request.UserId,
		AccountBank: request.AccountBank,
		Balance:     request.Balance,
		AccountType: domain.AccountType(request.AccountType),
	}

	account = service.AccountRepository.Save(ctx, tx, account)

	return helper.ToAccountResponse(account)
}

func (service *AccountServiceImpl) Update(ctx context.Context, accountId int, request web.AccountUpdateRequest) web.AccountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindById(ctx, tx, accountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	account.AccountBank = request.AccountBank
	account.Balance = request.Balance
	account.AccountType = request.AccountType

	account = service.AccountRepository.Update(ctx, tx, account)

	return helper.ToAccountResponse(account)
}

func (service *AccountServiceImpl) Delete(ctx context.Context, accountId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindById(ctx, tx, accountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AccountRepository.Delete(ctx, tx, account)
}

func (service *AccountServiceImpl) FindById(ctx context.Context, accountId int) web.AccountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindById(ctx, tx, accountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAccountResponse(account)
}

func (service *AccountServiceImpl) FindAll(ctx context.Context) []web.AccountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	accounts := service.AccountRepository.FindAll(ctx, tx)

	return helper.ToAccountResponses(accounts)
}

func (service *AccountServiceImpl) FindByUserId(ctx context.Context, userId int) []web.AccountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	accounts := service.AccountRepository.FindByUserId(ctx, tx, userId)

	return helper.ToAccountResponses(accounts)
}
