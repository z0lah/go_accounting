package user

import (
	"context"
	"errors"
	"go_accounting/internal/shared/token"
	"log"

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

func (u *userUsecase) GetAll(ctx context.Context) ([]User, error) {
	return u.repo.FindAll(ctx)
}
