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

//User ...
type User struct {
	Endpoint string
}

//NewUser ...
func NewUser() User {
	return User{Endpoint: "users"}
}

//Create ...
func (u User) Create(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var mdbUser mariadb.User
	var ok bool
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		if err == io.EOF {
			JSON(w, http.StatusOK, Failure(u.Endpoint, err))
			return
		}
		JSON(w, http.StatusOK, Failure(u.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Err(u.Endpoint, err))
			return
		}
	}
	mdbUser = mariadb.NewUser()
	lastID, err = mdbUser.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(u.Endpoint, res))
}

//Update ...
func (u User) Update(w http.ResponseWriter, r *http.Request) {
	var input input.User
	var mUser models.User
	var mdbUser mariadb.User
	var err error
	var ok bool
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(u.Endpoint, err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure(u.Endpoint, err))
		return
	}
	if err = validate.Struct(input); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			JSON(w, http.StatusOK, Failure(u.Endpoint, err))
			return
		}
	}
	mdbUser = mariadb.NewUser()
	mUser, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	input.Match(&mUser)
	if err = mdbUser.Update(id, mUser); err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(u.Endpoint, NewEmptyData()))
}

//Delete ...
func (u User) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbUser mariadb.User

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(u.Endpoint, err))
		return
	}
	mdbUser = mariadb.NewUser()
	if err = mdbUser.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(u.Endpoint, NewEmptyData()))
}

//DeleteByIDs ...
func (u User) DeleteByIDs(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbUser mariadb.User

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	mdbUser = mariadb.NewUser()
	if err = mdbUser.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	JSON(w, http.StatusOK, Success(u.Endpoint, NewEmptyData()))
}

//Get ...
func (u User) Get(w http.ResponseWriter, r *http.Request) {
	var mUser models.User
	var mdbUser mariadb.User
	var id int64
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(u.Endpoint, err))
		return
	}
	mdbUser = mariadb.NewUser()
	mUser, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	if mUser == (models.User{}) {
		JSON(w, http.StatusOK, NotFound("user"))
		return
	}
	JSON(w, http.StatusOK, Success("success", mUser))
}

//All ...
func (u User) All(w http.ResponseWriter, r *http.Request) {
	var mUsers []models.User
	var mdbUser mariadb.User
	var err error
	var total, offset, limit int64

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}
	mdbUser = mariadb.NewUser()
	mUsers, err = mdbUser.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(u.Endpoint, err))
		return
	}
	total, err = mdbUser.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(u.Endpoint, err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mUsers))
}
