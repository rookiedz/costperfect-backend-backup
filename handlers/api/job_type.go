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

//JobType ...
type JobType struct {
	Endpoint string
}

//NewJobType ...
func NewJobType() JobType {
	return JobType{Endpoint: "job_types"}
}

//Create ...
func (jt JobType) Create(w http.ResponseWriter, r *http.Request) {
	var input models.JobType
	var mdbJobType mariadb.JobType
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(jt.Endpoint, err))
			return
		}
	}
	mdbJobType = mariadb.NewJobType()
	lastID, err = mdbJobType.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(jt.Endpoint, res))
}

//Update ...
func (jt JobType) Update(w http.ResponseWriter, r *http.Request) {
	var input input.JobType
	var mJobType models.JobType
	var mdbJobType mariadb.JobType
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
			return
		}
	}
	mdbJobType = mariadb.NewJobType()
	mJobType, err = mdbJobType.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	input.Match(&mJobType)
	if err = mdbJobType.Update(id, mJobType); err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(jt.Endpoint, NewEmptyData()))
}

//Delete ...
func (jt JobType) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbJobType mariadb.JobType

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
		return
	}
	mdbJobType = mariadb.NewJobType()
	if err = mdbJobType.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(jt.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (jt JobType) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbJobType mariadb.JobType

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	mdbJobType = mariadb.NewJobType()
	if err = mdbJobType.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(jt.Endpoint, NewEmptyData()))
}

//Get ...
func (jt JobType) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mJobType models.JobType
	var mdbJobType mariadb.JobType

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(jt.Endpoint, err))
		return
	}
	mdbJobType = mariadb.NewJobType()
	mJobType, err = mdbJobType.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	if mJobType == (models.JobType{}) {
		JSON(w, http.StatusOK, NotFound(jt.Endpoint))
		return
	}
	JSON(w, http.StatusOK, Success(jt.Endpoint, mJobType))
}

//All ...
func (jt JobType) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit, total int64
	var mJobTypes []models.JobType
	var mdbJobType mariadb.JobType
	var err error

	mdbJobType = mariadb.NewJobType()
	mJobTypes, err = mdbJobType.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(jt.Endpoint, err))
		return
	}
	total, err = mdbJobType.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(jt.Endpoint, err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(jt.Endpoint, total, mJobTypes))
}
