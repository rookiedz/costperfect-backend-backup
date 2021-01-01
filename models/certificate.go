package models

import "time"

//Certificate ...
type Certificate struct {
	Model
	ContractID         int64                  `json:"contract_id"`
	VariationsValue    []CertificateVariation `json:"variations_value"`
	No                 string                 `json:"certificate_no"`
	DateOfValuation    time.Time              `json:"date_of_valuation"`
	WorkDoneValue      float64                `json:"work_done_value"`
	VariationWorkValue float64                `json:"variation_work_value"`
}
