package journal

import (
	"go_accounting/internal/account"
	"time"

	"github.com/google/uuid"
)

// request
type JournalDetailInput struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
	Debit     float64   `json:"debit" validate:"required"`
	Credit    float64   `json:"credit" validate:"required"`
}

type CreateJournalInput struct {
	Date        string               `json:"date" validate:"required"`
	Reference   string               `json:"reference" validate:"required"`
	Description string               `json:"description" validate:"required"`
	Details     []JournalDetailInput `json:"details" validate:"required"`
}

// response
type JournalDetailResponse struct {
	AccountID uuid.UUID                `json:"account_id"`
	Account   *account.AccountResponse `json:"account"`
	Debit     float64                  `json:"debit"`
	Credit    float64                  `json:"credit"`
}

type JournalResponse struct {
	ID          string                  `json:"id"`
	Date        time.Time               `json:"date"`
	Reference   string                  `json:"reference"`
	Description string                  `json:"description"`
	Details     []JournalDetailResponse `json:"details"`
}

// mapper
func ToJournalResponse(j *Journal, details []JournalDetailResponse) *JournalResponse {
	return &JournalResponse{
		ID:          j.ID.String(),
		Date:        j.Date,
		Reference:   j.Reference,
		Description: j.Description,
		Details:     details,
	}
}

func ToJournalDetailResponse(d JournalDetail, acc *account.AccountResponse) JournalDetailResponse {
	return JournalDetailResponse{
		AccountID: d.AccountID,
		Account:   acc,
		Debit:     d.Debit,
		Credit:    d.Credit,
	}
}
