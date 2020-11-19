package input

import "costperfect/backend/models"

//JobType ...
type JobType struct {
	Label *string `validate:"required"`
}

//Match ...
func (jt *JobType) Match(jobType *models.JobType) {
	if jt.Label != nil {
		jobType.Label = *jt.Label
	}
}
