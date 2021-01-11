package models

//User ...
type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	EmployeeID string `json:"employee_id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
}
