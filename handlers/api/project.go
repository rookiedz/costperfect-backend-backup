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

//Project ...
type Project struct {
	Endpoint string
}

//NewProject ...
func NewProject() Project {
	return Project{Endpoint: "projects"}
}

//Create ...
func (p Project) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Project
	var mdbProject mariadb.Project
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(p.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(p.Endpoint, err))
		return
	}

	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(p.Endpoint, err))
			return
		}
	}
	mdbProject = mariadb.NewProject()
	lastID, err = mdbProject.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(p.Endpoint, res))
}

//Update ...
func (p Project) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Project
	var mProject models.Project
	var mdbProject mariadb.Project
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(p.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(p.Endpoint, err))
			return
		}
	}
	mdbProject = mariadb.NewProject()
	mProject, err = mdbProject.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	input.Match(&mProject)
	if err = mdbProject.Update(id, mProject); err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(p.Endpoint, NewEmptyData()))
}

//Delete ...
func (p Project) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbProject mariadb.Project

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(p.Endpoint, err))
		return
	}
	mdbProject = mariadb.NewProject()
	if err = mdbProject.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(p.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (p Project) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbProject mariadb.Project

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	mdbProject = mariadb.NewProject()
	if err = mdbProject.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(p.Endpoint, NewEmptyData()))
}

//Get ...
func (p Project) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mProject models.Project
	var mdbProject mariadb.Project

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(p.Endpoint, err))
		return
	}
	mdbProject = mariadb.NewProject()
	mProject, err = mdbProject.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	if mProject == (models.Project{}) {
		JSON(w, http.StatusOK, NotFound(p.Endpoint))
		return
	}
	JSON(w, http.StatusOK, Success(p.Endpoint, mProject))
}

//All ...
func (p Project) All(w http.ResponseWriter, r *http.Request) {
	var offset, limit, total int64
	var mProjects []models.Project
	var mdbProject mariadb.Project
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	mdbProject = mariadb.NewProject()
	mProjects, err = mdbProject.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(p.Endpoint, err))
		return
	}
	total, err = mdbProject.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(p.Endpoint, err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(p.Endpoint, total, mProjects))
}
