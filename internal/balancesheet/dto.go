package balancesheet

type BalanceItem struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type BalanceSheet struct {
	Assets         []BalanceItem `json:"assets"`
	Liability      []BalanceItem `json:"liability"`
	Equity         []BalanceItem `json:"equity"`
	TotalAsset     float64       `json:"total_asset"`
	TotalLiability float64       `json:"total_liability"`
	TotalEquity    float64       `json:"total_equity"`
}
