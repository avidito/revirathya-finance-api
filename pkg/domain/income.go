package domain

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
)

type Income struct {
	ID           int64                `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Date         DateStandard         `json:"date"`
	IncomeTypeID int64                `json:"income_type_id"`
	LocationID   int64                `json:"location_id"`
	Description  string               `json:"description"`
	Amount       int64                `json:"amount"`
	IncomeType   models.RefIncomeType `json:"income_type" gorm:"foreignKey:IncomeTypeID;references:ID"`
	Location     models.RefLocation   `json:"location" gorm:"foreignKey:LocationID;references:ID"`
}

type IncomeRead struct {
	Date        DateStandard `json:"date"`
	IncomeType  string       `json:"income_type"`
	Location    string       `json:"location"`
	Description string       `json:"description"`
	Amount      int64        `json:"amount"`
}

type IncomeRepository interface {
	Create(income Income) (Income, error)
	Fetch(_date DateStandard, income_type string) ([]IncomeRead, error)
	Get(id int64) (IncomeRead, error)
	Update(id int64, income Income) (Income, error)
	Delete(id int64) (Income, error)
}

type IncomeUsecase interface {
	Create(income Income) (Income, error)
	Fetch(_date DateStandard, income_type string) ([]IncomeRead, error)
	Get(id int64) (IncomeRead, error)
	Update(id int64, income Income) (Income, error)
	Delete(id int64) (Income, error)
}
