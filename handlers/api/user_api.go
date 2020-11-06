package api

import (
	"costperfect/handlers/api/bind"
	"costperfect/handlers/api/match"
	"costperfect/models"
	"costperfect/stores/mariadb"
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

	r.ParseForm()
	input, _, err = bind.User(r)
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbUser = mariadb.NewUser()
	lastID, err = mdbUser.Create(input)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	res = make(map[string]int64)
	res["id"] = lastID
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}

//UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input, mUser models.User
	var mdbUser mariadb.User
	var err error
	var id int64
	var fields map[string]bool

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	r.ParseForm()
	input, fields, err = bind.User(r)
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbUser = mariadb.NewUser()
	mUser, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	match.User(&mUser, input, fields)
	if err = mdbUser.Update(id, mUser); err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	JSON(w, http.StatusNoContent, NewEmptyEntry("success", ""))
}

//DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbUser mariadb.User

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbUser = mariadb.NewUser()
	if err = mdbUser.Delete(id); err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	JSON(w, http.StatusNoContent, NewEmptyEntry("success", ""))
}

//GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	var mUser models.User
	var mdbUser mariadb.User
	var res map[string]models.User
	var id int64
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbUser = mariadb.NewUser()
	mUser, err = mdbUser.FindByID(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	res = make(map[string]models.User)
	res["user"] = mUser
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}

//GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var mUsers []models.User
	var mdbUser mariadb.User
	var res map[string][]models.User
	var err error

	mdbUser = mariadb.NewUser()
	mUsers, err = mdbUser.FindAll()
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	res = make(map[string][]models.User)
	res["users"] = mUsers
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}
