package models

import "time"

//Contract ...
type Contract struct {
	Model
	ProjectID                int64     `json:"project_id"`
	ContractorID             int64     `json:"contractor_id"`
	No                       string    `json:"no"`
	Name                     string    `json:"name"`
	Value                    int64     `json:"value"`
	Tax                      int64     `json:"tax"`
	CreatedData              time.Time `json:"created_date"`
	BeginDate                time.Time `json:"begin_date"`
	EndDate                  time.Time `json:"end_date"`
	DeliveryDate             time.Time `json:"delivery_date"`
	WarrantyPeriod           int64     `json:"warranty_period"`
	TermOfPayment            string    `json:"term_of_payment"`
	TermOfAdvancePayment     string    `json:"term_of_advance_payment"`
	AdvancePaymentPercentage int64     `json:"advance_payment_percentage"`
	DeductConfig             int64     `json:"deduct_config"`
	DeductPercentage         int64     `json:"deduct_percentage"`
	WarrantyConfig           int64     `json:"warranty_config"`
	WarrantyPercentage       int64     `json:"warranty_percentage"`
	CollateralPercentage     int64     `json:"collateral_percentage"`
	Note                     string    `json:"note"`
}
