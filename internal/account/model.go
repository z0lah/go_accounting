package account

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID             uuid.UUID   `gorm:"type:uuid;primary_key;"`
	Code           string      `gorm:"unique;not null"`
	Name           string      `gorm:"not null"`
	Type           AccountType `gorm:"type:varchar(10);not null;default:asset"`
	InitialBalance float64     `gorm:"default:0"`
	gorm.Model
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

type AccountType string

const (
	Asset     AccountType = "asset"
	Liability AccountType = "liability"
	Equity    AccountType = "equity"
	Revenue   AccountType = "revenue"
	Expense   AccountType = "expense"
)
