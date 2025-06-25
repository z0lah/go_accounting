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
