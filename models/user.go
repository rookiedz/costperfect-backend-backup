package models

//User ...
type User struct {
	Model
	EmployeeID string `json:"employee_id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
}
