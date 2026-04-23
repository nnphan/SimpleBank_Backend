package service

import "errors"

var (
    ErrInvalidCredentials = errors.New("invalid email or password")
    ErrUserInactive       = errors.New("user is inactive")
)

type UserAuthRepository interface {
	FindByEmail(email string) (*User, error)
	UpdatePassword(userID string, hashedPassword string) error
}

type User struct {
	ID       string
	Email    string
	Password string
	IsActive bool
}

