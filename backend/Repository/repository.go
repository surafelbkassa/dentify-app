package repository

import "dentify/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	LoginUser(email, password string) (*domain.User, error)
}
