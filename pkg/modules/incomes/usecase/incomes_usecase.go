package usecase

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type incomeUsecase struct {
	incomeRepository domain.IncomeRepository
}

func NewIncomeUsecease(r domain.IncomeRepository) domain.IncomeUsecase {
	return &incomeUsecase{
		incomeRepository: r,
	}
}

// Usecase
func (u incomeUsecase) Create(income domain.Income) (domain.Income, error) {
	income, err := u.incomeRepository.Create(income)
	return income, err
}

func (u incomeUsecase) Fetch(_date domain.DateStandard, income_type string) ([]domain.IncomeRead, error) {
	incomes, err := u.incomeRepository.Fetch(_date, income_type)
	return incomes, err
}

func (u incomeUsecase) Get(id int64) (domain.IncomeRead, error) {
	income, err := u.incomeRepository.Get(id)
	return income, err
}

func (u incomeUsecase) Update(id int64, income domain.Income) (domain.Income, error) {
	income, err := u.incomeRepository.Update(id, income)
	return income, err
}

func (u incomeUsecase) Delete(id int64) (domain.Income, error) {
	income, err := u.incomeRepository.Delete(id)
	return income, err
}
