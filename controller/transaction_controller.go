package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMyTransactionsByAccountId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Deposit(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Withdraw(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Transfer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMyTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMyTransactionById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
