package domain

import "github.com/shopspring/decimal"

type AccountType string

const (
	Savings  AccountType = "savings"
	Checking AccountType = "checking"
)

type Account struct {
	Id          int             `json:"id"`
	UserId      int             `json:"user_id"`
	AccountBank string          `json:"account_bank"`
	Balance     decimal.Decimal `json:"balance"`
	AccountType AccountType     `json:"account_type"`
}
