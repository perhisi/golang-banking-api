package helper

import (
	"context"
	"golang-banking-api/model/domain"
)

func GetUserId(ctx context.Context) int {
	return domain.UserIDFromContext(ctx)
}

func GetUserRole(ctx context.Context) domain.Role {
	return domain.RoleFromContext(ctx)
}
