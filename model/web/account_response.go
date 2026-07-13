package web

type AccountResponse struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	AccountBank string `json:"account_bank"`
	Balance     string `json:"balance"`
	AccountType string `json:"account_type"`
}
