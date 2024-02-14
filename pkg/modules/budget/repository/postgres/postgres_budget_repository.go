package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresBudgetRepository struct {
	db *gorm.DB
}

func NewPostgresBudgetRepository(db *gorm.DB) domain.BudgetRepository {
	return &postgresBudgetRepository{
		db: db,
	}
}

// Repository
func (r postgresBudgetRepository) Create(budget domain.Budget) (domain.Budget, error) {
	query := `
		INSERT INTO budgets (
			cycle,
			budget_type_id,
			amount
		)
		VALUES (?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		budget.Cycle,
		budget.BudgetTypeID,
		budget.Amount,
	).Row()

	var createdBudget domain.Budget
	err := row.Scan(
		&createdBudget.ID,
		&createdBudget.Cycle,
		&createdBudget.BudgetTypeID,
		&createdBudget.Amount,
	)
	if err != nil {
		return domain.Budget{}, err
	}

	return createdBudget, nil
}

func (r postgresBudgetRepository) Fetch(cycle string, budget_type string) ([]domain.BudgetRead, error) {
	query := `
		SELECT
		  bud.cycle,
		  bud_t.budget_type,
		  bud.amount
		FROM budgets AS bud
		JOIN ref_budget_types AS bud_t
		  ON bud.budget_type_id = bud_t.id
		WHERE TRUE
		  AND bud.cycle LIKE ?
		  AND bud_t.budget_type LIKE ?
	`

	rows, err := r.db.Raw(
		query,
		cycle,
		budget_type,
	).Rows()
	if err != nil {
		return nil, err
	}

	var budgetReadList []domain.BudgetRead
	var tmpBudgetRead domain.BudgetRead
	for rows.Next() {
		rows.Scan(
			&tmpBudgetRead.Cycle,
			&tmpBudgetRead.BudgetType,
			&tmpBudgetRead.Amount,
		)

		budgetReadList = append(budgetReadList, tmpBudgetRead)
	}
	return budgetReadList, nil
}

func (r postgresBudgetRepository) GetByID(id int64) (domain.BudgetRead, error) {
	query := `
		SELECT
		  bud.cycle,
		  bud_t.budget_type,
		  bud.amount
		FROM budgets AS bud
		JOIN ref_budget_types AS bud_t
		  ON bud.budget_type_id = bud_t.id
		WHERE TRUE
		  AND bud.id = ?
	`

	var budgetRead domain.BudgetRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&budgetRead.Cycle,
		&budgetRead.BudgetType,
		&budgetRead.Amount,
	)
	if err != nil {
		return domain.BudgetRead{}, err
	}
	return budgetRead, nil
}

func (r postgresBudgetRepository) Update(id int64, budget domain.Budget) (domain.Budget, error) {
	query := `
		UPDATE budgets
		SET
		  cycle = ?,
		  budget_type_id = ?,
		  amount = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		budget.Cycle,
		budget.BudgetTypeID,
		budget.Amount,
		id,
	).Row()

	var updatedBudget domain.Budget
	err := row.Scan(
		&updatedBudget.ID,
		&updatedBudget.Cycle,
		&updatedBudget.BudgetTypeID,
		&updatedBudget.Amount,
	)
	if err != nil {
		return domain.Budget{}, err
	}
	return updatedBudget, nil
}

func (r postgresBudgetRepository) Delete(id int64) (domain.Budget, error) {
	query := `
		DELETE FROM budgets
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedBudget domain.Budget
	err := row.Scan(
		&deletedBudget.ID,
		&deletedBudget.Cycle,
		&deletedBudget.BudgetTypeID,
		&deletedBudget.Amount,
	)
	if err != nil {
		return domain.Budget{}, err
	}
	return deletedBudget, nil
}
