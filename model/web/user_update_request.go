package web

import "golang-banking-api/model/domain"

type UserUpdateRequest struct {
	Id int `validate:"required"`
	Email string `validate:"required,email,min=1,max=100" json:"email"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Name string `validate:"required,min=1,max=100" json:"name"`
	Role domain.Role
}