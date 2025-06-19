package report

type AccountAmount struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type ProfitLossResponse struct {
	Revenue      []AccountAmount `json:"revenue"`
	Expense      []AccountAmount `json:"expense"`
	TotalRevenue float64         `json:"total_revenue"`
	TotalExpense float64         `json:"total_expense"`
	NetProfit    float64         `json:"net_profit"`
}
