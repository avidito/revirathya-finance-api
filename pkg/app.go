package pkg

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	_budget_handler "github.com/avidito/revirathya-finance-api/pkg/modules/budget/handler/http"
	_postgres_budget_repository "github.com/avidito/revirathya-finance-api/pkg/modules/budget/repository/postgres"
	_budget_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/budget/usecase"
	_income_handler "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/handler/http"
	_postgres_income_repository "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/repository/postgres"
	_income_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	postgresBudgetRepository domain.BudgetRepository
	postgresIncomeRepository domain.IncomeRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		postgresBudgetRepository: _postgres_budget_repository.NewPostgresBudgetRepository(db),
		postgresIncomeRepository: _postgres_income_repository.NewPostgresIncomesRepository(db),
	}
}

// Usecase
type Usecase struct {
	budgetUsecase domain.BudgetUsecase
	incomeUsecase domain.IncomeUsecase
}

func NewUsecase(r *Repository) *Usecase {
	return &Usecase{
		budgetUsecase: _budget_usecase.NewBudgetUsecase(r.postgresBudgetRepository),
		incomeUsecase: _income_usecase.NewIncomeUsecease(r.postgresIncomeRepository),
	}
}

// Handler
type Handler struct {
	budgetHandler _budget_handler.HttpBudgetHandler
	incomeHandler _income_handler.HttpIncomeHandler
}

func NewHandler(app *fiber.App, u *Usecase) {
	_budget_handler.NewHttpBudgetHandler(app, u.budgetUsecase)
	_income_handler.NewHttpIncomeHandler(app, u.incomeUsecase)
}
