package http

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

type ExpenseRequestBody struct {
	Date         domain.DateStandard `json:"date"`
	BudgetTypeID int64               `json:"budget_type_id"`
	LocationID   int64               `json:"location_id"`
	Description  string              `json:"description"`
	Amount       int64               `json:"amount"`
}

type ExpenseResponse struct {
	ID           int64               `json:"id"`
	Date         domain.DateStandard `json:"date"`
	BudgetTypeID int64               `json:"budget_type_id"`
	LocationID   int64               `json:"location_id"`
	Description  string              `json:"description"`
	Amount       int64               `json:"amount"`
}

type ExpenseReadResponse struct {
	Date        domain.DateStandard `json:"date"`
	BudgetType  string              `json:"budget_type"`
	Location    string              `json:"location"`
	Description string              `json:"description"`
	Amount      int64               `json:"amount"`
}
