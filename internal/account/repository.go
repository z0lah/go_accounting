package account

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(ctx context.Context, a *Account) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *accountRepository) FindByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	var account Account
	if err := r.db.WithContext(ctx).First(&account, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindAll(ctx context.Context, page, limit int) ([]Account, int64, error) {
	var accounts []Account
	var total int64

	offset := (page - 1) * limit

	// hitung total
	tx := r.db.WithContext(ctx).Model(&Account{})
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//ambil data
	if err := tx.Limit(limit).Offset(offset).Find(&accounts).Error; err != nil {
		return nil, 0, err
	}

	return accounts, total, nil
}

func (r *accountRepository) UpdateAccount(ctx context.Context, a *Account) error {
	return r.db.WithContext(ctx).Save(a).Error
}

func (r *accountRepository) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Account{}, "id = ?", id).Error
}

func (r *accountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Account{}, "id = ?", id).Error
}
