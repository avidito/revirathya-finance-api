package domain

type RefIncomeType struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	IncomeType string `json:"income_type"`
}

type RefLocation struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Location string `json:"location"`
}

type RefBudgetGroup struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	BudgetGroup string `json:"budget_group"`
}

type RefBudgetType struct {
	ID            int64          `json:"id" gorm:"primaryKey"`
	BudgetGroupID int64          `json:"budget_group_id"`
	BudgetType    string         `json:"budget_type"`
	BudgetGroup   RefBudgetGroup `json:"budget_group" gorm:"foreignKey:BudgetGroupID;references:ID"`
}

type RefBudgetTypeRead struct {
	ID          int64  `json:"id"`
	BudgetGroup string `json:"budget_group"`
	BudgetType  string `json:"budget_type"`
}

type RefSavingType struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	SavingType string `json:"saving_type"`
}

type ReferenceRepository interface {
	FetchIncomeTypes(name string) ([]RefIncomeType, error)
	FetchLocations(name string) ([]RefLocation, error)
	FetchBudgetGroups(name string) ([]RefBudgetGroup, error)
	FetchBudgetTypes(name string, budget_group string) ([]RefBudgetTypeRead, error)
	FetchSavingTypes(name string) ([]RefSavingType, error)
	GetIncomeTypeByID(id int64) (RefIncomeType, error)
	GetLocationByID(id int64) (RefLocation, error)
	GetBudgetGroupByID(id int64) (RefBudgetGroup, error)
	GetBudgetTypeByID(id int64) (RefBudgetTypeRead, error)
	GetSavingTypeByID(id int64) (RefSavingType, error)
}

type ReferenceUsecase interface {
	FetchIncomeTypes(name string) ([]RefIncomeType, error)
	FetchLocations(name string) ([]RefLocation, error)
	FetchBudgetGroups(name string) ([]RefBudgetGroup, error)
	FetchBudgetTypes(name string, budget_group string) ([]RefBudgetTypeRead, error)
	FetchSavingTypes(name string) ([]RefSavingType, error)
	GetIncomeTypeByID(id int64) (RefIncomeType, error)
	GetLocationByID(id int64) (RefLocation, error)
	GetBudgetGroupByID(id int64) (RefBudgetGroup, error)
	GetBudgetTypeByID(id int64) (RefBudgetTypeRead, error)
	GetSavingTypeByID(id int64) (RefSavingType, error)
}
