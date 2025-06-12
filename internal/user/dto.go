package user

// single User lengkap
type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:     u.ID.String(),
		Name:   u.Name,
		Email:  u.Email,
		Role:   u.Role,
		Phone:  u.Phone,
		Status: u.Status,
	}
}

// Login
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type LoginResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func ToLoginResponse(u *User) *LoginResponse {
	return &LoginResponse{
		ID:   u.ID.String(),
		Name: u.Name,
		Role: u.Role,
	}
}

type AuthResponse struct {
	Token string         `json:"token"`
	User  *LoginResponse `json:"user"`
}

// Register
type RegisterInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=255"`
	Phone           string `json:"phone" validate:"required,min=10,max=13"`
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

// Update user
type UpdateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
