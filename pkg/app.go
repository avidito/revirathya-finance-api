package pkg

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	_budget_handler "github.com/avidito/revirathya-finance-api/pkg/modules/budget/handler/http"
	_postgres_budget_repository "github.com/avidito/revirathya-finance-api/pkg/modules/budget/repository/postgres"
	_budget_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/budget/usecase"
	_expense_handler "github.com/avidito/revirathya-finance-api/pkg/modules/expense/handler/http"
	_postgres_expense_repository "github.com/avidito/revirathya-finance-api/pkg/modules/expense/repository/postgres"
	_expense_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/expense/usecase"
	_income_handler "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/handler/http"
	_postgres_income_repository "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/repository/postgres"
	_income_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/incomes/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	postgresBudgetRepository  domain.BudgetRepository
	postgresIncomeRepository  domain.IncomeRepository
	postgresExpenseRepository domain.ExpenseRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		postgresBudgetRepository:  _postgres_budget_repository.NewPostgresBudgetRepository(db),
		postgresIncomeRepository:  _postgres_income_repository.NewPostgresIncomesRepository(db),
		postgresExpenseRepository: _postgres_expense_repository.NewPostgresExpenseRepository(db),
	}
}

// Usecase
type Usecase struct {
	budgetUsecase  domain.BudgetUsecase
	incomeUsecase  domain.IncomeUsecase
	expenseUsecase domain.ExpenseUsecase
}

func NewUsecase(r *Repository) *Usecase {
	return &Usecase{
		budgetUsecase:  _budget_usecase.NewBudgetUsecase(r.postgresBudgetRepository),
		incomeUsecase:  _income_usecase.NewIncomeUsecease(r.postgresIncomeRepository),
		expenseUsecase: _expense_usecase.NewExpenseUsecase(r.postgresExpenseRepository),
	}
}

// Handler
type Handler struct {
	budgetHandler  _budget_handler.HttpBudgetHandler
	incomeHandler  _income_handler.HttpIncomeHandler
	expenseHandler _expense_handler.HttpExpenseHandler
}

func NewHandler(app *fiber.App, u *Usecase) {
	_budget_handler.NewHttpBudgetHandler(app, u.budgetUsecase)
	_income_handler.NewHttpIncomeHandler(app, u.incomeUsecase)
	_expense_handler.NewHttpExpenseHandler(app, u.expenseUsecase)
}
