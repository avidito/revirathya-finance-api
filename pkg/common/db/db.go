package db

import (
	"log"

	"github.com/avidito/revirathya-finance-api/pkg/common/models"
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	doMigration(db)

	return db
}

func doMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Expense{},
		&models.RefIncomeType{},
		&models.RefLocation{},
		&models.RefBudgetGroup{},
		&models.RefBudgetType{},
		&models.RefSavingType{},
		&domain.Income{},
	)
}
