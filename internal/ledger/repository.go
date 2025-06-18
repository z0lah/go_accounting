package ledger

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type LedgerRepository interface {
	GetLedgerSummary(ctx context.Context, startDate, endDate time.Time) ([]AccountSummary, error)
}

type AccountSummary struct {
	AccountID   string
	TotalDebit  float64
	TotalCredit float64
}

type ledgerRepository struct {
	db *gorm.DB
}

func NewLedgerRepository(db *gorm.DB) LedgerRepository {
	return &ledgerRepository{db}
}

func (r *ledgerRepository) GetLedgerSummary(ctx context.Context, startDate, endDate time.Time) ([]AccountSummary, error) {
	var result []AccountSummary

	tx := r.db.WithContext(ctx).
		Table("journal_details").
		Joins("JOIN journals ON journal_details.journal_id = journals.id")

	if !startDate.IsZero() {
		tx = tx.Where("journals.date >= ?", startDate)
	}
	if !endDate.IsZero() {
		tx = tx.Where("journals.date <= ?", endDate)
	}

	err := tx.Select("account_id, sum(debit) as total_debit, sum(credit) as total_credit").
		Group("account_id").
		Scan(&result).Error

	return result, err
}
