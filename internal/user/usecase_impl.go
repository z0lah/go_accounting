package user

import (
	"context"
	"errors"
	"go_accounting/internal/shared/token"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo   UserRepository
	tokens token.TokenGenerator
}

func NewUserUsecase(repo UserRepository, tokens token.TokenGenerator) UserUsecase {
	return &userUsecase{
		repo:   repo,
		tokens: tokens,
	}
}

func (u *userUsecase) Register(ctx context.Context, input RegisterInput) (*RegisterResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Phone:    input.Phone,
	}

	err = u.repo.Create(ctx, user)
	if err != nil {
		log.Println("Gagal insert ", err)
		return nil, err
	}
	return ToRegisterResponse(user), nil
}
func (u *userUsecase) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	user, err := u.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// generate token
	tokenStr, err := u.tokens.Generate(user.ID.String(), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: tokenStr,
		User:  ToLoginResponse(user),
	}, nil
}

func (u *userUsecase) GetAll(ctx context.Context, page int, limit int) ([]UserResponse, int64, error) {
	rawUser, total, err := u.repo.FindAll(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}
	users := make([]UserResponse, 0, len(rawUser))
	for _, user := range rawUser {
		users = append(users, *ToUserResponse(&user))
	}

	return users, total, nil
}

func (u *userUsecase) GetNotActive(ctx context.Context) ([]UserResponse, error) {
	users, err := u.repo.FindNotActive(ctx)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []UserResponse{}, nil
	}
	res := make([]UserResponse, 0, len(users))
	for _, user := range users {
		res = append(res, *ToUserResponse(&user))
	}
	return res, nil
}

func (u *userUsecase) UpdateRole(ctx context.Context, id string, input UpdateRoleInput) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}
	return u.repo.UpdateRole(ctx, uid, string(input.Role))
}

func (u *userUsecase) UpdateStatus(ctx context.Context, id string, input UpdateStatusInput) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}
	return u.repo.UpdateStatus(ctx, uid, string(input.Status))
}
