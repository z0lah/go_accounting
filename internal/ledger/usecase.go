package ledger

import (
	"context"
	"time"
)

type LedgerUseCase interface {
	GetAll(ctx context.Context, startDate, endDate time.Time) ([]LedgerResponse, error)
}
