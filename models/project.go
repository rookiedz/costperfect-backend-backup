package models

//Project ...
type Project struct {
	Model
	Name         string `json:"name"`
	OwnerName    string `json:"owner_name"`
	OwnerNameEng string `json:"owner_name_eng"`
	Manager      string `json:"manager"`
}
