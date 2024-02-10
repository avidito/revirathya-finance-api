package models

type RefIncomeType struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	IncomeType string `json:"income_type"`
}

type RefLocation struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Location string `json:"location"`
}

type RefBudgetGroup struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	BudgetGroup string `json:"budget_group"`
}

type RefBudgetType struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	BudgetGroupID int            `json:"id_budget_group"`
	BudgetGroup   RefBudgetGroup `json:"budget_group,omitempty" gorm:"foreignKey:BudgetGroupID;references:ID"`
	BudgetType    string         `json:"budget_type"`
}

type RefSavingType struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	SavingType string `json:"saving_type"`
}
