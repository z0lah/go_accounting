package dashboard

import (
	"context"
	"time"
)

type DashboardUsecase interface {
	GetSummary(ctx context.Context, start, end time.Time) (*DashboardSummary, error)
}
