package balancesheet

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type balanceSheet struct {
	db *gorm.DB
}

func NewBalanceSheetRepository(db *gorm.DB) BalanceSheetRepository {
	return &balanceSheet{db: db}
}

func (r *balanceSheet) GetAccountsByType(ctx context.Context, accountType string, upTo time.Time) ([]BalanceItem, float64, error) {
	var result []BalanceItem
	var total float64

	var sumExpr string
	if accountType == "asset" {
		sumExpr = "SUM(journal_details.debit) - SUM(journal_details.credit)"
	} else {
		sumExpr = "SUM(journal_details.credit) - SUM(journal_details.debit)"
	}
	tx := r.db.WithContext(ctx).
		Table("journal_details").
		Select("accounts.code, accounts.name, "+sumExpr+" as amount").
		Joins("JOIN accounts ON accounts.id = journal_details.account_id").
		Joins("JOIN journals ON journals.id = journal_details.journal_id").
		Where("accounts.type = ?", accountType).
		Where("journals.date <= ?", upTo).
		Group("accounts.code, accounts.name")

	if err := tx.Scan(&result).Error; err != nil {
		return nil, 0, err
	}
	// total
	for _, r := range result {
		total += r.Amount
	}

	return result, total, nil
}
