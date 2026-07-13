package web

import "github.com/shopspring/decimal"

type AccountResponse struct {
	Id          int             `json:"id"`
	UserId      int             `json:"user_id"`
	AccountBank string          `json:"account_bank"`
	Balance     decimal.Decimal `json:"balance"`
	AccountType string          `json:"account_type"`
}
