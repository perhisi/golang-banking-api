package web

import "golang-banking-api/model/domain"

type AccountUpdateRequest struct {
	AccountBank string             `validate:"required,min=1,max=100" json:"account_bank"`
	Balance     string             `validate:"required,min=1,decimal=2" json:"balance"`
	AccountType domain.AccountType `json:"account_type"`
}
