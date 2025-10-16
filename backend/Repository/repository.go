package repository

import "dentify/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
}
