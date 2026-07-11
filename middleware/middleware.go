package middleware

import (
	"context"
	"golang-banking-api/model/domain"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var publicPaths = map[string]bool{
	"/register": true,
	"/login":    true,
	"/refresh":  true,
	"/logout":   true,
}

func AuthRoleMiddleware(allowedRoles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if publicPaths[r.URL.Path] {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || len(authHeader) < 8 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			tokenStr := authHeader[7:]

			claims := &domain.JWTClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				return domain.JWTSecret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			roleValid := false
			for _, role := range allowedRoles {
				if claims.Role == role {
					roleValid = true
					break
				}
			}

			if !roleValid {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Simpan data klaim ke context jika handler butuh ID pengakses
			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
