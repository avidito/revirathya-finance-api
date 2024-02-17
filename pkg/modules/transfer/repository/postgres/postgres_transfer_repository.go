package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresTransferRepository struct {
	db *gorm.DB
}

func NewPostgresTransferRepository(db *gorm.DB) domain.TransferRepository {
	return &postgresTransferRepository{
		db: db,
	}
}

// Repository
func (r postgresTransferRepository) Create(transfer domain.Transfer) (domain.Transfer, error) {
	query := `
		INSERT INTO transfers (
			"date",
			source_id,
			destination_id,
			description,
			amount
		)
		VALUES (?, ?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		transfer.Date,
		transfer.SourceID,
		transfer.DestinationID,
		transfer.Description,
		transfer.Amount,
	).Row()

	var createdTransfer domain.Transfer
	err := row.Scan(
		&createdTransfer.ID,
		&createdTransfer.Date,
		&createdTransfer.SourceID,
		&createdTransfer.DestinationID,
		&createdTransfer.Description,
		&createdTransfer.Amount,
	)
	if err != nil {
		return domain.Transfer{}, err
	}

	return createdTransfer, nil
}

func (r postgresTransferRepository) Fetch(date string, source string, destination string) ([]domain.TransferRead, error) {
	query := `
		SELECT
		  t."date",
		  s_loc.location AS source,
		  d_loc.location AS destination,
		  t.description,
		  t.amount
		FROM transfers AS t
		JOIN ref_locations AS s_loc
		  ON t.source_id = s_loc.id
		JOIN ref_locations AS d_loc
		  ON t.destination_id = d_loc.id
		WHERE TRUE
		  AND DATE(t."date") = ?
		  AND s_loc.location LIKE ?
		  AND d_loc.location LIKE ?
	`

	rows, err := r.db.Raw(
		query,
		date,
		source,
		destination,
	).Rows()
	if err != nil {
		return nil, err
	}

	var transferReadList []domain.TransferRead
	var tmpTransferRead domain.TransferRead
	for rows.Next() {
		rows.Scan(
			&tmpTransferRead.Date,
			&tmpTransferRead.Source,
			&tmpTransferRead.Destination,
			&tmpTransferRead.Description,
			&tmpTransferRead.Amount,
		)

		transferReadList = append(transferReadList, tmpTransferRead)
	}
	return transferReadList, nil
}

func (r postgresTransferRepository) GetByID(id int64) (domain.TransferRead, error) {
	query := `
		SELECT
		  t."date",
		  s_loc.location AS source,
		  d_loc.location AS destination,
		  t.description,
		  t.amount
		FROM transfers AS t
		JOIN ref_locations AS s_loc
		  ON t.source_id = s_loc.id
		JOIN ref_locations AS d_loc
		  ON t.destination_id = d_loc.id
		WHERE TRUE
		  AND t.id = ?
	`

	var transferRead domain.TransferRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&transferRead.Date,
		&transferRead.Source,
		&transferRead.Destination,
		&transferRead.Description,
		&transferRead.Amount,
	)
	if err != nil {
		return domain.TransferRead{}, err
	}
	return transferRead, nil
}

func (r postgresTransferRepository) Update(id int64, transfer domain.Transfer) (domain.Transfer, error) {
	query := `
		UPDATE transfers
		SET
		  "date" = ?,
		  source_id = ?,
		  destination_id = ?,
		  description = ?,
		  amount = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		transfer.Date,
		transfer.SourceID,
		transfer.DestinationID,
		transfer.Description,
		transfer.Amount,
		id,
	).Row()

	var updatedTransfer domain.Transfer
	err := row.Scan(
		&updatedTransfer.ID,
		&updatedTransfer.Date,
		&updatedTransfer.SourceID,
		&updatedTransfer.DestinationID,
		&updatedTransfer.Description,
		&updatedTransfer.Amount,
	)
	if err != nil {
		return domain.Transfer{}, err
	}
	return updatedTransfer, nil
}

func (r postgresTransferRepository) Delete(id int64) (domain.Transfer, error) {
	query := `
		DELETE FROM transfers
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedTransfer domain.Transfer
	err := row.Scan(
		&deletedTransfer.ID,
		&deletedTransfer.Date,
		&deletedTransfer.SourceID,
		&deletedTransfer.DestinationID,
		&deletedTransfer.Description,
		&deletedTransfer.Amount,
	)
	if err != nil {
		return domain.Transfer{}, err
	}
	return deletedTransfer, nil
}
