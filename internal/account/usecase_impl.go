package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type accountUsecase struct {
	repo AccountRepository
}

func NewAccountUsecase(repo AccountRepository) AccountUsecase {
	return &accountUsecase{repo: repo}
}

func (u *accountUsecase) GetAll(ctx context.Context, page, limit int) ([]AccountResponse, int64, error) {
	accounts, total, err := u.repo.FindAll(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}
	res := make([]AccountResponse, 0, len(accounts))
	for _, account := range accounts {
		res = append(res, *ToAccountResponse(&account))
	}
	return res, total, nil
}

func (u *accountUsecase) GetByID(ctx context.Context, id string) (*AccountResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	account, err := u.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return ToAccountResponse(account), nil
}

func (u *accountUsecase) Create(ctx context.Context, input CreateAccountInput) (*AccountResponse, error) {
	account := &Account{
		Code:           input.Code,
		Name:           input.Name,
		Type:           AccountType(input.Type),
		InitialBalance: input.InitialBalance,
	}
	err := u.repo.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return ToAccountResponse(account), nil
}

func (u *accountUsecase) Update(ctx context.Context, id string, input UpdateAccountInput) (*AccountResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	account, err := u.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	account.Code = input.Code
	account.Name = input.Name
	account.Type = AccountType(input.Type)
	account.InitialBalance = input.InitialBalance
	err = u.repo.UpdateAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	return ToAccountResponse(account), nil
}

func (u *accountUsecase) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}
	return u.repo.Delete(ctx, uid)
}
