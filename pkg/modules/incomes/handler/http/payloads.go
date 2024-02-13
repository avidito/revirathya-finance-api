package http

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

type CreateIncomeRequestBody struct {
	Date         domain.DateStandard `json:"date"`
	IncomeTypeID int64               `json:"income_type_id"`
	LocationID   int64               `json:"location_id"`
	Description  string              `json:"description"`
	Amount       int64               `json:"amount"`
}

type UpdateIncomeRequestBody struct {
	Date         domain.DateStandard `json:"date"`
	IncomeTypeID int64               `json:"income_type_id"`
	LocationID   int64               `json:"location_id"`
	Description  string              `json:"description"`
	Amount       int64               `json:"amount"`
}

type IncomeReponse struct {
	Date         domain.DateStandard `json:"date"`
	IncomeTypeID int64               `json:"income_type_id"`
	LocationID   int64               `json:"location_id"`
	Description  string              `json:"description"`
	Amount       int64               `json:"amount"`
}

type IncomeResponseSubs struct {
	Date        domain.DateStandard `json:"date"`
	IncomeType  string              `json:"income_type"`
	Location    string              `json:"location"`
	Description string              `json:"description"`
	Amount      int64               `json:"amount"`
}
