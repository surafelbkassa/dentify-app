package usecase

import (
	"dentify/domain"
	"dentify/repository"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *UserUsecase) Signup(user *domain.User) error {
	// TODO Hash Password
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(id int64) (*domain.User, error) {
	return u.UserRepo.GetUserByID(id)
}
