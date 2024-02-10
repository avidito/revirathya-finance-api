package ref_income_types

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SeederRefIncomeTypes struct {
	db   *gorm.DB
	data [][]string
}

func NewSeederRefIncomeTypes(db *gorm.DB, data [][]string) *SeederRefIncomeTypes {
	return &SeederRefIncomeTypes{
		db:   db,
		data: data,
	}
}

func (s SeederRefIncomeTypes) CreateQuery() (string, []interface{}) {
	var listParamSet []string
	for i := 0; i < len(s.data); i++ {
		listParamSet = append(listParamSet, "(?, ?)")
	}
	query := fmt.Sprintf(`
		INSERT INTO ref_income_types (id, income_type)
		VALUES %s
		ON CONFLICT ON CONSTRAINT ref_income_types_pkey DO
			UPDATE SET
				income_type = excluded.income_type
	`, strings.Join(listParamSet, ", "))

	valueSet := make([]interface{}, 0, len(s.data)*2)
	for _, row := range s.data {

		id_val, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		valueSet = append(valueSet, id_val) // ID
		valueSet = append(valueSet, row[1]) // Income Types
	}

	return query, valueSet
}

func (s SeederRefIncomeTypes) Load() {
	fmt.Println("[SEED] Seeding: 'ref_income_types'")
	query, values := s.CreateQuery()
	s.db.Exec(query, values...)
}
