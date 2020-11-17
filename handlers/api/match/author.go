package match

import "costperfect/backend/models"

//Author ...
func Author(author *models.Author, input models.Author, fields map[string]bool) {
	if fields["username"] && author.Username != input.Username {
		author.Username = input.Username
	}
	if fields["password"] && author.Password != input.Password {
		author.Password = input.Password
	}
}
