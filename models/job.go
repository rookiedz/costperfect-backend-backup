package models

//JobType ...
type JobType struct {
	Model
	Name string `json:"name"`
}

//JobGroup ...
type JobGroup struct {
	Model
	Name   string `json:"name"`
	TypeID int64  `json:"type_id"`
}

//Job ...
type Job struct {
	Model
	Name     string `json:"name"`
	TypeID   int64  `json:"type_id"`
	GroupID  int64  `json:"group_id"`
	ParentID int64  `json:"parent_id"`
}
