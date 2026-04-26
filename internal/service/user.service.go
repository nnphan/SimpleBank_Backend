package service

import (
	"simplebank/internal/repository"
	"simplebank/response"
)

type IUserService interface {
	Register(fullName, email, password string) int
}

type userService struct {
	userRepo repository.IUserRepository
	//...
}


func NewUserService(
	userRepo repository.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(fullName string, email string, password string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.StatusUserHasExist
	}
	return response.StatusOK
}
