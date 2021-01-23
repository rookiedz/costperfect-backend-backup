package api

import (
	"costperfect/backend/handlers/api/input"
	"costperfect/backend/models"
	"costperfect/backend/stores/mariadb"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

//Employer ...
type Employer struct{}

//NewEmployer ...
func NewEmployer() Employer {
	return Employer{}
}

//Create ...
func (e Employer) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Employer
	var mdbEmployer mariadb.Employer
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(err))
			return
		}
		JSON(w, http.StatusOK, Failure(err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	mdbEmployer = mariadb.NewEmployer()
	lastID, err = mdbEmployer.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(res))
}

//Update ...
func (e Employer) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Employer
	var mEmployer models.Employer
	var mdbEmployer mariadb.Employer
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(err))
			return
		}
	}
	mdbEmployer = mariadb.NewEmployer()
	mEmployer, err = mdbEmployer.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	input.Match(&mEmployer)
	if err = mdbEmployer.Update(id, mEmployer); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Delete ...
func (e Employer) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbEmployer mariadb.Employer

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbEmployer = mariadb.NewEmployer()
	if err = mdbEmployer.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//DeleteByIDs ...
func (e Employer) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbEmployer mariadb.Employer

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	mdbEmployer = mariadb.NewEmployer()
	if err = mdbEmployer.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Get ...
func (e Employer) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mEmployer models.Employer
	var mdbEmployer mariadb.Employer

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbEmployer = mariadb.NewEmployer()
	mEmployer, err = mdbEmployer.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	if mEmployer == (models.Employer{}) {
		JSON(w, http.StatusOK, NotFound())
		return
	}
	JSON(w, http.StatusOK, Success(mEmployer))
}

//All ...
func (e Employer) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit, total int64
	var mEmployer []models.Employer
	var mdbEmployer mariadb.Employer
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	mdbEmployer = mariadb.NewEmployer()
	mEmployer, err = mdbEmployer.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbEmployer.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mEmployer))
}
