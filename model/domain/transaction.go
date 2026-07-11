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
	Amount        float64
	Type          TypeTransactions
	Description   string
}
