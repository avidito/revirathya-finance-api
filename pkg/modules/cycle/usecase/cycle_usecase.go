package usecase

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type cycleUsecase struct {
	cycleRepository domain.CycleRepository
}

func NewCycleUsecase(r domain.CycleRepository) domain.CycleUsecase {
	return &cycleUsecase{
		cycleRepository: r,
	}
}

// Usecase
func (u cycleUsecase) Create(cycle domain.Cycle) (domain.Cycle, error) {
	cycle.TotalDays = int64(cycle.EndDate.Time.Sub(cycle.StartDate.Time).Hours()) / 24
	createdCycle, err := u.cycleRepository.Create(cycle)
	return createdCycle, err
}

func (u cycleUsecase) Fetch(label string) ([]domain.CycleRead, error) {
	if label == "" {
		label = "%%"
	}

	cycleReadList, err := u.cycleRepository.Fetch(label)
	return cycleReadList, err
}

func (u cycleUsecase) GetByID(id int64) (domain.CycleRead, error) {
	cycleRead, err := u.cycleRepository.GetByID(id)
	return cycleRead, err
}

func (u cycleUsecase) Update(id int64, cycle domain.Cycle) (domain.Cycle, error) {
	cycle.TotalDays = int64(cycle.EndDate.Time.Sub(cycle.StartDate.Time).Hours()) / 24
	updatedCycle, err := u.cycleRepository.Update(id, cycle)
	return updatedCycle, err
}

func (u cycleUsecase) Delete(id int64) (domain.Cycle, error) {
	deletedCycle, err := u.cycleRepository.Delete(id)
	return deletedCycle, err
}
