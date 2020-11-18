package api

import (
	"costperfect/backend/models"
	"costperfect/backend/stores/mariadb"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

//CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var mdbUser mariadb.User
	var err error
	var lastID int64
	var res map[string]int64

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	mdbUser = mariadb.NewUser()
	lastID, err = mdbUser.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success("users", res))
}

//UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var mdbUser mariadb.User
	var err error
	var id int64

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	if err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	mdbUser = mariadb.NewUser()
	_, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	JSON(w, http.StatusOK, Success("users", NewEmptyData()))
}

//DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbUser mariadb.User

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	mdbUser = mariadb.NewUser()
	if err = mdbUser.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	JSON(w, http.StatusOK, Success("users", NewEmptyData()))
}

//DeleteUsers ...
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	var ids models.IDs
	var err error
	var mdbUser mariadb.User

	if err = json.NewDecoder(r.Body).Decode(&ids); err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	mdbUser = mariadb.NewUser()
	if err = mdbUser.DeleteByIDs(ids.IDs); err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	JSON(w, http.StatusOK, Success("users", NewEmptyData()))
}

//GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	var mUser models.User
	var mdbUser mariadb.User
	var id int64
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure("users", err))
		return
	}
	mdbUser = mariadb.NewUser()
	mUser, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	if mUser == (models.User{}) {
		JSON(w, http.StatusOK, NotFound("user"))
		return
	}
	JSON(w, http.StatusOK, Success("success", mUser))
}

//GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var mUsers []models.User
	var mdbUser mariadb.User
	var err error

	mdbUser = mariadb.NewUser()
	mUsers, err = mdbUser.FindAll()
	if err != nil {
		JSON(w, http.StatusOK, Err("users", err))
		return
	}
	JSON(w, http.StatusOK, Success("users", mUsers))
}
