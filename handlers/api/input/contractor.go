package input

import (
	"costperfect/backend/models"
)

//Contractor ...
type Contractor struct {
	Name      *string `json:"name" validate:"required"`
	NameEng   *string `json:"name_eng" validate:"-"`
	Acronym   *string `json:"acronym" validate:"required"`
	Address   *string `json:"address" validate:"-"`
	Telephone *string `json:"telephone" validate:"-"`
	Fax       *string `json:"fax" validate:"-"`
}

//Match ...
func (c Contractor) Match(contractor *models.Contractor) {
	if c.Name != nil {
		contractor.Name = *c.Name
	}
	if c.NameEng != nil {
		contractor.NameEng = *c.NameEng
	}
	if c.Acronym != nil {
		contractor.Acronym = *c.Acronym
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
