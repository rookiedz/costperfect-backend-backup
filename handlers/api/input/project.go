package input

import "costperfect/backend/models"

//Project ...
type Project struct {
	Name         *string `validate:"required"`
	OwnerName    *string `validate:"required"`
	OwnerNameEng *string `validate:"required"`
	Manager      *string `validate:"required"`
	Acronym      *string `validate:"required"`
	Expand       *string `validate:"required"`
}

//Match ...
func (p Project) Match(project *models.Project) {
	if p.Name != nil {
		project.Name = *p.Name
	}
	if p.OwnerName != nil {
		project.OwnerName = *p.OwnerName
	}
	if p.OwnerNameEng != nil {
		project.OwnerNameEng = *p.OwnerNameEng
	}
	if p.Manager != nil {
		project.Manager = *p.Manager
	}
	if p.Acronym != nil {
		project.Acronym = *p.Acronym
	}
	if p.Expand != nil {
		project.Expand = *p.Expand
	}
}
