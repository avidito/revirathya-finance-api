package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresCycleRepository struct {
	db *gorm.DB
}

func NewPostgresCycleRepository(db *gorm.DB) domain.CycleRepository {
	return &postgresCycleRepository{
		db: db,
	}
}

// Repository
func (r postgresCycleRepository) Create(cycle domain.Cycle) (domain.Cycle, error) {
	query := `
		INSERT INTO cycles (
			label,
			start_date,
			end_date,
			notes,
			total_days
		)
		VALUES (?, ?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		cycle.Label,
		cycle.StartDate,
		cycle.EndDate,
		cycle.Notes,
		cycle.TotalDays,
	).Row()

	var createdCycle domain.Cycle
	err := row.Scan(
		&createdCycle.ID,
		&createdCycle.Label,
		&createdCycle.StartDate,
		&createdCycle.EndDate,
		&createdCycle.Notes,
		&createdCycle.TotalDays,
	)
	if err != nil {
		return domain.Cycle{}, err
	}

	return createdCycle, nil
}

func (r postgresCycleRepository) Fetch(label string) ([]domain.CycleRead, error) {
	query := `
		SELECT
		  id,
		  label,
		  start_date,
		  end_date,
		  notes,
		  total_days
		FROM cycles
		WHERE TRUE
		  AND label LIKE ?
	`

	rows, err := r.db.Raw(query, label).Rows()

	if err != nil {
		return nil, err
	}

	var cycleReadList []domain.CycleRead
	var tmpCycleRead domain.CycleRead
	for rows.Next() {
		rows.Scan(
			&tmpCycleRead.ID,
			&tmpCycleRead.Label,
			&tmpCycleRead.StartDate,
			&tmpCycleRead.EndDate,
			&tmpCycleRead.Notes,
			&tmpCycleRead.TotalDays,
		)
		cycleReadList = append(cycleReadList, tmpCycleRead)
	}
	return cycleReadList, nil
}

func (r postgresCycleRepository) GetByID(id int64) (domain.CycleRead, error) {
	query := `
		SELECT
		  id,
		  label,
		  start_date,
		  end_date,
		  notes,
		  total_days
		FROM cycles
		WHERE TRUE
		  AND id = ?
	`

	var cycleRead domain.CycleRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&cycleRead.ID,
		&cycleRead.Label,
		&cycleRead.StartDate,
		&cycleRead.EndDate,
		&cycleRead.Notes,
		&cycleRead.TotalDays,
	)

	if err != nil {
		return domain.CycleRead{}, err
	}
	return cycleRead, nil
}

func (r postgresCycleRepository) Update(id int64, cycle domain.Cycle) (domain.Cycle, error) {
	query := `
		UPDATE cycles
		SET
		  label = ?,
		  start_date = ?,
		  end_date = ?,
		  notes = ?,
		  total_days = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		cycle.Label,
		cycle.StartDate,
		cycle.EndDate,
		cycle.Notes,
		cycle.TotalDays,
		id,
	).Row()

	var updatedCycle domain.Cycle
	err := row.Scan(
		&updatedCycle.ID,
		&updatedCycle.Label,
		&updatedCycle.StartDate,
		&updatedCycle.EndDate,
		&updatedCycle.Notes,
		&updatedCycle.TotalDays,
	)
	if err != nil {
		return domain.Cycle{}, err
	}
	return updatedCycle, nil
}

func (r postgresCycleRepository) Delete(id int64) (domain.Cycle, error) {
	query := `
		DELETE FROM cycles
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedCycle domain.Cycle
	err := row.Scan(
		&deletedCycle.ID,
		&deletedCycle.Label,
		&deletedCycle.StartDate,
		&deletedCycle.EndDate,
		&deletedCycle.Notes,
		&deletedCycle.TotalDays,
	)
	if err != nil {
		return domain.Cycle{}, err
	}
	return deletedCycle, nil
}
