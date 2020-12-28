package models

import "time"

//Contract ...
type Contract struct {
	Model
	ProjectID                int64     `json:"project_id"`
	ContractorID             int64     `json:"contractor_id"`
	EmployerID               int64     `json:"employer_id"`
	ContractNo               string    `json:"contract_no"`
	LetterOfIntentNo         string    `json:"loi_no"`
	Description              string    `json:"description"`
	Value                    int64     `json:"value"`
	Tax                      int64     `json:"tax"`
	SigningDate              time.Time `json:"signing_date"`
	BeginDate                time.Time `json:"begin_date"`
	EndDate                  time.Time `json:"end_date"`
	DeliveryDate             time.Time `json:"delivery_date"`
	WarrantyDays             int64     `json:"warranty_days"`  //days จำนวนวัน
	PaymentMethod            string    `json:"payment_method"` //percent, period
	PaymentPercentage        int64     `json:"payment_percentage"`
	PaymentPeriods           []Period  `json:"payment_periods"`
	AdvancePaymentMethod     string    `json:"advance_payment_method"` //percent, period
	AdvancePaymentPercentage int64     `json:"advance_payment_percentage"`
	AdvancePaymentPeriods    []Period  `json:"advance_payment_periods"`
	DeductMethod             int64     `json:"deduct_method"`
	DeductPercentage         int64     `json:"deduct_percentage"`
	WarrantyMethod           int64     `json:"warranty_method"`
	WarrantyPercentage       int64     `json:"warranty_percentage"`
	CollateralPercentage     int64     `json:"collateral_percentage"`
	Note                     string    `json:"note"`
	Attachments              []string  `json:"attachments"`
}
