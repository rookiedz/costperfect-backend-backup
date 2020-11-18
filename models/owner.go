package models

//Owner ...
type Owner struct {
	Model
	Name     string `json:"name"`
	NameEng  string `json:"name_eng"`
	Director string `json:"director"`
}
