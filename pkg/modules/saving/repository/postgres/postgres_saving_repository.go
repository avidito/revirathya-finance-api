package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresSavingRepository struct {
	db *gorm.DB
}

func NewPostgresSavingRepository(db *gorm.DB) domain.SavingRepository {
	return &postgresSavingRepository{
		db: db,
	}
}

// Repository
func (r postgresSavingRepository) Create(saving domain.Saving) (domain.Saving, error) {
	query := `
		INSERT INTO savings (
			"date",
			saving_type_id,
			flow,
			amount
		)
		VALUES (?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		saving.Date,
		saving.SavingTypeID,
		saving.Flow,
		saving.Amount,
	).Row()

	var createdSaving domain.Saving
	err := row.Scan(
		&createdSaving.ID,
		&createdSaving.Date,
		&createdSaving.SavingTypeID,
		&createdSaving.Flow,
		&createdSaving.Amount,
	)
	if err != nil {
		return domain.Saving{}, err
	}

	return createdSaving, nil
}

func (r postgresSavingRepository) Fetch(date string, saving_type string, flow string) ([]domain.SavingRead, error) {
	query := `
		SELECT
		  sv."date",
		  sv_t.saving_type,
		  sv.flow,
		  sv.amount
		FROM savings AS sv
		JOIN ref_saving_types AS sv_t
		  ON sv.saving_type_id = sv_t.id
		WHERE TRUE
		  AND DATE(sv."date") = ?
		  AND sv_t.saving_type LIKE ?
		  AND sv.flow LIKE ?
	`

	rows, err := r.db.Raw(
		query,
		date,
		saving_type,
		flow,
	).Rows()
	if err != nil {
		return nil, err
	}

	var savingReadList []domain.SavingRead
	var tmpSavingRead domain.SavingRead
	for rows.Next() {
		rows.Scan(
			&tmpSavingRead.Date,
			&tmpSavingRead.SavingType,
			&tmpSavingRead.Flow,
			&tmpSavingRead.Amount,
		)

		savingReadList = append(savingReadList, tmpSavingRead)
	}
	return savingReadList, nil
}

func (r postgresSavingRepository) GetByID(id int64) (domain.SavingRead, error) {
	query := `
		SELECT
		  sv."date",
		  sv_t.saving_type,
		  sv.flow,
		  sv.amount
		FROM savings AS sv
		JOIN ref_saving_types AS sv_t
		  ON sv.saving_type_id = sv_t.id
		WHERE TRUE
		  AND sv.id = ?
	`

	var savingRead domain.SavingRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&savingRead.Date,
		&savingRead.SavingType,
		&savingRead.Flow,
		&savingRead.Amount,
	)
	if err != nil {
		return domain.SavingRead{}, err
	}
	return savingRead, nil
}

func (r postgresSavingRepository) Update(id int64, saving domain.Saving) (domain.Saving, error) {
	query := `
		UPDATE savings
		SET
		  "date" = ?,
		  saving_type_id = ?,
		  flow = ?,
		  amount = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		saving.Date,
		saving.SavingTypeID,
		saving.Flow,
		saving.Amount,
		id,
	).Row()

	var updatedSaving domain.Saving
	err := row.Scan(
		&updatedSaving.ID,
		&updatedSaving.Date,
		&updatedSaving.SavingTypeID,
		&updatedSaving.Flow,
		&updatedSaving.Amount,
	)
	if err != nil {
		return domain.Saving{}, err
	}
	return updatedSaving, nil
}

func (r postgresSavingRepository) Delete(id int64) (domain.Saving, error) {
	query := `
		DELETE FROM savings
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedSaving domain.Saving
	err := row.Scan(
		&deletedSaving.ID,
		&deletedSaving.Date,
		&deletedSaving.SavingTypeID,
		&deletedSaving.Flow,
		&deletedSaving.Amount,
	)
	if err != nil {
		return domain.Saving{}, err
	}
	return deletedSaving, nil
}
