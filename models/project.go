package models

//Project ...
type Project struct {
	Model
	Name    string `json:"name"`
	OwnerID int64  `json:"owner_id"`
}
