package helper

import (
	"golang-banking-api/model/domain"
	"golang-banking-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:           user.Id,
		Email:        user.Email,
		Password:     user.Password,
		Name:         user.Name,
		Role:         string(user.Role),
		RefreshToken: user.RefreshToken,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToAccountResponse(account domain.Account) web.AccountResponse {
	return web.AccountResponse{
		Id:          account.Id,
		UserId:      account.UserId,
		AccountBank: account.AccountBank,
		Balance:     account.Balance,
		AccountType: string(account.AccountType),
	}
}

func ToAccountResponses(accounts []domain.Account) []web.AccountResponse {
	var accountResponses []web.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, ToAccountResponse(account))
	}
	return accountResponses
}
