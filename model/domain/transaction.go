package domain

import "github.com/shopspring/decimal"

type TransactionType string

const (
	Transfer   TransactionType = "transfer"
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
)

type Transaction struct {
	Id            int             `json:"id"`
	FromAccountId int             `json:"from_account_id"`
	ToAccountId   int             `json:"to_account_id"`
	Amount        decimal.Decimal `json:"amount"`
	Type          TransactionType `json:"type"`
	Description   string          `json:"description"`
}
