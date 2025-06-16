package account

type CreateAccountInput struct {
	Code           string      `json:"code" validate:"required"`
	Name           string      `json:"name" validate:"required"`
	Type           AccountType `json:"type" validate:"required,oneof=asset liability equity revenue expense"`
	InitialBalance float64     `json:"initial_balance"`
}

type UpdateAccountInput struct {
	ID             string      `json:"id"`
	Code           string      `json:"code" validate:"required"`
	Name           string      `json:"name" validate:"required"`
	Type           AccountType `json:"type" validate:"required,oneof=asset liability equity revenue expense"`
	InitialBalance float64     `json:"initial_balance"`
}

type AccountResponse struct {
	ID             string  `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	InitialBalance float64 `json:"initial_balance"`
}

func ToAccountResponse(a *Account) *AccountResponse {
	return &AccountResponse{
		ID:             a.ID.String(),
		Code:           a.Code,
		Name:           a.Name,
		Type:           string(a.Type),
		InitialBalance: a.InitialBalance,
	}
}
