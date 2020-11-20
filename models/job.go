package models

//JobType ...
type JobType struct {
	Model
	Label string `json:"label"`
}

//JobGroup ...
type JobGroup struct {
	Model
	Label  string `json:"label"`
	TypeID int64  `json:"type_id"`
}

//Job ...
type Job struct {
	Model
	Description string `json:"description"`
	TypeID      int64  `json:"type_id"`
	GroupID     int64  `json:"group_id"`
}
