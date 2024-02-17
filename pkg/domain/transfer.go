package domain

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
)

type Transfer struct {
	ID            int64              `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Date          DateStandard       `json:"date"`
	SourceID      int64              `json:"source_id"`
	DestinationID int64              `json:"destination_id"`
	Description   string             `json:"description"`
	Amount        int64              `json:"amount"`
	Source        models.RefLocation `json:"source" gorm:"foreignKey:SourceID;references:ID"`
	Destination   models.RefLocation `json:"destination" gorm:"foreignKey:DestinationID;references:ID"`
}

type TransferRead struct {
	Date        DateStandard `json:"date"`
	Source      string       `json:"source"`
	Destination string       `json:"destination"`
	Description string       `json:"description"`
	Amount      int64        `json:"amount"`
}

type TransferRepository interface {
	Create(transfer Transfer) (Transfer, error)
	Fetch(date string, source string, destination string) ([]TransferRead, error)
	GetByID(id int64) (TransferRead, error)
	Update(id int64, transfer Transfer) (Transfer, error)
	Delete(id int64) (Transfer, error)
}

type TransferUsecase interface {
	Create(transfer Transfer) (Transfer, error)
	Fetch(date string, source string, destination string) ([]TransferRead, error)
	GetByID(id int64) (TransferRead, error)
	Update(id int64, transfer Transfer) (Transfer, error)
	Delete(id int64) (Transfer, error)
}
