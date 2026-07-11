package repository

import (
	"database/sql"
	"golang-banking-api/model/domain"
)

type mysqlTokenRepository struct {
	db *sql.DB
}

func NewMySQLTokenRepository(db *sql.DB) domain.TokenRepository {
	return &mysqlTokenRepository{db: db}
}

func (m *mysqlTokenRepository) Store(t *domain.RefreshToken) error {
	query := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES (?, ?, ?)"
	_, err := m.db.Exec(query, t.UserID, t.Token, t.ExpiresAt)
	return err
}

func (m *mysqlTokenRepository) Get(tokenStr string) (*domain.RefreshToken, error) {
	var t domain.RefreshToken
	query := "SELECT user_id, token, expires_at FROM refresh_tokens WHERE token = ?"
	err := m.db.QueryRow(query, tokenStr).Scan(&t.UserID, &t.Token, &t.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m *mysqlTokenRepository) Delete(tokenStr string) error {
	query := "DELETE FROM refresh_tokens WHERE token = ?"
	_, err := m.db.Exec(query, tokenStr)
	return err
}
