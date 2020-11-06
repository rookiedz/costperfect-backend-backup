package models

//Author ...
type Author struct {
	Model
	Username string `json:"-"`
	Password string `json:"-"`
	Salt     string `json:"-"`
}
