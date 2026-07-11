package web

import "golang-banking-api/model/domain"

type UserCreateRequest struct {
	Email string `validate:"required,email,min=1,max=100" json:"email"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Name string `validate:"required,min=1,max=100" json:"name"`
	Role domain.Role
}