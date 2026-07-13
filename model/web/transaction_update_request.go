package web

import (
	"golang-banking-api/model/domain"

	"github.com/shopspring/decimal"
)

type TransactionUpdateRequest struct {
	FromAccountId int                    `validate:"required,min=1,max=100" json:"from_account_id"`
	ToAccountId   int                    `validate:"required,min=1,max=100" json:"to_account_id"`
	Amount        decimal.Decimal        `validate:"required" json:"amount"`
	Type          domain.TransactionType `json:"type"`
	Description   string                 `json:"description"`
}
