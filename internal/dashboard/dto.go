package dashboard

type DashboardSummary struct {
	CashBalance   float64 `json:"cash_balance"`
	TotalRevenue  float64 `json:"total_revenue"`
	TotalExpense  float64 `json:"total_expense"`
	TotalJournals int64   `json:"total_journals"`
	NetProfit     float64 `json:"net_profit"`
}
