package web

import "golang-banking-api/model/domain"

type AccountCreateRequest struct {
	UserId      int                `validate:"required,min=1,max=100" json:"user_id"`
	AccountBank string             `validate:"required,min=1,max=100" json:"account_bank"`
	Balance     string             `validate:"required,min=1,decimal=2" json:"balance"`
	AccountType domain.AccountType `json:"account_type"`
}
