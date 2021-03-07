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

//Contract ...
type Contract struct{}

//NewContract ...
func NewContract() Contract {
	return Contract{}
}

//Create ...
func (c Contract) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Contract
	var mdbContract mariadb.Contract
	//var mdbInstallment mariadb.Installment
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
	mdbContract = mariadb.NewContract()
	lastID, err = mdbContract.Create(input)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	//mdbInstallment = mariadb.NewInstallment()
	//Create Payment Installment by contract
	// for no, value := range input.PaymentInstallmentValues {
	// 	var mInstallment models.Installment
	// 	mInstallment = models.Installment{No: int64(no), Value: value, ContractID: lastID, Relations: "payment"}
	// 	mdbInstallment.Create(mInstallment)
	// }
	//Create Advance Payment Installment by contract
	// for no, value := range input.AdvancePaymentInstallmentValues {
	// 	var mInstallment models.Installment
	// 	mInstallment = models.Installment{No: int64(no), Value: value, ContractID: lastID, Relations: "advance"}
	// 	mdbInstallment.Create(mInstallment)
	// }
	res = make(map[string]int64)
	res["last_id"] = lastID
	JSON(w, http.StatusOK, Success(res))
}

//Update ...
func (c Contract) Update(w http.ResponseWriter, r *http.Request) {
	var input models.PContract
	var mContract models.Contract
	var mdbContract mariadb.Contract
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
	mdbContract = mariadb.NewContract()
	mContract, err = mdbContract.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	input.Match(&mContract)
	if err = mdbContract.Update(id, mContract); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Delete ...
func (c Contract) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var mdbContract mariadb.Contract

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbContract = mariadb.NewContract()
	if err = mdbContract.Delete(id); err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	JSON(w, http.StatusOK, Success(NewEmptyData()))
}

//Get ...
func (c Contract) Get(w http.ResponseWriter, r *http.Request) {
	var id int64
	var mContract models.Contract
	var mdbContract mariadb.Contract
	var err error

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	mdbContract = mariadb.NewContract()
	mContract, err = mdbContract.FindByID(id)
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	if reflect.DeepEqual(mContract, models.Contract{}) {
		JSON(w, http.StatusOK, NotFound())
		return
	}
	JSON(w, http.StatusOK, Success(mContract))
}

//All ...
func (c Contract) All(w http.ResponseWriter, r *http.Request) {
	var total, offset, limit int64
	var mContracts []models.Contract
	var mdbContract mariadb.Contract
	var err error

	offset, err = INT64(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 1
	}
	limit, err = INT64(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}
	mdbContract = mariadb.NewContract()
	mContracts, err = mdbContract.FindAll(mariadb.WithOffset(offset), mariadb.WithLimit(limit))
	if err != nil {
		JSON(w, http.StatusOK, Err(err))
		return
	}
	total, err = mdbContract.GetTotal()
	if err != nil {
		if err != nil {
			JSON(w, http.StatusOK, Err(err))
			return
		}
	}
	JSON(w, http.StatusOK, Total(total, mContracts))
}

//Installments ...
func (c Contract) Installments(w http.ResponseWriter, r *http.Request) {
	var id int64
	var mInstallments []models.Installment
	var mdbInstallment mariadb.Installment
	var err error
	var relations string

	id, err = ID64(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		JSON(w, http.StatusOK, Failure(err))
		return
	}
	relations = chi.URLParamFromCtx(r.Context(), "relations")
	mdbInstallment = mariadb.NewInstallment()
	mInstallments, err = mdbInstallment.FindByRelations(id, relations)
	if err != nil {
		JSON(w, http.StatusOK, mInstallments)
		return
	}
	JSON(w, http.StatusOK, Success(mInstallments))
}
