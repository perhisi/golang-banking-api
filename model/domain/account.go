package domain

type AccountType string

const (
	Savings  AccountType = "savings"
	Checking AccountType = "checking"
)

type Account struct {
	Id          int         `json:"id"`
	UserId      int         `json:"user_id"`
	AccountBank string      `json:"account_bank"`
	Balance     string      `json:"balance"`
	AccountType AccountType `json:"account_type"`
}
