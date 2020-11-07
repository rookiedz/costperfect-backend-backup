package match

import "costperfect/models"

//User ...
func User(user *models.User, input models.User, fields map[string]bool) {
	if user.EmployeeID != input.EmployeeID {
		user.EmployeeID = input.EmployeeID
	}
	if user.Name != input.Name {
		user.Name = input.Name
	}
	if user.Address != input.Address {
		user.Address = input.Address
	}
	if user.Telephone != input.Telephone {
		user.Telephone = input.Telephone
	}
}
