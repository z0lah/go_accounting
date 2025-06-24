package dashboard

import (
	"context"
	"time"
)

type DashboardRepository interface {
	GetCashBalance(ctx context.Context) (float64, error)
	GetTotalRevenue(ctx context.Context, start, end time.Time) (float64, error)
	GetTotalExpense(ctx context.Context, start, end time.Time) (float64, error)
	CountJournals(ctx context.Context, start, end time.Time) (int64, error)
}
