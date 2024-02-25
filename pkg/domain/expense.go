package domain

type Expense struct {
	ID           int64         `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Date         DateStandard  `json:"date"`
	BudgetTypeID int64         `json:"budget_type_id"`
	LocationID   int64         `json:"location_id"`
	Description  string        `json:"description"`
	Amount       int64         `json:"amount"`
	BudgetType   RefBudgetType `json:"budget_type" gorm:"foreignKey:BudgetTypeID;references:ID"`
	Location     RefLocation   `json:"location" gorm:"foreignKey:LocationID;references:ID"`
}

type ExpenseRead struct {
	Date        DateStandard `json:"date"`
	BudgetType  string       `json:"budget_type"`
	Location    string       `json:"location"`
	Description string       `json:"description"`
	Amount      int64        `json:"amount"`
}

type ExpenseRepository interface {
	Create(expense Expense) (Expense, error)
	Fetch(_date string, budget_type string) ([]ExpenseRead, error)
	GetByID(id int64) (ExpenseRead, error)
	Update(id int64, expense Expense) (Expense, error)
	Delete(id int64) (Expense, error)
}

type ExpenseUsecase interface {
	Create(expense Expense) (Expense, error)
	Fetch(_date string, budget_type string) ([]ExpenseRead, error)
	GetByID(id int64) (ExpenseRead, error)
	Update(id int64, expense Expense) (Expense, error)
	Delete(id int64) (Expense, error)
}
