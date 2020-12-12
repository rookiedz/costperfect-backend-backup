package input

import "costperfect/backend/models"

//Employer ...
type Employer struct {
	Fullname  *string `validate:"required"`
	ProjectID *int64  `validate:"required"`
}

//Match ...
func (e Employer) Match(employer *models.Employer) {
	if e.Fullname != nil {
		employer.Fullname = *e.Fullname
	}
	if e.ProjectID != nil {
		employer.ProjectID = *e.ProjectID
	}
}
