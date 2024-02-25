package domain

type Budget struct {
	ID           int64         `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Cycle        string        `json:"cycle"`
	BudgetTypeID int64         `json:"budget_type_id"`
	Amount       int64         `json:"amount"`
	BudgetType   RefBudgetType `json:"budget_type" gorm:"foreignKey:BudgetTypeID;references:ID"`
}

type BudgetRead struct {
	Cycle      string `json:"cycle"`
	BudgetType string `json:"budget_type"`
	Amount     int64  `json:"amount"`
}

type BudgetRepository interface {
	Create(budget Budget) (Budget, error)
	Fetch(cycle string, budget_type string) ([]BudgetRead, error)
	GetByID(id int64) (BudgetRead, error)
	Update(id int64, budget Budget) (Budget, error)
	Delete(id int64) (Budget, error)
}

type BudgetUsecase interface {
	Create(budget Budget) (Budget, error)
	Fetch(cycle string, budget_type string) ([]BudgetRead, error)
	GetByID(id int64) (BudgetRead, error)
	Update(id int64, budget Budget) (Budget, error)
	Delete(id int64) (Budget, error)
}
