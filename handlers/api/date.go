package api

import (
	"costperfect/backend/models"
	"encoding/json"
	"io"
	"net/http"
)

//DecodeDate ...
type DecodeDate struct {
	Created models.JSONDate `json:"created"`
}

//Date ...
type Date struct {
}

//Get ...
func (d Date) Get(w http.ResponseWriter, r *http.Request) {
	var input DecodeDate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(err))
			return
		}
		JSON(w, http.StatusOK, Failure(err))
		return
	}
}
