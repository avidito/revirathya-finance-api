package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresIncomeRepository struct {
	db *gorm.DB
}

func NewPostgresIncomesRepository(db *gorm.DB) domain.IncomeRepository {
	return &postgresIncomeRepository{
		db: db,
	}
}

// Repository
func (r postgresIncomeRepository) Create(income domain.Income) (domain.Income, error) {
	query := `
		INSERT INTO incomes (
		  "date",
		  income_type_id,
		  location_id,
		  description,
		  amount
		)
		VALUES (?, ?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		income.Date,
		income.IncomeTypeID,
		income.LocationID,
		income.Description,
		income.Amount,
	).Row()

	var createdIncome domain.Income
	err := row.Scan(
		&createdIncome.ID,
		&createdIncome.Date,
		&createdIncome.IncomeTypeID,
		&createdIncome.LocationID,
		&createdIncome.Description,
		&createdIncome.Amount,
	)
	if err != nil {
		return domain.Income{}, err
	}

	return createdIncome, nil
}

func (r postgresIncomeRepository) Fetch(_date domain.DateStandard, income_type string) ([]domain.IncomeRead, error) {
	var tmpIncome domain.IncomeRead
	var incomes []domain.IncomeRead
	query := `
		SELECT
		  inc.date,
		  inc_t.income_type,
		  loc.location,
		  inc.description,
		  inc.amount
		FROM incomes AS inc
		JOIN ref_income_types AS inc_t
		  ON inc.income_type_id = inc_t.id
		JOIN ref_locations AS loc
		  ON inc.location_id = loc.id
		WHERE TRUE
		  AND inc.date = ?
		  AND inc_t.income_type LIKE '%' || ? || '%'
	`

	rows, err := r.db.Raw(
		query,
		_date,
		income_type,
	).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(
			&tmpIncome.Date,
			&tmpIncome.IncomeType,
			&tmpIncome.Location,
			&tmpIncome.Description,
			&tmpIncome.Amount,
		)
		incomes = append(incomes, tmpIncome)
	}
	return incomes, nil
}

func (r postgresIncomeRepository) Get(id int64) (domain.IncomeRead, error) {
	var income domain.IncomeRead
	query := `
		SELECT
		  inc.date,
		  inc_t.income_type,
		  loc.location,
		  inc.description,
		  inc.amount
		FROM incomes AS inc
		JOIN ref_income_types AS inc_t
		  ON inc.income_type_id = inc_t.id
		JOIN ref_locations AS loc
		  ON inc.location_id = loc.id
		WHERE TRUE
		  AND inc.id = ?
	`

	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&income.Date,
		&income.IncomeType,
		&income.Location,
		&income.Description,
		&income.Amount,
	)
	if err != nil {
		return domain.IncomeRead{}, err
	}
	return income, nil
}

func (r postgresIncomeRepository) Update(id int64, income domain.Income) (domain.Income, error) {
	query := `
		UPDATE incomes
		SET
		  date = ?,
		  income_type_id = ?,
		  location_id = ?,
		  description = ?,
		  amount = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		income.Date,
		income.IncomeTypeID,
		income.LocationID,
		income.Description,
		income.Amount,
		id,
	).Row()

	var updatedIncome domain.Income
	err := row.Scan(
		&updatedIncome.ID,
		&updatedIncome.Date,
		&updatedIncome.IncomeTypeID,
		&updatedIncome.LocationID,
		&updatedIncome.Description,
		&updatedIncome.Amount,
	)
	if err != nil {
		return domain.Income{}, err
	}
	return updatedIncome, nil
}

func (r postgresIncomeRepository) Delete(id int64) (domain.Income, error) {
	query := `
		DELETE FROM incomes
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedIncome domain.Income
	err := row.Scan(
		&deletedIncome.ID,
		&deletedIncome.Date,
		&deletedIncome.IncomeTypeID,
		&deletedIncome.LocationID,
		&deletedIncome.Description,
		&deletedIncome.Amount,
	)
	if err != nil {
		return domain.Income{}, err
	}
	return deletedIncome, err
}
