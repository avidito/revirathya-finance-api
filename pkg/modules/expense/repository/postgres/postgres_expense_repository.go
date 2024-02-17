package postgres

import (
	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type postgresExpenseRepository struct {
	db *gorm.DB
}

func NewPostgresExpenseRepository(db *gorm.DB) domain.ExpenseRepository {
	return &postgresExpenseRepository{
		db: db,
	}
}

// Repository
func (r postgresExpenseRepository) Create(expense domain.Expense) (domain.Expense, error) {
	query := `
		INSERT INTO expenses (
			"date",
			budget_type_id,
			location_id,
			description,
			amount
		)
		VALUES (?, ?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.Raw(
		query,
		expense.Date,
		expense.BudgetTypeID,
		expense.LocationID,
		expense.Description,
		expense.Amount,
	).Row()

	var createdExpense domain.Expense
	err := row.Scan(
		&createdExpense.ID,
		&createdExpense.Date,
		&createdExpense.BudgetTypeID,
		&createdExpense.LocationID,
		&createdExpense.Description,
		&createdExpense.Amount,
	)
	if err != nil {
		return domain.Expense{}, err
	}

	return createdExpense, nil
}

func (r postgresExpenseRepository) Fetch(date string, budget_type string) ([]domain.ExpenseRead, error) {
	query := `
		SELECT
		  ex."date",
		  bud_t.budget_type,
		  loc."location",
		  ex.description,
		  ex.amount
		FROM expenses AS ex
		JOIN ref_budget_types AS bud_t
		  ON ex.budget_type_id = bud_t.id
		JOIN ref_locations AS loc
		  ON ex.location_id = loc.id
		WHERE TRUE
		  AND DATE(ex."date") = ?
		  AND bud_t.budget_type LIKE ?
	`

	rows, err := r.db.Raw(
		query,
		date,
		budget_type,
	).Rows()
	if err != nil {
		return nil, err
	}

	var expenseReadList []domain.ExpenseRead
	var tmpExpenseRead domain.ExpenseRead
	for rows.Next() {
		rows.Scan(
			&tmpExpenseRead.Date,
			&tmpExpenseRead.BudgetType,
			&tmpExpenseRead.Location,
			&tmpExpenseRead.Description,
			&tmpExpenseRead.Amount,
		)

		expenseReadList = append(expenseReadList, tmpExpenseRead)
	}
	return expenseReadList, nil
}

func (r postgresExpenseRepository) GetByID(id int64) (domain.ExpenseRead, error) {
	query := `
		SELECT
		  ex."date",
		  bud_t.budget_type,
		  loc."location",
		  ex.description,
		  ex.amount
		FROM expenses AS ex
		JOIN ref_budget_types AS bud_t
		  ON ex.budget_type_id = bud_t.id
		JOIN ref_locations AS loc
		  ON ex.location_id = loc.id
		WHERE TRUE
		  AND ex.id = ?
	`

	var expenseRead domain.ExpenseRead
	row := r.db.Raw(query, id).Row()
	err := row.Scan(
		&expenseRead.Date,
		&expenseRead.BudgetType,
		&expenseRead.Location,
		&expenseRead.Description,
		&expenseRead.Amount,
	)
	if err != nil {
		return domain.ExpenseRead{}, err
	}
	return expenseRead, nil
}

func (r postgresExpenseRepository) Update(id int64, expense domain.Expense) (domain.Expense, error) {
	query := `
		UPDATE expenses
		SET
		  "date" = ?,
		  budget_type_id = ?,
		  location_id = ?,
		  description = ?,
		  amount = ?
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		expense.Date,
		expense.BudgetTypeID,
		expense.LocationID,
		expense.Description,
		expense.Amount,
		id,
	).Row()

	var updatedExpense domain.Expense
	err := row.Scan(
		&updatedExpense.ID,
		&updatedExpense.Date,
		&updatedExpense.BudgetTypeID,
		&updatedExpense.LocationID,
		&updatedExpense.Description,
		&updatedExpense.Amount,
	)
	if err != nil {
		return domain.Expense{}, err
	}
	return updatedExpense, nil
}

func (r postgresExpenseRepository) Delete(id int64) (domain.Expense, error) {
	query := `
		DELETE FROM expenses
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(query, id).Row()

	var deletedExpense domain.Expense
	err := row.Scan(
		&deletedExpense.ID,
		&deletedExpense.Date,
		&deletedExpense.BudgetTypeID,
		&deletedExpense.LocationID,
		&deletedExpense.Description,
		&deletedExpense.Amount,
	)
	if err != nil {
		return domain.Expense{}, err
	}
	return deletedExpense, nil
}
