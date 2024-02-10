package ref_locations

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SeederRefLocations struct {
	db   *gorm.DB
	data [][]string
}

func NewSeederRefLocations(db *gorm.DB, data [][]string) *SeederRefLocations {
	return &SeederRefLocations{
		db:   db,
		data: data,
	}
}

func (s SeederRefLocations) CreateQuery() (string, []interface{}) {
	var listParamSet []string
	for i := 0; i < len(s.data); i++ {
		listParamSet = append(listParamSet, "(?, ?)")
	}
	query := fmt.Sprintf(`
		INSERT INTO ref_locations (id, location)
		VALUES %s
		ON CONFLICT ON CONSTRAINT ref_locations_pkey DO
			UPDATE SET
				location = excluded.location
	`, strings.Join(listParamSet, ", "))

	valueSet := make([]interface{}, 0, len(s.data)*2)
	for _, row := range s.data {

		id_val, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		valueSet = append(valueSet, id_val) // ID
		valueSet = append(valueSet, row[1]) // Location
	}

	return query, valueSet
}

func (s SeederRefLocations) Load() {
	fmt.Println("[SEED] Seeding: 'ref_locations'")
	query, values := s.CreateQuery()
	s.db.Exec(query, values...)
}
