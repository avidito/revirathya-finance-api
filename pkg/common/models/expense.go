package models

type Expense struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	BudgetType  string `json:"budget_type"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
