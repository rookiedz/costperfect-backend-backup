package models

import "time"

//Model ...
type Model struct {
	ID        int64      `json:"id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
