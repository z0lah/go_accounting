package ledger

import (
	"context"
	"go_accounting/internal/account"
	"time"

	"github.com/google/uuid"
)

type ledgerUsecase struct {
	ledgerRepo  LedgerRepository
	accountRepo account.AccountRepository
}

func NewLedgerUsecase(lr LedgerRepository, ar account.AccountRepository) LedgerUseCase {
	return &ledgerUsecase{
		ledgerRepo:  lr,
		accountRepo: ar,
	}
}

func (u *ledgerUsecase) GetAll(ctx context.Context, startDate, endDate time.Time) ([]LedgerResponse, error) {
	summaries, err := u.ledgerRepo.GetLedgerSummary(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var result []LedgerResponse

	for _, s := range summaries {
		acc, err := u.accountRepo.FindByID(ctx, uuid.MustParse(s.AccountID))
		if err != nil {
			continue // skip if account not found
		}

		balance := s.TotalDebit - s.TotalCredit
		if acc.Type.IsCreditNormal() {
			balance = s.TotalCredit - s.TotalDebit
		}

		result = append(result, LedgerResponse{
			AccountID:   s.AccountID,
			Code:        acc.Code,
			Name:        acc.Name,
			Type:        string(acc.Type),
			TotalDebit:  s.TotalDebit,
			TotalCredit: s.TotalCredit,
			Balance:     balance,
		})

	}

	return result, nil
}
