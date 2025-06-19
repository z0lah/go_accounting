package report

import (
	"context"
	"time"
)

type profitLossUsecase struct {
	repo ProfitLossRepository
}

func NewProfitLossUsecase(repo ProfitLossRepository) ProfitLossUsecase {
	return &profitLossUsecase{repo: repo}
}

func (u *profitLossUsecase) GetReport(ctx context.Context, start, end time.Time) (*ProfitLossResponse, error) {
	revenues, totalRevenue, err := u.repo.GetRevenueAccounts(ctx, start, end)
	if err != nil {
		return nil, err
	}

	expenses, totalExpense, err := u.repo.GetExpenseAccounts(ctx, start, end)
	if err != nil {
		return nil, err
	}

	netProfit := totalRevenue - totalExpense

	return &ProfitLossResponse{
		Revenue:      revenues,
		Expense:      expenses,
		TotalRevenue: totalRevenue,
		TotalExpense: totalExpense,
		NetProfit:    netProfit,
	}, nil
}
