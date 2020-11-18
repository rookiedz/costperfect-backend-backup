package api

import (
	"costperfect/backend/models"
	"costperfect/backend/stores/mariadb"
	"encoding/json"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

//CreateContractor ...
func CreateContractor(w http.ResponseWriter, r *http.Request) {
	var input models.Contractor
	var mdbContractor mariadb.Contractor
	var validate *validator.Validate
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
	}
	validate = validator.New()
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err("users", err))
			return
		}
	}
	mdbContractor = mariadb.NewContractor()
	lastID, err = mdbContractor.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success("users", res))
}
