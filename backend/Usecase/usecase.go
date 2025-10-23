package usecase

import (
	"dentify/domain"
	"dentify/repository"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *UserUsecase) Signup(user *domain.User) error {
	// TODO Hash Password
	if user == nil {
		return errors.New("user is nil")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password cannot be empty")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	if existing, _ := u.UserRepo.GetUserByEmail(user.Email); existing != nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(id int64) (*domain.User, error) {
	return u.UserRepo.GetUserByID(id)
}

func (u *UserUsecase) Login(email, password string) (*domain.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
