package controller

import (
	"golang-banking-api/exception"
	"golang-banking-api/helper"
	"golang-banking-api/model/web"
	"golang-banking-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

// @Summary		Create account
// @Description	Create a new account
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.AccountCreateRequest	true	"Account creation request"
// @Success		200		{object}	web.AccountResponse
// @Router			/api/admin/accounts [post]
func (controller *AccountControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountCreateRequest := web.AccountCreateRequest{}
	helper.ReadFromRequestBody(request, &accountCreateRequest)

	accountResponse := controller.AccountService.Create(request.Context(), accountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Update account
// @Description	Update account by ID
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path		int							true	"Account ID"
// @Param			request		body		web.AccountUpdateRequest	true	"Account update request"
// @Success		200			{object}	web.AccountResponse
// @Router			/api/admin/accounts/{accountId} [put]
func (controller *AccountControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountUpdateRequest := web.AccountUpdateRequest{}
	helper.ReadFromRequestBody(request, &accountUpdateRequest)

	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.Update(request.Context(), id, accountUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Delete account
// @Description	Delete account by ID
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path	int	true	"Account ID"
// @Success		200
// @Router			/api/admin/accounts/{accountId} [delete]
func (controller *AccountControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	controller.AccountService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get account by ID
// @Description	Get account details by ID (admin only)
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path		int	true	"Account ID"
// @Success		200			{object}	web.AccountResponse
// @Router			/api/admin/accounts/{accountId} [get]
func (controller *AccountControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		List accounts
// @Description	Get all accounts
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{array}	web.AccountResponse
// @Router			/api/admin/accounts [get]
func (controller *AccountControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountResponses := controller.AccountService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get my accounts
// @Description	Get accounts belonging to the authenticated user
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{array}	web.AccountResponse
// @Router			/api/user/accounts [get]
func (controller *AccountControllerImpl) GetMyAccounts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	accountResponses := controller.AccountService.FindByUserId(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Create my account
// @Description	Create a new account for the authenticated user
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.AccountCreateRequest	true	"Account creation request"
// @Success		200		{object}	web.AccountResponse
// @Router			/api/user/accounts [post]
func (controller *AccountControllerImpl) CreateMyAccount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountCreateRequest := web.AccountCreateRequest{}
	helper.ReadFromRequestBody(request, &accountCreateRequest)

	accountCreateRequest.UserId = helper.GetUserId(request.Context())

	accountResponse := controller.AccountService.Create(request.Context(), accountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Update my account
// @Description	Update an account belonging to the authenticated user
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path		int							true	"Account ID"
// @Param			request		body		web.AccountUpdateRequest	true	"Account update request"
// @Success		200			{object}	web.AccountResponse
// @Router			/api/user/accounts/{accountId} [patch]
func (controller *AccountControllerImpl) UpdateMyAccount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.FindById(request.Context(), id)
	if accountResponse.UserId != userId {
		panic(exception.NewNotFoundError("account is not found"))
	}

	accountUpdateRequest := web.AccountUpdateRequest{}
	helper.ReadFromRequestBody(request, &accountUpdateRequest)

	accountResponse = controller.AccountService.Update(request.Context(), id, accountUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Delete my account
// @Description	Delete an account belonging to the authenticated user
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path	int	true	"Account ID"
// @Success		200
// @Router			/api/user/accounts/{accountId} [delete]
func (controller *AccountControllerImpl) DeleteMyAccount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.FindById(request.Context(), id)
	if accountResponse.UserId != userId {
		panic(exception.NewNotFoundError("account is not found"))
	}

	controller.AccountService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get my account by ID
// @Description	Get account details by ID for the authenticated user
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			accountId	path		int	true	"Account ID"
// @Success		200			{object}	web.AccountResponse
// @Router			/api/user/accounts/{accountId} [get]
func (controller *AccountControllerImpl) GetMyAccountById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.FindById(request.Context(), id)
	if accountResponse.UserId != userId {
		panic(exception.NewNotFoundError("account is not found"))
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
