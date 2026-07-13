package web

import (
	"golang-banking-api/model/domain"

	"github.com/shopspring/decimal"
)

type TransactionResponse struct {
	Id            int                    `json:"id"`
	FromAccountId int                    `json:"from_account_id"`
	ToAccountId   int                    `json:"to_account_id"`
	Amount        decimal.Decimal        `json:"amount"`
	Type          domain.TransactionType `json:"type"`
	Description   string                 `json:"description"`
}
