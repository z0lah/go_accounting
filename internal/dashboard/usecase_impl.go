package dashboard

import (
	"context"
	"time"
)

type dashboardUsecase struct {
	repo DashboardRepository
}

func NewDashboardUsecase(repo DashboardRepository) DashboardUsecase {
	return &dashboardUsecase{repo: repo}
}

func (u *dashboardUsecase) GetSummary(ctx context.Context, start, end time.Time) (*DashboardSummary, error) {
	cash, err := u.repo.GetCashBalance(ctx)
	if err != nil {
		return nil, err
	}

	revenue, err := u.repo.GetTotalRevenue(ctx, start, end)
	if err != nil {
		return nil, err
	}

	expense, err := u.repo.GetTotalExpense(ctx, start, end)
	if err != nil {
		return nil, err
	}

	count, err := u.repo.CountJournals(ctx, start, end)
	if err != nil {
		return nil, err
	}

	summary := &DashboardSummary{
		CashBalance:   cash,
		TotalRevenue:  revenue,
		TotalExpense:  expense,
		TotalJournals: count,
		NetProfit:     revenue - expense,
	}

	return summary, nil
}
