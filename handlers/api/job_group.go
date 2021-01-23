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

//JobGroup ...
type JobGroup struct{}

//NewJobGroup ...
func NewJobGroup() JobGroup {
	return JobGroup{}
}

//Create ...
func (jg JobGroup) Create(w http.ResponseWriter, r *http.Request) {
	var input models.JobGroup
	var mdbJobGroup mariadb.JobGroup
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
	mdbJobGroup = mariadb.NewJobGroup()
	lastID, err = mdbJobGroup.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(res))
}

//Update ...
func (jg JobGroup) Update(w http.ResponseWriter, r *http.Request) {
	var input input.JobGroup
	var mJobGroup models.JobGroup
	var mdbJobGroup mariadb.JobGroup
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
	mdbJobGroup = mariadb.NewJobGroup()
	mJobGroup, err = mdbJobGroup.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	input.Match(&mJobGroup)
	if err = mdbJobGroup.Update(id, mJobGroup); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Delete ...
func (jg JobGroup) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbJobGroup mariadb.JobGroup

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbJobGroup = mariadb.NewJobGroup()
	if err = mdbJobGroup.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//DeleteByIDs ...
func (jg JobGroup) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbJobGroup mariadb.JobGroup

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	mdbJobGroup = mariadb.NewJobGroup()
	if err = mdbJobGroup.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Get ...
func (jg JobGroup) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mJobGroup models.JobGroup
	var mdbJobGroup mariadb.JobGroup

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbJobGroup = mariadb.NewJobGroup()
	mJobGroup, err = mdbJobGroup.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	if mJobGroup == (models.JobGroup{}) {
		JSON(w, http.StatusOK, NotFound())
		return
	}
	JSON(w, http.StatusOK, Success(mJobGroup))
}

//All ...
func (jg JobGroup) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit, total int64
	var mJobGroups []models.JobGroup
	var mdbJobGroup mariadb.JobGroup
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	mdbJobGroup = mariadb.NewJobGroup()
	mJobGroups, err = mdbJobGroup.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbJobGroup.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mJobGroups))
}

//Jobs ...
func (jg JobGroup) Jobs(w http.ResponseWriter, r *http.Request) {
	var id, offset, limit, total int64
	var mJobs []models.Job
	var mdbJob mariadb.Job
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbJob = mariadb.NewJob()
	mJobs, err = mdbJob.FindByGroup(id, mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbJob.GetTotalByGroup(id)
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mJobs))
}
