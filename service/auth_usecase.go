package service

import (
	"errors"
	"golang-banking-api/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = domain.JWTSecret

type AuthUsecase interface {
	Register(name, email, password string, role domain.Role) error
	Login(name, password string) (*domain.TokenResponse, error)
	Refresh(refreshTokenStr string) (string, error)
	Logout(refreshTokenStr string) error
}

type authUsecase struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
}

func NewAuthUsecase(ur domain.UserRepository, tr domain.TokenRepository) AuthUsecase {
	return &authUsecase{userRepo: ur, tokenRepo: tr}
}

func (u *authUsecase) Register(name, email, password string, role domain.Role) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if role == "" {
		role = domain.RoleUser
	}
	user := &domain.User{Name: name, Email: email, Password: string(hashedPassword), Role: role}
	return u.userRepo.Create(user)
}

func (u *authUsecase) Login(name, password string) (*domain.TokenResponse, error) {
	user, err := u.userRepo.GetByUsername(name)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username atau password salah")
	}

	// Buat Access Token
	accessToken, err := u.generateAccessToken(user.Id, user.Role)
	if err != nil {
		return nil, err
	}

	// Buat & Simpan Refresh Token
	rfTokenStr := uuid.New().String()
	rfToken := &domain.RefreshToken{
		UserID:    user.Id,
		Token:     rfTokenStr,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := u.tokenRepo.Store(rfToken); err != nil {
		return nil, err
	}

	return &domain.TokenResponse{AccessToken: accessToken, RefreshToken: rfTokenStr}, nil
}

func (u *authUsecase) Refresh(rfTokenStr string) (string, error) {
	storedToken, err := u.tokenRepo.Get(rfTokenStr)
	if err != nil {
		return "", errors.New("token tidak ditemukan")
	}

	if time.Now().After(storedToken.ExpiresAt) {
		u.tokenRepo.Delete(rfTokenStr)
		return "", errors.New("token kedaluwarsa")
	}

	// Untuk mendapat Role, idealnya ambil dari userRepo lewat storedToken.UserID
	// Di sini disederhanakan dengan asumsi penanganan kueri kustom jika dibutuhkan
	return u.generateAccessToken(storedToken.UserID, domain.RoleUser) // Sesuaikan role aslinya
}

func (u *authUsecase) Logout(rfTokenStr string) error {
	return u.tokenRepo.Delete(rfTokenStr)
}

func (u *authUsecase) generateAccessToken(userID int, role domain.Role) (string, error) {
	claims := &domain.JWTClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}
