package usecase

import (
	"time"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type savingUsecase struct {
	savingRepository domain.SavingRepository
}

func NewSavingUsecase(r domain.SavingRepository) domain.SavingUsecase {
	return &savingUsecase{
		savingRepository: r,
	}
}

// Usecase
func (u savingUsecase) Create(saving domain.Saving) (domain.Saving, error) {
	createdSaving, err := u.savingRepository.Create(saving)
	return createdSaving, err
}

func (u savingUsecase) Fetch(date string, saving_type string, flow string) ([]domain.SavingRead, error) {
	if date == "" {
		date = time.Now().Format("2006-01-01")
	}

	if saving_type == "" {
		saving_type = "%%"
	}

	if flow == "" {
		flow = "%%"
	}

	savingReadList, err := u.savingRepository.Fetch(date, saving_type, flow)
	return savingReadList, err
}

func (u savingUsecase) GetByID(id int64) (domain.SavingRead, error) {
	savingRead, err := u.savingRepository.GetByID(id)
	return savingRead, err
}

func (u savingUsecase) Update(id int64, saving domain.Saving) (domain.Saving, error) {
	updatedSaving, err := u.savingRepository.Update(id, saving)
	return updatedSaving, err
}

func (u savingUsecase) Delete(id int64) (domain.Saving, error) {
	deletedSaving, err := u.savingRepository.Delete(id)
	return deletedSaving, err
}
