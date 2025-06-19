package report

import (
	"context"
	"time"
)

type ProfitLossUsecase interface {
	GetReport(ctx context.Context, start, end time.Time) (*ProfitLossResponse, error)
}
