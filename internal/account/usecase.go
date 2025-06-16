package account

import (
	"context"

	"github.com/google/uuid"
)

type AccountRepository interface {
	Create(ctx context.Context, a *Account) error
	FindByID(ctx context.Context, id uuid.UUID) (*Account, error)
	FindAll(ctx context.Context, page, limit int) ([]Account, int64, error)
	UpdateAccount(ctx context.Context, a *Account) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AccountUsecase interface {
	Create(ctx context.Context, input CreateAccountInput) (*AccountResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]AccountResponse, int64, error)
	GetByID(ctx context.Context, id string) (*AccountResponse, error)
	Update(ctx context.Context, id string, input UpdateAccountInput) (*AccountResponse, error)
	Delete(ctx context.Context, id string) error
}
