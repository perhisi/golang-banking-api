package web

import "github.com/shopspring/decimal"

type TransactionDepositRequest struct {
	AccountId   int             `validate:"required,min=1" json:"account_id"`
	Amount      decimal.Decimal `validate:"required" json:"amount"`
	Description string          `json:"description"`
}

type TransactionWithdrawRequest struct {
	AccountId   int             `validate:"required,min=1" json:"account_id"`
	Amount      decimal.Decimal `validate:"required" json:"amount"`
	Description string          `json:"description"`
}

type TransactionTransferRequest struct {
	FromAccountId int             `validate:"required,min=1" json:"from_account_id"`
	ToAccountId   int             `validate:"required,min=1" json:"to_account_id"`
	Amount        decimal.Decimal `validate:"required" json:"amount"`
	Description   string          `json:"description"`
}
