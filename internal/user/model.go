package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Phone    string    `gorm:"not null"`
	Role     string    `gorm:"type:varchar(10);not null;default:staff"`
	Status   string    `gorm:"type:varchar(10);not null;default:not_active"`
	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

// enum role
type Role string

const (
	RoleAdmin Role = "admin"
	RoleStaff Role = "staff"
)

func IsValidRole(r string) bool {
	return r == string(RoleAdmin) || r == string(RoleStaff)
}

// enum status
type Status string

const (
	StatusActive    Status = "active"
	StatusNotActive Status = "not_active"
)

func IsValidStatus(s string) bool {
	return s == string(StatusActive) || s == string(StatusNotActive)
}
