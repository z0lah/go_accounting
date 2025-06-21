package balancesheet

import (
	"context"
	"time"
)

type BalanceSheetRepository interface {
	GetAccountsByType(ctx context.Context, accountType string, upTo time.Time) ([]BalanceItem, float64, error)
}
