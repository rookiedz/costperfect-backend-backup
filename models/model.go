package models

import "time"

//Model ...
type Model struct {
	ID        int64      `json:"id"`
	Deleted   bool       `json:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
