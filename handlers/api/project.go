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
type Project struct{}

//NewProject ...
func NewProject() Project {
	return Project{}
}

//Create ...
func (p Project) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Project
	var mdbProject mariadb.Project
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
	mdbProject = mariadb.NewProject()
	lastID, err = mdbProject.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	//Create Employer by project
	mdbEmployer = mariadb.NewEmployer()
	for _, value := range input.Employers {
		var mEmployer models.Employer
		mEmployer = models.Employer{ProjectID: lastID, Fullname: value}
		mdbEmployer.Create(mEmployer)
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(res))
}

//Update ...
func (p Project) Update(w http.ResponseWriter, r *http.Request) {
	var input input.Project
	var mProject models.Project
	var mdbProject mariadb.Project
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
	mdbProject = mariadb.NewProject()
	mProject, err = mdbProject.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	input.Match(&mProject)
	if err = mdbProject.Update(id, mProject); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}

	mdbEmployer = mariadb.NewEmployer()
	if err = mdbEmployer.DeleteByProject(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	for _, value := range input.Employers {
		var mEmployer models.Employer
		mEmployer = models.Employer{ProjectID: id, Fullname: *value}
		mdbEmployer.Create(mEmployer)
	}

	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Delete ...
func (p Project) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbProject mariadb.Project
	var mdbEmployer mariadb.Employer

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbProject = mariadb.NewProject()
	if err = mdbProject.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	mdbEmployer = mariadb.NewEmployer()
	if err = mdbEmployer.DeleteByProject(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}

	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//DeleteByIDs ...
func (p Project) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbProject mariadb.Project
	var mdbEmployer mariadb.Employer

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	mdbProject = mariadb.NewProject()
	if err = mdbProject.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	mdbEmployer = mariadb.NewEmployer()
	if err = mdbEmployer.DeleteByProjectIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Get ...
func (p Project) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mProject models.Project
	var mdbProject mariadb.Project

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbProject = mariadb.NewProject()
	mProject, err = mdbProject.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(mProject))
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
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbProject.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mProjects))
}

//Contracts ...
func (p Project) Contracts(w http.ResponseWriter, r *http.Request) {
	var id, total, offset, limit int64
	var mContracts []models.Contract
	var mdbContract mariadb.Contract
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}

	mdbContract = mariadb.NewContract()
	mContracts, err = mdbContract.FindAllByProject(id, mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbContract.GetTotalByProject(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Total(total, mContracts))
}

//Contractors ...
func (p Project) Contractors(w http.ResponseWriter, r *http.Request) {
	var id, total, offset, limit int64
	var mContractors []models.Contractor
	var mdbContractor mariadb.Contractor
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	mdbContractor = mariadb.NewContractor()
	mContractors, err = mdbContractor.FindByProject(id, mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbContractor.GetTotalByProject(id)
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mContractors))
}

//Employers ...
func (p Project) Employers(w http.ResponseWriter, r *http.Request) {
	var id, offset, limit, total int64
	var mEmployer []models.Employer
	var mdbEmployer mariadb.Employer
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	mdbEmployer = mariadb.NewEmployer()
	mEmployer, err = mdbEmployer.FindByProject(id, mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbEmployer.GetTotalByProject(id)
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mEmployer))
}
