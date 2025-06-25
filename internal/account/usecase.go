package account

import (
	"context"
)

type AccountUsecase interface {
	Create(ctx context.Context, input CreateAccountInput) (*AccountResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]AccountResponse, int64, error)
	GetByID(ctx context.Context, id string) (*AccountResponse, error)
	Update(ctx context.Context, id string, input UpdateAccountInput) (*AccountResponse, error)
	Delete(ctx context.Context, id string) error
}
