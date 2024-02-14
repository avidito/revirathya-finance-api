package usecase

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type budgetUsecase struct {
	budgetRepository domain.BudgetRepository
}

func NewBudgetUsecase(r domain.BudgetRepository) domain.BudgetUsecase {
	return &budgetUsecase{
		budgetRepository: r,
	}
}

// Usecase
func (u budgetUsecase) Create(budget domain.Budget) (domain.Budget, error) {
	createdBudget, err := u.budgetRepository.Create(budget)
	return createdBudget, err
}

func (u budgetUsecase) Fetch(cycle string, budget_type string) ([]domain.BudgetRead, error) {
	if cycle == "" {
		cycle = "%%"
	}

	if budget_type == "" {
		budget_type = "%%"
	}

	budgetReadList, err := u.budgetRepository.Fetch(cycle, budget_type)
	return budgetReadList, err
}

func (u budgetUsecase) GetByID(id int64) (domain.BudgetRead, error) {
	budgetRead, err := u.budgetRepository.GetByID(id)
	return budgetRead, err
}

func (u budgetUsecase) Update(id int64, budget domain.Budget) (domain.Budget, error) {
	updatedBudget, err := u.budgetRepository.Update(id, budget)
	return updatedBudget, err
}

func (u budgetUsecase) Delete(id int64) (domain.Budget, error) {
	deletedBudget, err := u.budgetRepository.Delete(id)
	return deletedBudget, err
}
