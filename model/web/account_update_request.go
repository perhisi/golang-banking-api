package web

import (
	"golang-banking-api/model/domain"

	"github.com/shopspring/decimal"
)

type AccountUpdateRequest struct {
	AccountBank string             `validate:"required,min=1,max=100" json:"account_bank"`
	Balance     decimal.Decimal    `validate:"required" json:"balance"`
	AccountType domain.AccountType `json:"account_type"`
}
