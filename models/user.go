package models

//User ...
type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
}
