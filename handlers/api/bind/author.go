package bind

import (
	"costperfect/models"
	"net/http"
)

//Author ...
func Author(r *http.Request) (models.Author, map[string]bool, error) {
	var input models.Author
	var fields map[string]bool

	fields = make(map[string]bool)
	if ok := r.PostForm["username"]; ok != nil {
		input.Username = r.PostFormValue("username")
		fields["username"] = true
	}
	if ok := r.PostForm["password"]; ok != nil {
		input.Password = r.PostFormValue("password")
		fields["password"] = true
	}
	return input, fields, nil
}
