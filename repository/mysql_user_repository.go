package repository

import (
	"database/sql"
	"golang-banking-api/model/domain"
)

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{db: db}
}

func (m *mysqlUserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (name, password, role, email) VALUES (?, ?, ?, ?)"
	_, err := m.db.Exec(query, user.Name, user.Password, user.Role, user.Email)
	return err
}

func (m *mysqlUserRepository) GetByUsername(name string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, password, role, email FROM users WHERE name = ?"
	err := m.db.QueryRow(query, name).Scan(&user.Id, &user.Name, &user.Password, &user.Role, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlUserRepository) GetById(userId int) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, password, role, email FROM users WHERE id = ?"
	err := m.db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Password, &user.Role, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
