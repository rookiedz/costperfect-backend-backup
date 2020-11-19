package input

import "costperfect/backend/models"

//JobGroup ...
type JobGroup struct {
	TypeID *int64  `validate:"required"`
	Label  *string `validate:"required"`
}

//Match ...
func (jg *JobGroup) Match(jobGroup *models.JobGroup) {
	if jg.TypeID != nil {
		jobGroup.TypeID = *jg.TypeID
	}
	if jg.Label != nil {
		jobGroup.Label = *jg.Label
	}
}
