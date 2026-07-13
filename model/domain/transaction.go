package domain

type TypeTransactions struct {
	Transfer   string
	Deposit    string
	WithDrawal string
}
type Transaction struct {
	Id            int
	FromAccountId string
	ToAccountId   string
	Amount        string
	Type          TypeTransactions
	Description   string
}
