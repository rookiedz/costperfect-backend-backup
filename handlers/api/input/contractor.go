package input

import (
	"costperfect/backend/models"
)

//Contractor ...
type Contractor struct {
	ID        *int64  `validate:"-"`
	Name      *string `validate:"required"`
	NameEng   *string `validate:"-"`
	Address   *string `validate:"-"`
	Telephone *string `validate:"-"`
	Fax       *string `validate:"-"`
}

//Match ...
func (c Contractor) Match(contractor *models.Contractor) {
	if c.Name != nil {
		contractor.Name = *c.Name
	}
	if c.NameEng != nil {
		contractor.NameEng = *c.NameEng
	}
	if c.Address != nil {
		contractor.Address = *c.Address
	}
	if c.Telephone != nil {
		contractor.Telephone = *c.Telephone
	}
	if c.Fax != nil {
		contractor.Fax = *c.Fax
	}
}
