package pkg

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	_income_handler "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/handler/http"
	_postgres_income_repository "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/repository/postgres"
	_income_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	postgresIncomeRepository domain.IncomeRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		postgresIncomeRepository: _postgres_income_repository.NewPostgresIncomesRepository(db),
	}
}

// Usecase
type Usecase struct {
	incomeUsecase domain.IncomeUsecase
}

func NewUsecase(r *Repository) *Usecase {
	return &Usecase{
		incomeUsecase: _income_usecase.NewIncomeUsecease(r.postgresIncomeRepository),
	}
}

// Handler
type Handler struct {
	incomeHandler _income_handler.HttpIncomeHandler
}

func NewHandler(app *fiber.App, u *Usecase) {
	_income_handler.NewHttpIncomeHandler(app, u.incomeUsecase)
}
