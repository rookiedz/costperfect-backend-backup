package input

import "costperfect/backend/models"

//User ...
type User struct {
	EmployeeID *string `json:"employee_id"`
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	Telephone  *string `json:"telephone"`
}

//Match ...
func (u User) Match(user *models.User) {
	if u.EmployeeID != nil {
		user.EmployeeID = *u.EmployeeID
	}
	if u.Name != nil {
		user.Name = *u.Name
	}
	if u.Address != nil {
		user.Address = *u.Address
	}
	if u.Telephone != nil {
		user.Telephone = *u.Telephone
	}
}
