package domain

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Role         Role   `json:"role"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
	GetById(userId int) (*User, error)
}
