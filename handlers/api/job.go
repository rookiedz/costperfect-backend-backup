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

//Job ...
type Job struct {
	Endpoint string
}

//NewJob ...
func NewJob() Job {
	return Job{Endpoint: "jobs"}
}

//Create ...
func (j Job) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Job
	var mdbJob mariadb.Job
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(j.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(j.Endpoint, err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(j.Endpoint, err))
			return
		}
	}
	mdbJob = mariadb.NewJob()
	lastID, err = mdbJob.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(j.Endpoint, res))
}

//Update ...
func (j Job) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Job
	var mJob models.Job
	var mdbJob mariadb.Job
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(j.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(j.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(j.Endpoint, err))
			return
		}
	}
	mdbJob = mariadb.NewJob()
	mJob, err = mdbJob.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	input.Match(&mJob)
	if err = mdbJob.Update(id, mJob); err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(j.Endpoint, NewEmptyData()))
}

//Delete ...
func (j Job) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbJob mariadb.Job

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(j.Endpoint, err))
		return
	}
	mdbJob = mariadb.NewJob()
	if err = mdbJob.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(j.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (j Job) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbJob mariadb.Job

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	mdbJob = mariadb.NewJob()
	if err = mdbJob.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(j.Endpoint, NewEmptyData()))
}

//Get ...
func (j Job) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mJob models.Job
	var mdbJob mariadb.Job

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(j.Endpoint, err))
		return
	}
	mdbJob = mariadb.NewJob()
	mJob, err = mdbJob.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	if mJob == (models.Job{}) {
		JSON(w, http.StatusOK, NotFound(j.Endpoint))
		return
	}
	JSON(w, http.StatusOK, Success(j.Endpoint, mJob))
}

//All ...
func (j Job) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit int
	var mJobs []models.Job
	var mdbJob mariadb.Job
	var err error

	mdbJob = mariadb.NewJob()
	mJobs, err = mdbJob.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(j.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(j.Endpoint, mJobs))
}
