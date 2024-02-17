package domain

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
)

type Saving struct {
	ID           int64                `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Date         DateStandard         `json:"date"`
	SavingTypeID int64                `json:"saving_type_id"`
	Flow         string               `json:"flow"`
	Amount       int64                `json:"amount"`
	SavingType   models.RefSavingType `json:"saving_type_id" gorm:"foreignKey:SavingTypeID;references:ID"`
}

type SavingRead struct {
	Date       DateStandard `json:"date"`
	SavingType string       `json:"saving"`
	Flow       string       `json:"flow"`
	Amount     int64        `json:"amount"`
}

type SavingRepository interface {
	Create(saving Saving) (Saving, error)
	Fetch(date string, saving_type string, flow string) ([]SavingRead, error)
	GetByID(id int64) (SavingRead, error)
	Update(id int64, saving Saving) (Saving, error)
	Delete(id int64) (Saving, error)
}

type SavingUsecase interface {
	Create(saving Saving) (Saving, error)
	Fetch(date string, saving_type string, flow string) ([]SavingRead, error)
	GetByID(id int64) (SavingRead, error)
	Update(id int64, saving Saving) (Saving, error)
	Delete(id int64) (Saving, error)
}
