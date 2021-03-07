package models

import (
	"encoding/json"
	"strings"
	"time"
)

//JSONDate ...
type JSONDate time.Time

//UnmarshalJSON ...
func (d *JSONDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = JSONDate(t)
	return nil
}

//MarshalJSON ...
func (d JSONDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

//Format ...
func (d JSONDate) Format(s string) string {
	t := time.Time(d)
	return t.Format(s)
}
