package nullType

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Valid = false
		return
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return
	case []byte:
		nt.Time, err = parseDateTime(string(v), time.UTC)
		nt.Valid = err == nil
		return
	case string:
		if v == "" {
			nt.Valid = false
			return
		}
		nt.Time, err = parseDateTime(v, time.UTC)
		nt.Valid = err == nil
		return
	}

	nt.Valid = false
	return fmt.Errorf("Can't convert %T to time.Time", value)
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func parseDateTime(str string, loc *time.Location) (t time.Time, err error) {
	t, err = time.Parse(
		"2006-01-02T03:04:05-07:00", str)
	return
}
