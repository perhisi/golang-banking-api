package controller

import (
	"golang-banking-api/helper"
	"golang-banking-api/model/web"
	"golang-banking-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// @Summary		Create user
// @Description	Create a new user (admin only)
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.UserCreateRequest	true	"User creation request"
// @Success		200		{object}	web.UserResponse
// @Router			/api/admin/users [post]
func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Update user
// @Description	Update user by ID (admin only)
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			userId	path		int						true	"User ID"
// @Param			request	body		web.UserUpdateRequest	true	"User update request"
// @Success		200		{object}	web.UserResponse
// @Router			/api/admin/users/{userId} [put]
func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Delete user
// @Description	Delete user by ID (admin only)
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			userId	path	int	true	"User ID"
// @Success		200
// @Router			/api/admin/users/{userId} [delete]
func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get user by ID
// @Description	Get user details by ID (admin only)
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			userId	path		int	true	"User ID"
// @Success		200		{object}	web.UserResponse
// @Router			/api/admin/users/{userId} [get]
func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		List users
// @Description	Get all users (admin only)
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{array}	web.UserResponse
// @Router			/api/admin/users [get]
func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Get current user profile
// @Description	Get authenticated user profile
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	web.UserResponse
// @Router			/api/user/profile [get]
func (controller *UserControllerImpl) GetMe(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.GetUserId(request.Context())
	userResponse := controller.UserService.GetMe(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary		Update current user profile
// @Description	Update authenticated user profile
// @Tags			Users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		web.UserUpdateMeRequest	true	"User update request"
// @Success		200		{object}	web.UserResponse
// @Router			/api/user/profile [put]
func (controller *UserControllerImpl) UpdateMe(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateMeRequest := web.UserUpdateMeRequest{}
	helper.ReadFromRequestBody(request, &userUpdateMeRequest)

	userResponse := controller.UserService.UpdateMe(request.Context(), userUpdateMeRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}