package user

import (
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, input RegisterInput) (*RegisterResponse, error)
	Login(ctx context.Context, input LoginInput) (*AuthResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]UserResponse, int64, error)
	GetNotActive(ctx context.Context) ([]UserResponse, error)
	UpdateRole(ctx context.Context, id string, input UpdateRoleInput) error
	UpdateStatus(ctx context.Context, id string, input UpdateStatusInput) error
}
