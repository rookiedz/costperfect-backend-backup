package models

//Installment ...
type Installment struct {
	Model
	No       int64  `json:"no"`
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}
