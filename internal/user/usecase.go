package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindAll(ctx context.Context, page int, limit int) ([]User, int64, error)
	Create(ctx context.Context, u *User) error
	FindNotActive(ctx context.Context) ([]User, error)
	UpdateRole(ctx context.Context, id uuid.UUID, input string) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type UserUsecase interface {
	Register(ctx context.Context, input RegisterInput) (*RegisterResponse, error)
	Login(ctx context.Context, input LoginInput) (*AuthResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]UserResponse, int64, error)
	GetNotActive(ctx context.Context) ([]UserResponse, error)
	UpdateRole(ctx context.Context, id string, input UpdateRoleInput) error
	UpdateStatus(ctx context.Context, id string, input UpdateStatusInput) error
}
