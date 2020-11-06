package models

//Author ...
type Author struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	UserID   int64  `json:"user_id"`
}
