package dashboard

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetCashBalance(ctx context.Context) (float64, error) {
	var total float64

	err := r.db.WithContext(ctx).
		Table("journal_details").
		Select("SUM(journal_details.debit) - SUM(journal_details.credit)").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Where("accounts.type = ? AND accounts.code LIKE ?", "asset", "1-%").
		Scan(&total).Error

	return total, err
}

func (r *dashboardRepository) GetTotalExpense(ctx context.Context, start, end time.Time) (float64, error) {
	var total float64
	err := r.db.WithContext(ctx).
		Table("journal_details").
		Select("SUM(journal_details.debit)").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Joins("JOIN journals ON journals.id = journal_details.journal_id").
		Where("accounts.type = ?", "expense").
		Where("journals.date BETWEEN ? AND ?", start, end).
		Scan(&total).Error

	return total, err
}

func (r *dashboardRepository) GetTotalRevenue(ctx context.Context, start, end time.Time) (float64, error) {
	var total float64
	err := r.db.WithContext(ctx).
		Table("journal_details").
		Select("SUM(journal_details.credit)").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Joins("JOIN journals ON journals.id = journal_details.journal_id").
		Where("accounts.type = ?", "revenue").
		Where("journals.date BETWEEN ? AND ?", start, end).
		Scan(&total).Error

	return total, err
}

func (r *dashboardRepository) CountJournals(ctx context.Context, start, end time.Time) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Table("journals").
		Where("date BETWEEN ? AND ?", start, end).
		Count(&count).Error
	return count, err
}
