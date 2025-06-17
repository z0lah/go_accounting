package journal

import (
	"context"
	"errors"
	"go_accounting/internal/account"
	"time"

	"github.com/google/uuid"
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
	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, err
	}
	journal := &Journal{
		Date:        parsedDate,
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

func (u *journalUsecase) GetAll(ctx context.Context, page, limit int) ([]JournalResponse, int64, error) {
	journals, total, err := u.repo.FindAll(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var result []JournalResponse

	for _, j := range journals {
		var details []JournalDetailResponse
		for _, d := range j.Details {
			acc, err := u.accountRepo.FindByID(ctx, d.AccountID)
			if err != nil {
				return nil, 0, err
			}
			details = append(details, ToJournalDetailResponse(d, account.ToAccountResponse(acc)))
		}
		result = append(result, *ToJournalResponse(&j, details))
	}
	return result, total, nil
}

func (j *journalUsecase) GetByID(ctx context.Context, id string) (*JournalResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	journal, err := j.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	var details []JournalDetailResponse
	for _, d := range journal.Details {
		acc, err := j.accountRepo.FindByID(ctx, d.AccountID)
		if err != nil {
			return nil, err
		}
		details = append(details, ToJournalDetailResponse(d, account.ToAccountResponse(acc)))
	}

	return ToJournalResponse(journal, details), nil
}
