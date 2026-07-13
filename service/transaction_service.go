package service

import (
	"context"
	"golang-banking-api/model/web"
)

type TransactionService interface {
	Create(ctx context.Context, request web.TransactionCreateRequest) web.TransactionResponse
	Update(ctx context.Context, transactionId int, request web.TransactionUpdateRequest) web.TransactionResponse
	Delete(ctx context.Context, transactionId int)
	FindById(ctx context.Context, transactionId int) web.TransactionResponse
	FindAll(ctx context.Context) []web.TransactionResponse
	FindByAccountId(ctx context.Context, accountId int) []web.TransactionResponse
	Deposit(ctx context.Context, request web.TransactionDepositRequest) web.TransactionResponse
	Withdraw(ctx context.Context, request web.TransactionWithdrawRequest) web.TransactionResponse
	Transfer(ctx context.Context, request web.TransactionTransferRequest) web.TransactionResponse
	FindMyTransactions(ctx context.Context, userId int) []web.TransactionResponse
	FindMyTransactionById(ctx context.Context, userId int, transactionId int) web.TransactionResponse
}
