package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}
func (r *userRepository) FindAll(ctx context.Context, page int, limit int) ([]User, int64, error) {
	var users []User
	var total int64
	offset := (page - 1) * limit

	// hitung total
	tx := r.db.WithContext(ctx).Model(&User{})
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//ambil data
	if err := tx.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
