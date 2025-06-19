package report

import (
	"context"
	"go_accounting/internal/account"
	"time"

	"gorm.io/gorm"
)

type profitLossRepository struct {
	db *gorm.DB
}

func NewProfitLossRepository(db *gorm.DB) ProfitLossRepository {
	return &profitLossRepository{db: db}
}

func (r *profitLossRepository) GetRevenueAccounts(ctx context.Context, start, end time.Time) ([]AccountAmount, float64, error) {
	var results []AccountAmount
	var total float64

	tx := r.db.WithContext(ctx).
		Table("journal_details").
		Select("accounts.code, accounts.name, SUM(journal_details.credit) as amount").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Joins("JOIN journals ON journals.id = journal_details.journal_id").
		Where("accounts.type = ?", account.AccountTypeRevenue).
		Group("accounts.code, accounts.name")

	if !start.IsZero() {
		tx = tx.Where("journals.date >= ?", start)
	}
	if !end.IsZero() {
		tx = tx.Where("journals.date <= ?", end)
	}

	if err := tx.Debug().Scan(&results).Error; err != nil {
		return nil, 0, err
	}

	//hitung
	for _, r := range results {
		total += r.Amount
	}
	return results, total, nil
}

func (r *profitLossRepository) GetExpenseAccounts(ctx context.Context, start, end time.Time) ([]AccountAmount, float64, error) {
	var results []AccountAmount
	var total float64

	tx := r.db.WithContext(ctx).
		Table("journal_details").
		Select("accounts.code, accounts.name, SUM(journal_details.debit) as amount").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Joins("JOIN journals ON journals.id = journal_details.journal_id").
		Where("accounts.type = ?", account.AccountTypeExpense).
		Group("accounts.code, accounts.name")

	if !start.IsZero() {
		tx = tx.Where("journals.date >= ?", start)
	}
	if !end.IsZero() {
		tx = tx.Where("journals.date <= ?", end)
	}

	if err := tx.Scan(&results).Error; err != nil {
		return nil, 0, err
	}

	for _, r := range results {
		total += r.Amount
	}

	return results, total, nil
}
