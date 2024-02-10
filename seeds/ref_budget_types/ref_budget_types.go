package ref_budget_types

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SeederRefBudgetTypes struct {
	db   *gorm.DB
	data [][]string
}

func NewSeederRefBudgetTypes(db *gorm.DB, data [][]string) *SeederRefBudgetTypes {
	return &SeederRefBudgetTypes{
		db:   db,
		data: data,
	}
}

func (s SeederRefBudgetTypes) CreateQuery() (string, []interface{}) {
	var listParamSet []string
	for i := 0; i < len(s.data); i++ {
		listParamSet = append(listParamSet, "(?, ?, ?)")
	}
	query := fmt.Sprintf(`
		INSERT INTO ref_budget_types (id, budget_group_id, budget_type)
		VALUES %s
		ON CONFLICT ON CONSTRAINT ref_budget_types_pkey DO
			UPDATE SET
				budget_group_id = excluded.budget_group_id,
				budget_type = excluded.budget_type
	`, strings.Join(listParamSet, ", "))

	valueSet := make([]interface{}, 0, len(s.data)*3)
	for _, row := range s.data {

		id_val, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		bg_id_val, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		valueSet = append(valueSet, id_val)    // ID
		valueSet = append(valueSet, bg_id_val) // ID Budget Group
		valueSet = append(valueSet, row[2])    // Budget Type
	}

	return query, valueSet
}

func (s SeederRefBudgetTypes) Load() {
	fmt.Println("[SEED] Seeding: 'ref_budget_types'")
	query, values := s.CreateQuery()
	s.db.Exec(query, values...)
}
