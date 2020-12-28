package models

//Period ...
type Period struct {
	Model
	No        int64  `json:"no"`
	Value     int64  `json:"value"`
	Relations string `json:"relations"`
}
