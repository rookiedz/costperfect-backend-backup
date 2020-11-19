package input

import "costperfect/backend/models"

//Project ...
type Project struct {
	Name    *string `validate:"required"`
	OwnerID *int64  `validate:"required"`
}

//Match ...
func (p Project) Match(project *models.Project) {
	if p.Name != nil {
		project.Name = *p.Name
	}
	if p.OwnerID != nil {
		project.OwnerID = *p.OwnerID
	}
}
