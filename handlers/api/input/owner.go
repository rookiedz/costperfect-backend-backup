package input

import "costperfect/backend/models"

//Owner ...
type Owner struct {
	Name     *string `validate:"required"`
	NameEng  *string `validate:"-"`
	Director *string `validate:"require"`
}

//Match ...
func (o Owner) Match(owner *models.Owner) {
	if o.Name != nil {
		owner.Name = *o.Name
	}
	if o.NameEng != nil {
		owner.NameEng = *o.NameEng
	}
	if o.Director != nil {
		owner.Director = *o.Director
	}
}
