package usecase

import (
	"time"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type expenseUsecase struct {
	expenseRepository domain.ExpenseRepository
}

func NewExpenseUsecase(r domain.ExpenseRepository) domain.ExpenseUsecase {
	return &expenseUsecase{
		expenseRepository: r,
	}
}

// Usecase
func (u expenseUsecase) Create(expense domain.Expense) (domain.Expense, error) {
	createdExpense, err := u.expenseRepository.Create(expense)
	return createdExpense, err
}

func (u expenseUsecase) Fetch(_date string, budget_type string) ([]domain.ExpenseRead, error) {
	if _date == "" {
		_date = time.Now().Format("2006-01-01")
	}

	if budget_type == "" {
		budget_type = "%%"
	}

	expenseReadList, err := u.expenseRepository.Fetch(_date, budget_type)
	return expenseReadList, err
}

func (u expenseUsecase) GetByID(id int64) (domain.ExpenseRead, error) {
	expenseRead, err := u.expenseRepository.GetByID(id)
	return expenseRead, err
}

func (u expenseUsecase) Update(id int64, expense domain.Expense) (domain.Expense, error) {
	updatedExpense, err := u.expenseRepository.Update(id, expense)
	return updatedExpense, err
}

func (u expenseUsecase) Delete(id int64) (domain.Expense, error) {
	deletedExpense, err := u.expenseRepository.Delete(id)
	return deletedExpense, err
}
