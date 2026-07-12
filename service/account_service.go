package service

import (
	"context"
	"golang-banking-api/model/web"
)

type AccountService interface {
	Create(ctx context.Context, request web.AccountCreateRequest) web.AccountResponse
	Update(ctx context.Context, accountId int, request web.AccountUpdateRequest) web.AccountResponse
	Delete(ctx context.Context, accountId int)
	FindById(ctx context.Context, accountId int) web.AccountResponse
	FindAll(ctx context.Context) []web.AccountResponse
	FindByUserId(ctx context.Context, userId int) []web.AccountResponse
}
