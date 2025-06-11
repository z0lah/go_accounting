package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, u *User) error
}

type UserUsecase interface {
	Register(ctx context.Context, input RegisterInput) (*RegisterResponse, error)
	Login(ctx context.Context, input LoginInput) (*AuthResponse, error)
	GetAll(ctx context.Context) ([]User, error)
}
