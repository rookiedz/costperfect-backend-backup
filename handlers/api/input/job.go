package input

import "costperfect/backend/models"

//Job ...
type Job struct {
	TypeID  *int64  `validate:"required"`
	GroupID *int64  `validate:"required"`
	Label   *string `validate:"required"`
}

//Match ...
func (j Job) Match(job *models.Job) {
	if j.TypeID != nil {
		job.TypeID = *j.TypeID
	}
	if j.GroupID != nil {
		job.GroupID = *j.GroupID
	}
	if j.Label != nil {
		job.Label = *j.Label
	}
}
