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
