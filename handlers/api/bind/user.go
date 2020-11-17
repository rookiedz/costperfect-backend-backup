package bind

import (
	"costperfect/backend/models"
	"net/http"
)

//User ...
func User(r *http.Request) (models.User, map[string]bool, error) {
	var input models.User
	var fields map[string]bool

	fields = make(map[string]bool)
	if ok := r.PostForm["employee_id"]; ok != nil {
		input.EmployeeID = r.PostFormValue("employee_id")
		fields["employee_id"] = true
	}
	if ok := r.PostForm["name"]; ok != nil {
		input.Name = r.PostFormValue("name")
		fields["name"] = true
	}
	if ok := r.PostForm["address"]; ok != nil {
		input.Address = r.PostFormValue("address")
		fields["address"] = true
	}
	if ok := r.PostForm["telephone"]; ok != nil {
		input.Telephone = r.PostFormValue("telephone")
		fields["telephone"] = true
	}
	return input, fields, nil
}
