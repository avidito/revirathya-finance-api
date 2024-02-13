package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DateStandard struct {
	time.Time
}

func (ds *DateStandard) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	ds.Time = t
	return nil
}

func (ds DateStandard) MarshalJSON() ([]byte, error) {
	return json.Marshal(ds.Time.Format("2006-01-02"))
}

func (ds DateStandard) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := ds.Time

	// Check whether the given time is equal to the default zero time
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (ds *DateStandard) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*ds = DateStandard{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to date (standard)", v)
}
