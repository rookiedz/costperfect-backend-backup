package models

//Contractor ...
type Contractor struct {
	Model
	Name      string `json:"name"`
	NameEng   string `json:"name_eng"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Fax       string `json:"fax"`
}
