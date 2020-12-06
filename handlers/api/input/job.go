package input

import "costperfect/backend/models"

//Job ...
type Job struct {
	TypeID      *int64  `json:"type_id" validate:"required"`
	GroupID     *int64  `json:"group_id" validate:"required"`
	Description *string `json:"description" validate:"required"`
}

//Match ...
func (j Job) Match(job *models.Job) {
	if j.TypeID != nil {
		job.TypeID = *j.TypeID
	}
	if j.GroupID != nil {
		job.GroupID = *j.GroupID
	}
	if j.Description != nil {
		job.Description = *j.Description
	}
}
