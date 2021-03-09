package api

import (
	"costperfect/backend/models"
	"encoding/json"
	"io"
	"net/http"
)

//CreatedDate ...
type CreatedDate struct {
	Created models.JSONDate `json:"created"`
}

//Test ...
type Test struct{}

//NewTest ...
func NewTest() Test {
	return Test{}
}

//Create ...
func (t Test) Create(w http.ResponseWriter, r *http.Request) {
	var input CreatedDate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(err))
			return
		}
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	JSON(w, http.StatusOK, Success(input))
}

//Get ...
func (t Test) Get(w http.ResponseWriter, r *http.Request) {
	var input CreatedDate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(err))
			return
		}
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	JSON(w, http.StatusOK, Success(input))
}
