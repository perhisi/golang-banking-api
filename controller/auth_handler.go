package controller

import (
	"encoding/json"
	"golang-banking-api/model/domain"
	"golang-banking-api/service"
	"net/http"
)

type AuthHandler struct {
	authUsecase service.AuthUsecase
}

func NewAuthHandler(au service.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: au}
}

// @Summary		Register user
// @Description	Register a new user account
// @Tags			Auth
// @Accept			json
// @Produce		plain
// @Param			request	body	domain.User	true	"User registration"
// @Success		201
// @Router			/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.authUsecase.Register(req.Name, req.Email, req.Password, req.Role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registrasi sukses"))
}

// @Summary		Login
// @Description	Login user and return access and refresh tokens
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		domain.User	true	"Login credentials"
// @Success		200		{object}	domain.TokenResponse
// @Router			/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	res, err := h.authUsecase.Login(req.Name, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// @Summary		Refresh token
// @Description	Refresh access token using refresh token
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		map[string]string	true	"Refresh token request"
// @Success		200		{object}	map[string]string
// @Router			/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	newAccess, err := h.authUsecase.Refresh(req["refresh_token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": newAccess})
}

// @Summary		Logout
// @Description	Logout user by invalidating refresh token
// @Tags			Auth
// @Accept			json
// @Produce		plain
// @Param			request	body	map[string]string	true	"Logout request"
// @Success		200
// @Router			/logout [post]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	h.authUsecase.Logout(req["refresh_token"])
	w.Write([]byte("Logout sukses"))
}
