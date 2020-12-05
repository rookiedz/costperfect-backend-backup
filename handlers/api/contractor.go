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

//Contractor ...
type Contractor struct {
	Endpoint string
}

//NewContractor ...
func NewContractor() Contractor {
	return Contractor{Endpoint: "contractors"}
}

//Create ...
func (c Contractor) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Contractor
	var mdbContractor mariadb.Contractor
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(c.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(c.Endpoint, err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(c.Endpoint, err))
			return
		}
	}
	mdbContractor = mariadb.NewContractor()
	lastID, err = mdbContractor.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(c.Endpoint, res))
}

//Update ...
func (c Contractor) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Contractor
	var mContractor models.Contractor
	var mdbContractor mariadb.Contractor
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(c.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(c.Endpoint, err))
			return
		}
	}
	mdbContractor = mariadb.NewContractor()
	mContractor, err = mdbContractor.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	input.Match(&mContractor)
	if err = mdbContractor.Update(id, mContractor); err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(c.Endpoint, NewEmptyData()))
}

//Delete ...
func (c Contractor) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbContractor mariadb.Contractor

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(c.Endpoint, err))
		return
	}
	mdbContractor = mariadb.NewContractor()
	if err = mdbContractor.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(c.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (c Contractor) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbContractor mariadb.Contractor

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	mdbContractor = mariadb.NewContractor()
	if err = mdbContractor.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(c.Endpoint, NewEmptyData()))
}

//Get ...
func (c Contractor) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mContractor models.Contractor
	var mdbContractor mariadb.Contractor

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(c.Endpoint, err))
		return
	}
	mdbContractor = mariadb.NewContractor()
	mContractor, err = mdbContractor.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	if mContractor == (models.Contractor{}) {
		JSON(w, http.StatusOK, NotFound(c.Endpoint))
		return
	}
	JSON(w, http.StatusOK, Success(c.Endpoint, mContractor))
}

//All ...
func (c Contractor) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit int64
	var mContractors []models.Contractor
	var mdbContractor mariadb.Contractor
	var err error

	mdbContractor = mariadb.NewContractor()
	mContractors, err = mdbContractor.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(c.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(c.Endpoint, mContractors))
}
