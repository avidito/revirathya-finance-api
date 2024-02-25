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
	_reference_handler "github.com/avidito/revirathya-finance-api/pkg/modules/reference/handler/http"
	_postgres_reference_repository "github.com/avidito/revirathya-finance-api/pkg/modules/reference/repository/postgres"
	_reference_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/reference/usecase"
	_saving_handler "github.com/avidito/revirathya-finance-api/pkg/modules/saving/handler/http"
	_postgres_saving_repository "github.com/avidito/revirathya-finance-api/pkg/modules/saving/repository/postgres"
	_saving_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/saving/usecase"
	_transfer_handler "github.com/avidito/revirathya-finance-api/pkg/modules/transfer/handler/http"
	_postgres_transfer_repository "github.com/avidito/revirathya-finance-api/pkg/modules/transfer/repository/postgres"
	_transfer_usecase "github.com/avidito/revirathya-finance-api/pkg/modules/transfer/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	postgresBudgetRepository    domain.BudgetRepository
	postgresIncomeRepository    domain.IncomeRepository
	postgresExpenseRepository   domain.ExpenseRepository
	postgresTransferRepository  domain.TransferRepository
	postgresSavingRepository    domain.SavingRepository
	postgresReferenceRepository domain.ReferenceRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		postgresBudgetRepository:    _postgres_budget_repository.NewPostgresBudgetRepository(db),
		postgresIncomeRepository:    _postgres_income_repository.NewPostgresIncomesRepository(db),
		postgresExpenseRepository:   _postgres_expense_repository.NewPostgresExpenseRepository(db),
		postgresTransferRepository:  _postgres_transfer_repository.NewPostgresTransferRepository(db),
		postgresSavingRepository:    _postgres_saving_repository.NewPostgresSavingRepository(db),
		postgresReferenceRepository: _postgres_reference_repository.NewPostgresReferenceRepository(db),
	}
}

// Usecase
type Usecase struct {
	budgetUsecase    domain.BudgetUsecase
	incomeUsecase    domain.IncomeUsecase
	expenseUsecase   domain.ExpenseUsecase
	transferUsecase  domain.TransferUsecase
	savingUsecase    domain.SavingUsecase
	referenceUsecase domain.ReferenceUsecase
}

func NewUsecase(r *Repository) *Usecase {
	return &Usecase{
		budgetUsecase:    _budget_usecase.NewBudgetUsecase(r.postgresBudgetRepository),
		incomeUsecase:    _income_usecase.NewIncomeUsecease(r.postgresIncomeRepository),
		expenseUsecase:   _expense_usecase.NewExpenseUsecase(r.postgresExpenseRepository),
		transferUsecase:  _transfer_usecase.NewTransferUsecase(r.postgresTransferRepository),
		savingUsecase:    _saving_usecase.NewSavingUsecase(r.postgresSavingRepository),
		referenceUsecase: _reference_usecase.NewReferenceUsecase(r.postgresReferenceRepository),
	}
}

// Handler
type Handler struct {
	budgetHandler   _budget_handler.HttpBudgetHandler
	incomeHandler   _income_handler.HttpIncomeHandler
	expenseHandler  _expense_handler.HttpExpenseHandler
	transferHandler _transfer_handler.HttpTransferHandler
	savingHandler   _saving_handler.HttpSavingHandler
}

func NewHandler(app *fiber.App, u *Usecase) {
	_budget_handler.NewHttpBudgetHandler(app, u.budgetUsecase)
	_income_handler.NewHttpIncomeHandler(app, u.incomeUsecase)
	_expense_handler.NewHttpExpenseHandler(app, u.expenseUsecase)
	_transfer_handler.NewHttpTransferHandler(app, u.transferUsecase)
	_saving_handler.NewHttpSavingHandler(app, u.savingUsecase)
	_reference_handler.NewHttpReferenceHandler(app, u.referenceUsecase)
}
