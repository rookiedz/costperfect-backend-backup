package models

//Employer ...
type Employer struct {
	Model
	Fullname  string `json:"fullname,omitempty"`
	ProjectID string `json:"project_id"`
}
