package seeder

import (
	"fmt"
	"go_accounting/internal/account"
	"go_accounting/internal/journal"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Seed accounts
	accounts := []account.Account{
		{ID: uuid.New(), Code: "1-1001", Name: "Kas", Type: account.AccountTypeAsset},
		{ID: uuid.New(), Code: "2-1001", Name: "Utang Dagang", Type: account.AccountTypeLiability},
		{ID: uuid.New(), Code: "3-1001", Name: "Modal", Type: account.AccountTypeEquity},
		{ID: uuid.New(), Code: "4-1001", Name: "Pendapatan Penjualan", Type: account.AccountTypeRevenue},
		{ID: uuid.New(), Code: "5-1001", Name: "Biaya Listrik", Type: account.AccountTypeExpense},
	}

	for _, a := range accounts {
		db.FirstOrCreate(&a, account.Account{Code: a.Code})
	}

	// Seed journal
	journalID := uuid.New()
	kasID := getAccountIDByCode("1-1001", accounts)
	pendapatanID := getAccountIDByCode("4-1001", accounts)

	j := journal.Journal{
		ID:          journalID,
		Date:        time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
		Reference:   "INV-001",
		Description: "Penjualan tunai",
		Details: []journal.JournalDetail{
			{ID: uuid.New(), JournalID: journalID, AccountID: kasID, Debit: 1500000},
			{ID: uuid.New(), JournalID: journalID, AccountID: pendapatanID, Credit: 1500000},
		},
	}

	return db.Create(&j).Error
}

func getAccountIDByCode(code string, accs []account.Account) uuid.UUID {
	for _, a := range accs {
		if a.Code == code {
			return a.ID
		}
	}
	return uuid.Nil
}

func SeedJournals(db *gorm.DB) error {
	// Hardcoded Account IDs from your provided data
	kasID := uuid.MustParse("77bd30cc-707b-4bb9-aa98-6147848f1e93")
	utangID := uuid.MustParse("c944ec01-d781-4c17-a10f-ebfe18172450")
	modalID := uuid.MustParse("762cfb4d-1e5e-4a52-87c1-79f7c7f51352")
	pendapatanID := uuid.MustParse("ba5ea603-9061-412b-bc4b-360e6649c34e")
	listrikID := uuid.MustParse("8e56bd25-ad1d-4318-8a25-9bd360f5d34c")

	// Generate 10 dummy journals
	for i := 1; i <= 10; i++ {
		jid := uuid.New()
		tanggal := time.Date(2025, 6, i, 0, 0, 0, 0, time.UTC)

		var details []journal.JournalDetail
		ref := "INV-202506" + twoDigit(i)
		desc := ""

		switch {
		case i <= 4:
			// Penjualan tunai
			details = []journal.JournalDetail{
				{ID: uuid.New(), JournalID: jid, AccountID: kasID, Debit: 1000000 + float64(i*10000)},
				{ID: uuid.New(), JournalID: jid, AccountID: pendapatanID, Credit: 1000000 + float64(i*10000)},
			}
			desc = fmt.Sprintf("Penjualan tunai ke-%d", i)
		case i <= 7:
			// Pembayaran utang
			details = []journal.JournalDetail{
				{ID: uuid.New(), JournalID: jid, AccountID: utangID, Debit: 500000 + float64(i*5000)},
				{ID: uuid.New(), JournalID: jid, AccountID: kasID, Credit: 500000 + float64(i*5000)},
			}
			desc = fmt.Sprintf("Pembayaran utang ke-%d", i)
		case i == 8:
			// Investasi modal
			details = []journal.JournalDetail{
				{ID: uuid.New(), JournalID: jid, AccountID: kasID, Debit: 2000000},
				{ID: uuid.New(), JournalID: jid, AccountID: modalID, Credit: 2000000},
			}
			desc = "Investasi modal tambahan"
		default:
			// Biaya listrik
			details = []journal.JournalDetail{
				{ID: uuid.New(), JournalID: jid, AccountID: listrikID, Debit: 300000},
				{ID: uuid.New(), JournalID: jid, AccountID: kasID, Credit: 300000},
			}
			desc = fmt.Sprintf("Pembayaran biaya listrik ke-%d", i)
		}

		j := journal.Journal{
			ID:          jid,
			Date:        tanggal,
			Reference:   ref,
			Description: desc,
			Details:     details,
		}

		if err := db.Create(&j).Error; err != nil {
			return err
		}
	}

	return nil
}

func twoDigit(n int) string {
	if n < 10 {
		return "0" + fmt.Sprint(n)
	}
	return fmt.Sprint(n)
}
