package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("kunci_rahasia_super_aman")

type RefreshToken struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type JWTClaims struct {
	UserID int  `json:"user_id"`
	Role   Role `json:"role"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenRepository adalah blueprint manajemen token di DB
type TokenRepository interface {
	Store(token *RefreshToken) error
	Get(tokenStr string) (*RefreshToken, error)
	Delete(tokenStr string) error
}
