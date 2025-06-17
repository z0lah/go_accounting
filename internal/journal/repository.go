package journal

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type journalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) *journalRepository {
	return &journalRepository{db: db}
}

func (r *journalRepository) Create(ctx context.Context, journal *Journal) error {
	return r.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(journal).Error; err != nil {
				return err
			}
			for i := range journal.Details {
				journal.Details[i].JournalID = journal.ID
				if err := tx.Create(&journal.Details[i]).Error; err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (r *journalRepository) FindAll(ctx context.Context, page int, limit int) ([]Journal, int64, error) {
	var journals []Journal
	var total int64

	offset := (page - 1) * limit

	tx := r.db.WithContext(ctx).Model(&Journal{})

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := tx.Preload("Details").Order("date desc").
		Limit(limit).Offset(offset).
		Find(&journals).Error; err != nil {
		return nil, 0, err
	}

	return journals, total, nil
}

func (r *journalRepository) FindByID(ctx context.Context, id uuid.UUID) (*Journal, error) {
	var journal Journal
	if err := r.db.WithContext(ctx).
		Preload("Details").
		First(&journal, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &journal, nil
}
