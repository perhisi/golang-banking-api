package domain

import "context"

type contextKey string

var (
	JWTContextUserIDKey  contextKey = "user_id"
	JWTContextUserRoleKey contextKey = "user_role"
)

func ContextWithUser(ctx context.Context, userID int, role Role) context.Context {
	ctx = context.WithValue(ctx, JWTContextUserIDKey, userID)
	return context.WithValue(ctx, JWTContextUserRoleKey, role)
}

func UserIDFromContext(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	val, ok := ctx.Value(JWTContextUserIDKey).(int)
	if !ok {
		return 0
	}
	return val
}

func RoleFromContext(ctx context.Context) Role {
	if ctx == nil {
		return ""
	}
	val, ok := ctx.Value(JWTContextUserRoleKey).(Role)
	if !ok {
		return ""
	}
	return val
}
