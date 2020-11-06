package models

//User ...
type User struct {
	Model
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
}
