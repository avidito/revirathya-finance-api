package usecase

import (
	"time"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
)

// Define
type transferUsecase struct {
	transferRepository domain.TransferRepository
}

func NewTransferUsecase(r domain.TransferRepository) domain.TransferUsecase {
	return &transferUsecase{
		transferRepository: r,
	}
}

// Usecase
func (u transferUsecase) Create(transfer domain.Transfer) (domain.Transfer, error) {
	createdTransfer, err := u.transferRepository.Create(transfer)
	return createdTransfer, err
}

func (u transferUsecase) Fetch(date string, source string, destination string) ([]domain.TransferRead, error) {
	if date == "" {
		date = time.Now().Format("2006-01-01")
	}

	if source == "" {
		source = "%%"
	}

	if destination == "" {
		destination = "%%"
	}

	transferReadList, err := u.transferRepository.Fetch(date, source, destination)
	return transferReadList, err
}

func (u transferUsecase) GetByID(id int64) (domain.TransferRead, error) {
	transferRead, err := u.transferRepository.GetByID(id)
	return transferRead, err
}

func (u transferUsecase) Update(id int64, transfer domain.Transfer) (domain.Transfer, error) {
	updatedTransfer, err := u.transferRepository.Update(id, transfer)
	return updatedTransfer, err
}

func (u transferUsecase) Delete(id int64) (domain.Transfer, error) {
	deletedTransfer, err := u.transferRepository.Delete(id)
	return deletedTransfer, err
}
