package user

import "github.com/google/uuid"

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type RegisterResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func ToRegisterResponse(u *User) *RegisterResponse {
	return &RegisterResponse{
		ID:     u.ID.String(),
		Name:   u.Name,
		Email:  u.Email,
		Phone:  u.Phone,
		Role:   u.Role,
		Status: u.Status,
	}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  *User
}
