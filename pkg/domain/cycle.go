package domain

type Cycle struct {
	ID        int64        `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Label     string       `json:"label"`
	StartDate DateStandard `json"start_date"`
	EndDate   DateStandard `json:"end_date"`
	Notes     string       `json:"notes"`
	TotalDays int64        `json:"total_days"`
}

type CycleRead struct {
	ID        int64        `json:"id"`
	Label     string       `json:"label"`
	StartDate DateStandard `json"start_date"`
	EndDate   DateStandard `json:"end_date"`
	Notes     string       `json:"notes"`
	TotalDays int64        `json:"total_days"`
}

type CycleRepository interface {
	Create(cycle Cycle) (Cycle, error)
	Fetch(label string) ([]CycleRead, error)
	GetByID(id int64) (CycleRead, error)
	Update(id int64, cycle Cycle) (Cycle, error)
	Delete(id int64) (Cycle, error)
}

type CycleUsecase interface {
	Create(cycle Cycle) (Cycle, error)
	Fetch(label string) ([]CycleRead, error)
	GetByID(id int64) (CycleRead, error)
	Update(id int64, cycle Cycle) (Cycle, error)
	Delete(id int64) (Cycle, error)
}
