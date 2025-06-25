package user

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(ctx context.Context, u *User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}
func (m *mockUserRepository) FindAll(ctx context.Context, page, limit int) ([]User, int64, error) {
	return nil, 0, nil
}

func (m *mockUserRepository) FindNotActive(ctx context.Context) ([]User, error) {
	return nil, nil
}

func (m *mockUserRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return nil
}

func (m *mockUserRepository) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return nil
}

func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) != nil {
		return args.Get(0).(*User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return nil, nil
}

// mock repository
type mockTokenGenerator struct {
}

func (m *mockTokenGenerator) Generate(UserID string, email string, role string) (string, error) {
	return "dummy.jwt.token", nil
}

// test
// Register
func TestUserUsecase_Register(t *testing.T) {
	mockRepo := new(mockUserRepository)
	mockToken := &mockTokenGenerator{}

	userUsecase := NewUserUsecase(mockRepo, mockToken)

	input := RegisterInput{
		Name:     "zolah",
		Email:    "zolah@mail.com",
		Password: "password",
		Phone:    "08123456789",
	}
	mockRepo.
		On("Create", mock.Anything, mock.MatchedBy(func(u *User) bool {
			return u.Email == input.Email
		})).
		Return(nil)

	//act
	res, err := userUsecase.Register(context.Background(), input)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, input.Name, res.Name)
	assert.Equal(t, input.Email, res.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Register_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(mockUserRepository)
	mockToken := &mockTokenGenerator{}
	usecase := NewUserUsecase(mockRepo, mockToken)

	input := RegisterInput{
		Name:            "Zolah",
		Email:           "zolah@mail.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Phone:           "081234567890",
	}

	// Simulasikan error saat Create dipanggil
	mockRepo.
		On("Create", mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to save user"))

	// Act
	res, err := usecase.Register(context.Background(), input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.EqualError(t, err, "failed to save user")
	mockRepo.AssertExpectations(t)
}

// Login
func TestUserUsecase_Login_UserNotFound(t *testing.T) {
	mockRepo := new(mockUserRepository)
	mockToken := &mockTokenGenerator{}
	usecase := NewUserUsecase(mockRepo, mockToken)

	input := LoginInput{
		Email:    "notfound@mail.com",
		Password: "any-password",
	}

	mockRepo.On("FindByEmail", mock.Anything, input.Email).Return(nil, errors.New("user not found"))

	res, err := usecase.Login(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.EqualError(t, err, "user not found")
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Login_WrongPassword(t *testing.T) {
	mockRepo := new(mockUserRepository)
	mockToken := &mockTokenGenerator{}
	usecase := NewUserUsecase(mockRepo, mockToken)

	input := LoginInput{
		Email:    "zolah@mail.com",
		Password: "wrong-password",
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.DefaultCost)
	mockRepo.On("FindByEmail", mock.Anything, input.Email).Return(&User{
		ID:       uuid.New(),
		Email:    input.Email,
		Password: string(hashed),
		Role:     "staff",
	}, nil)

	res, err := usecase.Login(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.EqualError(t, err, "invalid password")
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Login_Success(t *testing.T) {
	mockRepo := new(mockUserRepository)
	mockToken := &mockTokenGenerator{}
	usecase := NewUserUsecase(mockRepo, mockToken)

	input := LoginInput{
		Email:    "zolah@mail.com",
		Password: "password123",
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	mockRepo.On("FindByEmail", mock.Anything, input.Email).Return(&User{
		ID:       uuid.New(),
		Email:    input.Email,
		Password: string(hashed),
		Role:     "staff",
		Status:   "active",
	}, nil)

	res, err := usecase.Login(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Token)
	mockRepo.AssertExpectations(t)
}
