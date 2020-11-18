package models

//Contractor ...
type Contractor struct {
	Model
	Name      string `json:"name" validate:"required"`
	NameEng   string `json:"name_eng" validate:"-"`
	Address   string `json:"address" validate:"-"`
	Telephone string `json:"telephone" validate:"-"`
	Fax       string `json:"fax" validate:"-"`
}
