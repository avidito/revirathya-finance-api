package usecase

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type referenceUsecase struct {
	referenceRepository domain.ReferenceRepository
}

func NewReferenceUsecase(r domain.ReferenceRepository) domain.ReferenceUsecase {
	return &referenceUsecase{
		referenceRepository: r,
	}
}

// Usecase
func (u referenceUsecase) GetIncomeTypeByID(id int64) (domain.RefIncomeType, error) {
	refIncomeType, err := u.referenceRepository.GetIncomeTypeByID(id)
	return refIncomeType, err
}

func (u referenceUsecase) GetLocationByID(id int64) (domain.RefLocation, error) {
	refLocation, err := u.referenceRepository.GetLocationByID(id)
	return refLocation, err
}

func (u referenceUsecase) GetBudgetGroupByID(id int64) (domain.RefBudgetGroup, error) {
	refBudgetGroupRead, err := u.referenceRepository.GetBudgetGroupByID(id)
	return refBudgetGroupRead, err
}

func (u referenceUsecase) GetBudgetTypeByID(id int64) (domain.RefBudgetTypeRead, error) {
	refBudgetType, err := u.referenceRepository.GetBudgetTypeByID(id)
	return refBudgetType, err
}

func (u referenceUsecase) GetSavingTypeByID(id int64) (domain.RefSavingType, error) {
	refSavingType, err := u.referenceRepository.GetSavingTypeByID(id)
	return refSavingType, err
}

func (u referenceUsecase) FetchIncomeTypes(name string) ([]domain.RefIncomeType, error) {
	if name == "" {
		name = "%%"
	}

	refIncomeType, err := u.referenceRepository.FetchIncomeTypes(name)
	return refIncomeType, err
}

func (u referenceUsecase) FetchLocations(name string) ([]domain.RefLocation, error) {
	if name == "" {
		name = "%%"
	}

	refLocation, err := u.referenceRepository.FetchLocations(name)
	return refLocation, err
}

func (u referenceUsecase) FetchBudgetGroups(name string) ([]domain.RefBudgetGroup, error) {
	if name == "" {
		name = "%%"
	}

	refBudgetGroup, err := u.referenceRepository.FetchBudgetGroups(name)
	return refBudgetGroup, err
}

func (u referenceUsecase) FetchBudgetTypes(name string, budget_group string) ([]domain.RefBudgetTypeRead, error) {
	if name == "" {
		name = "%%"
	}

	if budget_group == "" {
		budget_group = "%%"
	}

	refBudgetTypeRead, err := u.referenceRepository.FetchBudgetTypes(name, budget_group)
	return refBudgetTypeRead, err
}

func (u referenceUsecase) FetchSavingTypes(name string) ([]domain.RefSavingType, error) {
	if name == "" {
		name = "%%"
	}

	refSavingType, err := u.referenceRepository.FetchSavingTypes(name)
	return refSavingType, err
}
