package journal

import (
	"context"

	"github.com/google/uuid"
)

type JournalRepository interface {
	Create(ctx context.Context, journal *Journal) error
	FindAll(ctx context.Context, page, limit int) ([]Journal, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Journal, error)
}
