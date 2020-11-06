package match

import "costperfect/models"

//User ...
func User(user *models.User, input models.User, fields map[string]bool) {
	if fields["employee_id"] && user.EmployeeID != input.EmployeeID {
		user.EmployeeID = input.EmployeeID
	}
	if fields["name"] && user.Name != input.Name {
		user.Name = input.Name
	}
	if fields["address"] && user.Address != input.Address {
		user.Address = input.Address
	}
	if fields["telephone"] && user.Telephone != input.Telephone {
		user.Telephone = input.Telephone
	}
}
