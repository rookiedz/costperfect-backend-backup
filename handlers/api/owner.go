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

//Owner ...
type Owner struct {
	Endpoint string
}

//NewOwner ...
func NewOwner() Owner {
	return Owner{Endpoint: "owners"}
}

//Create ...
func (o Owner) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Owner
	var mdbOwner mariadb.Owner
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(o.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(o.Endpoint, err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(o.Endpoint, err))
			return
		}
	}
	mdbOwner = mariadb.NewOwner()
	lastID, err = mdbOwner.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(o.Endpoint, res))
}

//Update ...
func (o Owner) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Owner
	var mOwner models.Owner
	var mdbOwner mariadb.Owner
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(o.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(o.Endpoint, err))
			return
		}
	}
	mdbOwner = mariadb.NewOwner()
	mOwner, err = mdbOwner.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	input.Match(&mOwner)
	if err = mdbOwner.Update(id, mOwner); err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(o.Endpoint, NewEmptyData()))
}

//Delete ...
func (o Owner) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbOwner mariadb.Owner

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(o.Endpoint, err))
		return
	}
	mdbOwner = mariadb.NewOwner()
	if err = mdbOwner.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(o.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (o Owner) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbOwner mariadb.Owner

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	mdbOwner = mariadb.NewOwner()
	if err = mdbOwner.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(o.Endpoint, NewEmptyData()))
}

//Get ...
func (o Owner) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mOwner models.Owner
	var mdbOwner mariadb.Owner

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(o.Endpoint, err))
		return
	}
	mdbOwner = mariadb.NewOwner()
	mOwner, err = mdbOwner.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	if mOwner == (models.Owner{}) {
		JSON(w, http.StatusOK, NotFound(o.Endpoint))
		return
	}
	JSON(w, http.StatusOK, Success(o.Endpoint, mOwner))
}

//All ...
func (o Owner) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit int
	var mOwners []models.Owner
	var mdbOwner mariadb.Owner
	var err error

	mdbOwner = mariadb.NewOwner()
	mOwners, err = mdbOwner.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(o.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(o.Endpoint, mOwners))
}
