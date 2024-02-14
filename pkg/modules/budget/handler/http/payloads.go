package http

type BudgetRequestBody struct {
	Cycle        string `json:"cycle"`
	BudgetTypeID int64  `json:"budget_type_id"`
	Amount       int64  `json:"amount"`
}

type BudgetResponse struct {
	ID           int64  `json:"ID"`
	Cycle        string `json:"cycle"`
	BudgetTypeID int64  `json:"budget_type_id"`
	Amount       int64  `json:"amount"`
}

type BudgetReadResponse struct {
	Cycle      string `json:"cycle"`
	BudgetType string `json:"budget_type"`
	Amount     int64  `json:"amount"`
}
