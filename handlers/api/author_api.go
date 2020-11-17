package api

import (
	"costperfect/backend/handlers/api/bind"
	"costperfect/backend/handlers/api/match"
	"costperfect/backend/models"
	"costperfect/backend/stores/mariadb"
	"net/http"

	"github.com/go-chi/chi"
)

//CreateAuthor ...
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var input models.Author
	var mdbAuthor mariadb.Author
	var err error
	var lastID int64
	var res map[string]int64

	r.ParseForm()
	input, _, err = bind.Author(r)
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbAuthor = mariadb.NewAuthor()
	lastID, err = mdbAuthor.Create(input)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	res = make(map[string]int64)
	res["id"] = lastID
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}

//UpdateAuthor ...
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var input, mAuthor models.Author
	var mdbAuthor mariadb.Author
	var err error
	var id int64
	var fields map[string]bool

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	r.ParseForm()
	input, fields, err = bind.Author(r)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	mdbAuthor = mariadb.NewAuthor()
	mAuthor, err = mdbAuthor.FindByID(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	match.Author(&mAuthor, input, fields)
	if err = mdbAuthor.Update(id, mAuthor); err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	JSON(w, http.StatusNoContent, NewEmptyEntry("success", ""))
}

//DeleteAuthor ...
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbAuthor mariadb.Author

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbAuthor = mariadb.NewAuthor()
	if err = mdbAuthor.Delete(id); err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	JSON(w, http.StatusNoContent, NewEmptyEntry("success", ""))
}

//GetAuthor ...
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	var mAuthor models.Author
	var mdbAuthor mariadb.Author
	var res map[string]models.Author
	var id int64
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusBadRequest, NewEmptyEntry("fail", err.Error()))
		return
	}
	mdbAuthor = mariadb.NewAuthor()
	mAuthor, err = mdbAuthor.FindByID(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	if mAuthor == (models.Author{}) {
		JSON(w, http.StatusOK, NewEntry("success", "", NewNotFound("author")))
		return
	}
	res = make(map[string]models.Author)
	res["author"] = mAuthor
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}

//GetAuthors ...
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var mAuthors []models.Author
	var mdbAuthor mariadb.Author
	var res map[string][]models.Author
	var err error

	mdbAuthor = mariadb.NewAuthor()
	mAuthors, err = mdbAuthor.FindAll()
	if err != nil {
		JSON(w, http.StatusInternalServerError, NewEmptyEntry("error", err.Error()))
		return
	}
	res = make(map[string][]models.Author)
	res["authors"] = mAuthors
	JSON(w, http.StatusOK, NewEntry("success", "", res))
}
