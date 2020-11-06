package models

import "time"

//Model ...
type Model struct {
	ID        int64
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
