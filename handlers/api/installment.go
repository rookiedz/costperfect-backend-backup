package api

import (
	"costperfect/backend/models"
	"costperfect/backend/stores/mariadb"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

//Installment ...
type Installment struct{}

//NewInstallment ...
func NewInstallment() Installment {
	return Installment{}
}

//Create ...
func (i Installment) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Installment
	var mdbInstallment mariadb.Installment
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
	mdbInstallment = mariadb.NewInstallment()
	lastID, err = mdbInstallment.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(res))
}

//Update ...
func (i Installment) Update(w http.ResponseWriter, r *http.Request) {
	var input models.PInstallment
	var mInstallment models.Installment
	var mdbInstallment mariadb.Installment
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
	mdbInstallment = mariadb.NewInstallment()
	mInstallment, err = mdbInstallment.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	input.Match(&mInstallment)
	if err = mdbInstallment.Update(id, mInstallment); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Delete ...
func (i Installment) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbInstallment mariadb.Installment

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbInstallment = mariadb.NewInstallment()
	if err = mdbInstallment.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Get ...
func (i Installment) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var mInstallment models.Installment
	var mdbInstallment mariadb.Installment
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbInstallment = mariadb.NewInstallment()
	mInstallment, err = mdbInstallment.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	if reflect.DeepEqual(mInstallment, models.Installment{}) {
		JSON(w, http.StatusOK, NotFound())
		return
	}
	JSON(w, http.StatusOK, Success(mInstallment))
}

//All ...
func (i Installment) All(w http.ResponseWriter, r *http.Request) {
	var total, offset, limit int64
	var mInstallments []models.Installment
	var mdbInstallment mariadb.Installment
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}
	mdbInstallment = mariadb.NewInstallment()
	mInstallments, err = mdbInstallment.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbInstallment.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mInstallments))
}
