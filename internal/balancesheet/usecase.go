package balancesheet

import (
	"context"
	"time"
)

type BalanceSheetUsecase interface {
	GetBalanceSheet(ctx context.Context, date time.Time) (*BalanceSheet, error)
}
