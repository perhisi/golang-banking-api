package web

import "golang-banking-api/model/domain"

type AccountUpdateRequest struct {
	AccountBank string             `validate:"required,min=1,max=100" json:"account_bank"`
	Balance     float64            `validate:"required,min=1" json:"balance"`
	AccountType domain.AccountType `validate:"required" json:"account_type"`
}
