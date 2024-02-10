package ref_budget_groups

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SeederRefBudgetGroups struct {
	db   *gorm.DB
	data [][]string
}

func NewSeederRefBudgetGroups(db *gorm.DB, data [][]string) *SeederRefBudgetGroups {
	return &SeederRefBudgetGroups{
		db:   db,
		data: data,
	}
}

func (s SeederRefBudgetGroups) CreateQuery() (string, []interface{}) {
	var listParamSet []string
	for i := 0; i < len(s.data); i++ {
		listParamSet = append(listParamSet, "(?, ?)")
	}
	query := fmt.Sprintf(`
		INSERT INTO ref_budget_groups (id, budget_group)
		VALUES %s
		ON CONFLICT ON CONSTRAINT ref_budget_groups_pkey DO
			UPDATE SET
				budget_group = excluded.budget_group
	`, strings.Join(listParamSet, ", "))

	valueSet := make([]interface{}, 0, len(s.data)*2)
	for _, row := range s.data {

		id_val, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		valueSet = append(valueSet, id_val) // ID
		valueSet = append(valueSet, row[1]) // Budget Group
	}

	return query, valueSet
}

func (s SeederRefBudgetGroups) Load() {
	fmt.Println("[SEED] Seeding: 'ref_budget_groups'")
	query, values := s.CreateQuery()
	s.db.Exec(query, values...)
}
