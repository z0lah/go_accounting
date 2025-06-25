package journal

import (
	"context"
)

type JournalUsecase interface {
	Create(ctx context.Context, input CreateJournalInput) (*JournalResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]JournalResponse, int64, error)
	GetByID(ctx context.Context, id string) (*JournalResponse, error)
}
