package input

import "costperfect/backend/models"

//JobGroup ...
type JobGroup struct {
	TypeID *int64  `json:"label" validate:"required"`
	Label  *string `json:"type_id" validate:"required"`
}

//Match ...
func (jg JobGroup) Match(jobGroup *models.JobGroup) {
	if jg.TypeID != nil {
		jobGroup.TypeID = *jg.TypeID
	}
	if jg.Label != nil {
		jobGroup.Label = *jg.Label
	}
}
