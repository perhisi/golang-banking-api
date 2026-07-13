package controller

import (
	"golang-banking-api/helper"
	"golang-banking-api/model/web"
	"golang-banking-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

// @Summary		Create transaction
// @Description	Create a new transaction
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.TransactionCreateRequest	true	"Transaction creation request"
// @Success		200		{object}	web.TransactionResponse
// @Router			/api/admin/transactions [post]
func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionCreateRequest)

	transactionResponse := controller.TransactionService.Create(request.Context(), transactionCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Update transaction
// @Description	Update transaction by ID
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			transactionId	path		int								true	"Transaction ID"
// @Param			request			body		web.TransactionUpdateRequest	true	"Transaction update request"
// @Success		200				{object}	web.TransactionResponse
// @Router			/api/admin/transactions/{transactionId} [put]
func (controller *TransactionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionUpdateRequest := web.TransactionUpdateRequest{}
	helper.ReadFromRequestBody(request, &transactionUpdateRequest)

	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.Update(request.Context(), id, transactionUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Delete transaction
// @Description	Delete transaction by ID
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			transactionId	path	int	true	"Transaction ID"
// @Success		200
// @Router			/api/admin/transactions/{transactionId} [delete]
func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	controller.TransactionService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get transaction by ID
// @Description	Get transaction details by ID (admin only)
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			transactionId	path		int	true	"Transaction ID"
// @Success		200				{object}	web.TransactionResponse
// @Router			/api/admin/transactions/{transactionId} [get]
func (controller *TransactionControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		List transactions
// @Description	Get all transactions
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{array}	web.TransactionResponse
// @Router			/api/admin/transactions [get]
func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionResponses := controller.TransactionService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Deposit
// @Description	Deposit an amount into an account
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.TransactionDepositRequest	true	"Deposit request"
// @Success		200		{object}	web.TransactionResponse
// @Router			/api/user/transactions/deposit [post]
func (controller *TransactionControllerImpl) Deposit(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionDepositRequest := web.TransactionDepositRequest{}
	helper.ReadFromRequestBody(request, &transactionDepositRequest)

	transactionResponse := controller.TransactionService.Deposit(request.Context(), transactionDepositRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Withdraw
// @Description	Withdraw an amount from an account
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.TransactionWithdrawRequest	true	"Withdraw request"
// @Success		200		{object}	web.TransactionResponse
// @Router			/api/user/transactions/withdraw [post]
func (controller *TransactionControllerImpl) Withdraw(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionWithdrawRequest := web.TransactionWithdrawRequest{}
	helper.ReadFromRequestBody(request, &transactionWithdrawRequest)

	transactionResponse := controller.TransactionService.Withdraw(request.Context(), transactionWithdrawRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Transfer
// @Description	Transfer an amount between accounts
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.TransactionTransferRequest	true	"Transfer request"
// @Success		200		{object}	web.TransactionResponse
// @Router			/api/user/transactions/transfer [post]
func (controller *TransactionControllerImpl) Transfer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTransferRequest := web.TransactionTransferRequest{}
	helper.ReadFromRequestBody(request, &transactionTransferRequest)

	transactionResponse := controller.TransactionService.Transfer(request.Context(), transactionTransferRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		List my transactions
// @Description	List all transactions across the authenticated user's accounts
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{array}	web.TransactionResponse
// @Router			/api/user/transactions [get]
func (controller *TransactionControllerImpl) GetMyTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	transactionResponses := controller.TransactionService.FindMyTransactions(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get my transaction by ID
// @Description	Get transaction details by ID for the authenticated user
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			transactionId	path		int	true	"Transaction ID"
// @Success		200				{object}	web.TransactionResponse
// @Router			/api/user/transactions/{transactionId} [get]
func (controller *TransactionControllerImpl) GetMyTransactionById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.FindMyTransactionById(request.Context(), userId, id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get my transactions by account ID
// @Description	Get transactions for a specific account belonging to the authenticated user
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path	int	true	"Account ID"
// @Success		200			{array}	web.TransactionResponse
// @Router			/api/user/transactions/{accountId} [get]
func (controller *TransactionControllerImpl) GetMyTransactionsByAccountId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	transactionResponses := controller.TransactionService.FindByAccountId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
