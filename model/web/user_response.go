package web

type UserResponse struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Role        string `json:"role"`
	RefreshToken string `json:"refresh_token"`
}
