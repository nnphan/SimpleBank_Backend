package service

import (
	"context"
	db "simplebank/internal/db/sqlc"
	"simplebank/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserUserRepository(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepository.GetUserInfo()
}

func (us *UserService) CreateUser(ctx context.Context, fullName, email, password string) (db.User, error) {
 	user, err := us.userRepository.CreateUser(
		ctx,
		fullName,
		email,
		hashPassword(password),
	)
	return user, err
}

func hashPassword(password string) string {
	if len(password) < 6 {
        return ""
    }
    hash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost, // cost = 10 (recommended)
    )
    if err != nil {
        return ""
    }
    return string(hash)
}



