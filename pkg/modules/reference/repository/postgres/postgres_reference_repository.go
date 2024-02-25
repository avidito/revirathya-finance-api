package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresReferenceRepository struct {
	db *gorm.DB
}

func NewPostgresReferenceRepository(db *gorm.DB) domain.ReferenceRepository {
	return &postgresReferenceRepository{
		db: db,
	}
}

// Repository
func (r postgresReferenceRepository) GetIncomeTypeByID(id int64) (domain.RefIncomeType, error) {
	query := `
		SELECT
		  id,
		  income_type
		FROM ref_income_types
		WHERE id = ?
	`

	var refIncomeTypeRead domain.RefIncomeType
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&refIncomeTypeRead.ID,
		&refIncomeTypeRead.IncomeType,
	)
	if err != nil {
		return domain.RefIncomeType{}, err
	}
	return refIncomeTypeRead, nil
}

func (r postgresReferenceRepository) GetLocationByID(id int64) (domain.RefLocation, error) {
	query := `
		SELECT
		  id,
		  location
		FROM ref_locations
		WHERE id = ?
	`

	var refLocationRead domain.RefLocation
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&refLocationRead.ID,
		&refLocationRead.Location,
	)
	if err != nil {
		return domain.RefLocation{}, err
	}
	return refLocationRead, nil
}

func (r postgresReferenceRepository) GetBudgetGroupByID(id int64) (domain.RefBudgetGroup, error) {
	query := `
		SELECT
		  id,
		  budget_group
		FROM ref_budget_groups
		WHERE id = ?
	`

	var refBudgetGroupRead domain.RefBudgetGroup
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&refBudgetGroupRead.ID,
		&refBudgetGroupRead.BudgetGroup,
	)
	if err != nil {
		return domain.RefBudgetGroup{}, err
	}
	return refBudgetGroupRead, nil
}

func (r postgresReferenceRepository) GetBudgetTypeByID(id int64) (domain.RefBudgetTypeRead, error) {
	query := `
		SELECT
		  bt.id,
		  bg.budget_group,
		  bt.budget_type
		FROM ref_budget_types AS bt
		JOIN ref_budget_groups AS bg
		  ON bt.budget_group_id = bg.id
		WHERE bt.id = ?
	`

	var refBudgetTypeRead domain.RefBudgetTypeRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&refBudgetTypeRead.ID,
		&refBudgetTypeRead.BudgetGroup,
		&refBudgetTypeRead.BudgetType,
	)
	if err != nil {
		return domain.RefBudgetTypeRead{}, err
	}
	return refBudgetTypeRead, nil
}

func (r postgresReferenceRepository) GetSavingTypeByID(id int64) (domain.RefSavingType, error) {
	query := `
		SELECT
		  id,
		  saving_type
		FROM ref_saving_types
		WHERE id = ?
	`

	var refSavingTypeRead domain.RefSavingType
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&refSavingTypeRead.ID,
		&refSavingTypeRead.SavingType,
	)
	if err != nil {
		return domain.RefSavingType{}, err
	}
	return refSavingTypeRead, nil
}

func (r postgresReferenceRepository) FetchIncomeTypes(name string) ([]domain.RefIncomeType, error) {
	query := `
		SELECT
		  id,
		  income_type
		FROM ref_income_types
		WHERE income_type LIKE ?
	`

	rows, err := r.db.Raw(query, name).Rows()
	if err != nil {
		return nil, err
	}

	var refIncomeTypeList []domain.RefIncomeType
	var tmpRefIncomeType domain.RefIncomeType
	for rows.Next() {
		rows.Scan(
			&tmpRefIncomeType.ID,
			&tmpRefIncomeType.IncomeType,
		)

		refIncomeTypeList = append(refIncomeTypeList, tmpRefIncomeType)
	}
	return refIncomeTypeList, nil
}

func (r postgresReferenceRepository) FetchLocations(name string) ([]domain.RefLocation, error) {
	query := `
		SELECT
		  id,
		  location
		FROM ref_locations
		WHERE location LIKE ?
	`

	rows, err := r.db.Raw(query, name).Rows()
	if err != nil {
		return nil, err
	}

	var refLocationList []domain.RefLocation
	var tmpRefLocation domain.RefLocation
	for rows.Next() {
		rows.Scan(
			&tmpRefLocation.ID,
			&tmpRefLocation.Location,
		)

		refLocationList = append(refLocationList, tmpRefLocation)
	}
	return refLocationList, nil
}

func (r postgresReferenceRepository) FetchBudgetGroups(name string) ([]domain.RefBudgetGroup, error) {
	query := `
		SELECT
		  id,
		  budget_group
		FROM ref_budget_groups
		WHERE budget_group LIKE ?
	`

	rows, err := r.db.Raw(query, name).Rows()
	if err != nil {
		return nil, err
	}

	var refBudgetGroupList []domain.RefBudgetGroup
	var tmpRefBudgetGroup domain.RefBudgetGroup
	for rows.Next() {
		rows.Scan(
			&tmpRefBudgetGroup.ID,
			&tmpRefBudgetGroup.BudgetGroup,
		)

		refBudgetGroupList = append(refBudgetGroupList, tmpRefBudgetGroup)
	}
	return refBudgetGroupList, nil
}

func (r postgresReferenceRepository) FetchBudgetTypes(name string, budget_group string) ([]domain.RefBudgetTypeRead, error) {
	query := `
		SELECT
		  bt.id,
		  bg.budget_group,
		  bt.budget_type
		FROM ref_budget_types AS bt
		JOIN ref_budget_groups AS bg
		  ON bt.budget_group_id = bg.id
		WHERE budget_type LIKE ?
		  AND budget_group LIKE ?
	`

	rows, err := r.db.Raw(query, name, budget_group).Rows()
	if err != nil {
		return nil, err
	}

	var refBudgetTypeReadList []domain.RefBudgetTypeRead
	var tmpRefBudgetTypeRead domain.RefBudgetTypeRead
	for rows.Next() {
		rows.Scan(
			&tmpRefBudgetTypeRead.ID,
			&tmpRefBudgetTypeRead.BudgetGroup,
			&tmpRefBudgetTypeRead.BudgetType,
		)

		refBudgetTypeReadList = append(refBudgetTypeReadList, tmpRefBudgetTypeRead)
	}
	return refBudgetTypeReadList, nil
}

func (r postgresReferenceRepository) FetchSavingTypes(name string) ([]domain.RefSavingType, error) {
	query := `
		SELECT
		  id,
		  saving_type
		FROM ref_saving_types
		WHERE saving_type LIKE ?
	`

	rows, err := r.db.Raw(query, name).Rows()
	if err != nil {
		return nil, err
	}

	var refSavingTypeList []domain.RefSavingType
	var tmpRefSavingType domain.RefSavingType
	for rows.Next() {
		rows.Scan(
			&tmpRefSavingType.ID,
			&tmpRefSavingType.SavingType,
		)

		refSavingTypeList = append(refSavingTypeList, tmpRefSavingType)
	}
	return refSavingTypeList, nil
}
