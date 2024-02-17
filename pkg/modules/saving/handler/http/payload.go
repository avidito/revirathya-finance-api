package http

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

type SavingRequestBody struct {
	Date         domain.DateStandard `json:"date"`
	SavingTypeID int64               `json:"saving_type_id"`
	Flow         string              `json:"flow"`
	Amount       int64               `json:"amount"`
}

type SavingResponse struct {
	ID           int64               `json:"id"`
	Date         domain.DateStandard `json:"date"`
	SavingTypeID int64               `json:"saving_type_id"`
	Flow         string              `json:"flow"`
	Amount       int64               `json:"amount"`
}

type SavingReadResponse struct {
	Date       domain.DateStandard `json:"date"`
	SavingType string              `json:"saving_type"`
	Flow       string              `json:"flow"`
	Amount     int64               `json:"amount"`
}
