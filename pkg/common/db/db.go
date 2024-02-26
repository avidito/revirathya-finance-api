package db

import (
	"log"

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
		&domain.RefIncomeType{},
		&domain.RefLocation{},
		&domain.RefBudgetGroup{},
		&domain.RefBudgetType{},
		&domain.RefSavingType{},
		&domain.Cycle{},
		&domain.Income{},
		&domain.Budget{},
		&domain.Expense{},
		&domain.Transfer{},
		&domain.Saving{},
	)
}
