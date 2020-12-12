package models

//Employer ...
type Employer struct {
	Model
	Fullname  string `json:"fullname,omitempty"`
	ProjectID int64  `json:"project_id"`
}
