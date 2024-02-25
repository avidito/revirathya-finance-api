package http

type RefIncomeTypeResponse struct {
	ID         int64  `json:"id"`
	IncomeType string `json:"income_type"`
}

type RefLocationResponse struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
}

type RefBudgetGroupResponse struct {
	ID          int64  `json:"id"`
	BudgetGroup string `json:"budget_group"`
}

type RefBudgetTypeReadResponse struct {
	ID          int64  `json:"id"`
	BudgetGroup string `json:"budget_group"`
	BudgetType  string `json:"budget_type"`
}

type RefSavingTypeResponse struct {
	ID         int64  `json:"id"`
	SavingType string `json:"saving_type"`
}
