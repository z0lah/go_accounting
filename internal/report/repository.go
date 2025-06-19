package report

import (
	"context"
	"time"
)

type ProfitLossRepository interface {
	GetRevenueAccounts(ctx context.Context, start, end time.Time) ([]AccountAmount, float64, error)
	GetExpenseAccounts(ctx context.Context, start, end time.Time) ([]AccountAmount, float64, error)
}
