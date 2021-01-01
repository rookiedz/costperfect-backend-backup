package models

import "time"

//Variation ...
type Variation struct {
	Model
	ContractID  int64     `json:"contract_id"`
	JobID       int64     `json:"job_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Value       int64     `json:"value"`
	RefRFANo    string    `json:"ref_rfa_no"`
	RefVENo     string    `json:"ref_ve_no"`
	RefPONo     string    `json:"ref_po_no"`
	RefSLNo     string    `json:"ref_sl_no"`
}
