package balancesheet

import (
	"context"
	"time"
)

type balanceSheetUsecase struct {
	repo BalanceSheetRepository
}

func NewBalanceSheetUsecase(repo BalanceSheetRepository) BalanceSheetUsecase {
	return &balanceSheetUsecase{repo: repo}
}

func (u *balanceSheetUsecase) GetBalanceSheet(ctx context.Context, date time.Time) (*BalanceSheet, error) {
	assets, totalAssets, err := u.repo.GetAccountsByType(ctx, "asset", date)
	if err != nil {
		return nil, err
	}

	liabilities, totalLiabilities, err := u.repo.GetAccountsByType(ctx, "liability", date)
	if err != nil {
		return nil, err
	}

	equity, totalEquity, err := u.repo.GetAccountsByType(ctx, "equity", date)
	if err != nil {
		return nil, err
	}

	return &BalanceSheet{
		Assets:         assets,
		Liability:      liabilities,
		Equity:         equity,
		TotalAsset:     totalAssets,
		TotalLiability: totalLiabilities,
		TotalEquity:    totalEquity,
	}, nil
}
