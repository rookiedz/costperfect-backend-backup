package models

//Project ...
type Project struct {
	Model
	Name         string `json:"name"`
	OwnerName    string `json:"owner_name"`
	OwnerNameEng string `json:"owner_name_eng"`
	Director     string `json:"director"`
	Supervisor   string `json:"supervisor"`
}
