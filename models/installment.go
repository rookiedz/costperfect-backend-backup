package models

//Installment ...
type Installment struct {
	Model
	No         int64   `json:"no"`
	Value      float64 `json:"value"`
	ContractID int64   `json:"contract_id"`
	Relations  string  `json:"relations"` //Payment, AdvancePayment
}

//PInstallment ...
type PInstallment struct {
	No         *int64   `json:"no"`
	Value      *float64 `json:"value"`
	ContractID *int64   `json:"contract_id"`
	Relations  *string  `json:"relations"` //Payment, AdvancePayment
}

//Match ...
func (pi PInstallment) Match(installment *Installment) {
	if pi.No != nil {
		installment.No = *pi.No
	}
	if pi.Value != nil {
		installment.Value = *pi.Value
	}
	if pi.ContractID != nil {
		installment.ContractID = *pi.ContractID
	}
	if pi.Relations != nil {
		installment.Relations = *pi.Relations
	}
}
