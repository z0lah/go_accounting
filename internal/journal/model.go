package journal

import (
	"go_accounting/internal/account"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Journal struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key;"`
	Date        time.Time       `gorm:"not null"`
	Reference   string          `gorm:"not null"`
	Description string          `gorm:"not null type:text"`
	Details     []JournalDetail `gorm:"foreignKey:JournalID"`
	gorm.Model
}

type JournalDetail struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key;"`
	JournalID uuid.UUID       `gorm:"not null type:uuid index"`
	AccountID uuid.UUID       `gorm:"not null type:uuid index"`
	Account   account.Account `gorm:"foreignKey:AccountID"`
	Debit     float64         `gorm:"Default:0"`
	Credit    float64         `gorm:"Default:0"`
	gorm.Model
}

func (j *Journal) BeforeCreate(tx *gorm.DB) (err error) {
	j.ID = uuid.New()
	return
}

func (j *JournalDetail) BeforeCreate(tx *gorm.DB) (err error) {
	j.ID = uuid.New()
	return
}
