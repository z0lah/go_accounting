package journal

import (
	"context"
	"errors"
	"go_accounting/internal/account"
)

type journalUsecase struct {
	repo        JournalRepository
	accountRepo account.AccountRepository
}

func NewJournalUsecase(repo JournalRepository, accountRepo account.AccountRepository) JournalUsecase {
	return &journalUsecase{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (u *journalUsecase) Create(ctx context.Context, input CreateJournalInput) (*JournalResponse, error) {
	var totalDebit, totalCredit float64

	// validasi total debit dan credit harus seimbang
	for _, d := range input.Details {
		totalDebit += d.Debit
		totalCredit += d.Credit

		if d.Debit > 0 && d.Credit > 0 {
			return nil, errors.New("each detail must have either debit or credit, not both")
		}
		if d.Debit == 0 && d.Credit == 0 {
			return nil, errors.New("each detail must have debit or credit")
		}
	}
	if totalDebit != totalCredit {
		return nil, errors.New("total debit and credit must be equal")
	}

	// maping journal dan detail
	journal := &Journal{
		Date:        input.Date,
		Reference:   input.Reference,
		Description: input.Description,
	}
	for _, d := range input.Details {
		journal.Details = append(journal.Details, JournalDetail{
			AccountID: d.AccountID,
			Debit:     d.Debit,
			Credit:    d.Credit,
		})
	}

	//save
	if err := u.repo.Create(ctx, journal); err != nil {
		return nil, err
	}

	// inject account for each detail
	var detailResponse []JournalDetailResponse
	for _, d := range journal.Details {
		acc, err := u.accountRepo.FindByID(ctx, d.AccountID)
		if err != nil {
			return nil, err
		}
		detailResponse = append(detailResponse, ToJournalDetailResponse(d, account.ToAccountResponse(acc)))
	}
	return ToJournalResponse(journal, detailResponse), nil
}

func (j *journalUsecase) GetAll(ctx context.Context, page, limit int) ([]JournalResponse, int64, error) {
	return nil, 0, nil
}

func (j *journalUsecase) GetByID(ctx context.Context, id string) (*JournalResponse, error) {
	return nil, nil
}
