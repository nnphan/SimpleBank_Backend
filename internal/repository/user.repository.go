package repository

import (
	"context"
	db "simplebank/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository struct{q *db.Queries}

func NewUserUserRepository() *UserRepository {
	return &UserRepository{ q: &db.Queries{} }
}

func (ur *UserRepository) GetUserInfo() string {
	return "Phan"
}

func (ur *UserRepository) CreateUser(ctx context.Context, fullName, email, passwordHash string) (db.User, error) {
	return ur.q.CreateUser(ctx, db.CreateUserParams{
		ID: uuid.New(),
		FullName: fullName,
		Email: email,
		PasswordHash: passwordHash,
	})
}
