package http

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

type CycleRequestBody struct {
	Label     string              `json:"label"`
	StartDate domain.DateStandard `json:"start_date"`
	EndDate   domain.DateStandard `json:"end_date"`
	Notes     string              `json:"notes"`
}

type CycleResponse struct {
	ID        int64               `json:"id"`
	Label     string              `json:"label"`
	StartDate domain.DateStandard `json:"start_date"`
	EndDate   domain.DateStandard `json:"end_date"`
	Notes     string              `json:"notes"`
	TotalDays int64               `json:"total_days"`
}

type CycleReadResponse struct {
	ID        int64               `json:"id"`
	Label     string              `json:"label"`
	StartDate domain.DateStandard `json:"start_date"`
	EndDate   domain.DateStandard `json:"end_date"`
	Notes     string              `json:"notes"`
	TotalDays int64               `json:"total_days"`
}
