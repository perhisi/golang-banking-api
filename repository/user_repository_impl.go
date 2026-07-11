package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-banking-api/helper"
	"golang-banking-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(email, password, name, role) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Name, user.Role)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set email = ?, password = ?, name = ?, role = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Name, user.Role, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select id, email, password, name, role from users where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return user, errors.New("user is not found")
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, email, password, name, role from users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Role)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return users
}
