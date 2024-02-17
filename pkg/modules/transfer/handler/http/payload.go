package http

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

type TransferRequestBody struct {
	Date          domain.DateStandard `json:"date"`
	SourceID      int64               `json:"source_id"`
	DestinationID int64               `json:"destination_id"`
	Description   string              `json:"description"`
	Amount        int64               `json:"amount"`
}

type TransferResponse struct {
	ID            int64               `json:"id"`
	Date          domain.DateStandard `json:"date"`
	SourceID      int64               `json:"source_id"`
	DestinationID int64               `json:"destination_id"`
	Description   string              `json:"description"`
	Amount        int64               `json:"amount"`
}

type TransferReadResponse struct {
	Date        domain.DateStandard `json:"date"`
	Source      string              `json:"source"`
	Destination string              `json:"destination"`
	Description string              `json:"description"`
	Amount      int64               `json:"amount"`
}
