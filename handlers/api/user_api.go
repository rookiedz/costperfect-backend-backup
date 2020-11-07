package api

import (
	"costperfect/handlers/api/match"
	"costperfect/models"
	"costperfect/stores/mariadb"
	"encoding/json"
	"log"
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
		log.Println(err.Error())
	}
	mdbUser = mariadb.NewUser()
	lastID, err = mdbUser.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, NewEmptyEntry("error", err.Error()))
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
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
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
	if mUser == (models.User{}) {
		JSON(w, http.StatusOK, NewEntry("success", "", NewNotFound("user")))
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
