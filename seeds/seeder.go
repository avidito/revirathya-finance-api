package seeds

import (
	"github.com/avidito/revirathya-finance-api/seeds/ref_budget_groups"
	"github.com/avidito/revirathya-finance-api/seeds/ref_budget_types"
	"github.com/avidito/revirathya-finance-api/seeds/ref_income_types"
	"github.com/avidito/revirathya-finance-api/seeds/ref_locations"
	"github.com/avidito/revirathya-finance-api/seeds/ref_saving_types"
	"gorm.io/gorm"
)

type SeedData struct {
	seedDataRefIncomeTypes  [][]string
	seedDataRefLocations    [][]string
	seedDataRefBudgetGroups [][]string
	seedDataRefBudgetTypes  [][]string
	seedDataRefSavingTypes  [][]string
}

type Seeder struct {
	seederRefIncomeTypes  *ref_income_types.SeederRefIncomeTypes
	seederRefLocations    *ref_locations.SeederRefLocations
	seederRefBudgetGroups *ref_budget_groups.SeederRefBudgetGroups
	seederRefBudgetTypes  *ref_budget_types.SeederRefBudgetTypes
	seederRefSavingTypes  *ref_saving_types.SeederRefSavingTypes
}

func NewSeedData() *SeedData {
	return &SeedData{
		seedDataRefIncomeTypes:  readSeedCSV("ref_income_types"),
		seedDataRefLocations:    readSeedCSV("ref_locations"),
		seedDataRefBudgetGroups: readSeedCSV("ref_budget_groups"),
		seedDataRefBudgetTypes:  readSeedCSV("ref_budget_types"),
		seedDataRefSavingTypes:  readSeedCSV("ref_saving_types"),
	}
}

func NewSeeder(db *gorm.DB, seedData *SeedData) *Seeder {
	return &Seeder{
		seederRefIncomeTypes:  ref_income_types.NewSeederRefIncomeTypes(db, seedData.seedDataRefIncomeTypes),
		seederRefLocations:    ref_locations.NewSeederRefLocations(db, seedData.seedDataRefLocations),
		seederRefBudgetGroups: ref_budget_groups.NewSeederRefBudgetGroups(db, seedData.seedDataRefBudgetGroups),
		seederRefBudgetTypes:  ref_budget_types.NewSeederRefBudgetTypes(db, seedData.seedDataRefBudgetTypes),
		seederRefSavingTypes:  ref_saving_types.NewSeederRefSavingTypes(db, seedData.seedDataRefSavingTypes),
	}
}

func (s Seeder) Load() {
	s.seederRefIncomeTypes.Load()
	s.seederRefLocations.Load()
	s.seederRefBudgetGroups.Load()
	s.seederRefBudgetTypes.Load()
	s.seederRefSavingTypes.Load()
}
