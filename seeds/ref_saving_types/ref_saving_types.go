package ref_saving_types

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SeederRefSavingTypes struct {
	db   *gorm.DB
	data [][]string
}

func NewSeederRefSavingTypes(db *gorm.DB, data [][]string) *SeederRefSavingTypes {
	return &SeederRefSavingTypes{
		db:   db,
		data: data,
	}
}

func (s SeederRefSavingTypes) CreateQuery() (string, []interface{}) {
	var listParamSet []string
	for i := 0; i < len(s.data); i++ {
		listParamSet = append(listParamSet, "(?, ?)")
	}
	query := fmt.Sprintf(`
		INSERT INTO ref_saving_types (id, saving_type)
		VALUES %s
		ON CONFLICT ON CONSTRAINT ref_saving_types_pkey DO
			UPDATE SET
				saving_type = excluded.saving_type
	`, strings.Join(listParamSet, ", "))

	valueSet := make([]interface{}, 0, len(s.data)*2)
	for _, row := range s.data {

		id_val, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		valueSet = append(valueSet, id_val) // ID
		valueSet = append(valueSet, row[1]) // Saving Type
	}
	return query, valueSet
}

func (s SeederRefSavingTypes) Load() {
	fmt.Println("[SEED] Seeding: 'ref_saving_types'")
	query, values := s.CreateQuery()
	s.db.Exec(query, values...)
}
