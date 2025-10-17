package repository

import (
	"database/sql"
	"dentify/domain"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

func (r *PostgresUserRepo) CreateUser(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepo) GetUserByID(id int64) (*domain.User, error) {
	var user domain.User
	err := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	row := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1 LIMIT 1", email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
