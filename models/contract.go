package models

import "time"

//Contract ...
type Contract struct {
	Model
	ProjectID                      int64         `json:"project_id"`
	ContractorID                   int64         `json:"contractor_id"`
	EmployerID                     int64         `json:"employer_id"`
	Name                           string        `json:"name"`
	ContractNo                     string        `json:"contract_no"`
	LetterOfIntentNo               string        `json:"loi_no"`
	Value                          float64       `json:"value"`
	Tax                            int64         `json:"tax"`
	TaxValue                       float64       `json:"tax_value"`
	NetValue                       float64       `json:"net_value"`
	SigningDate                    time.Time     `json:"signing_date"`
	BeginDate                      time.Time     `json:"begin_date"`
	EndDate                        time.Time     `json:"end_date"`
	DeliveryDate                   time.Time     `json:"delivery_date"`
	WarrantyDays                   int64         `json:"warranty_days"`  //days จำนวนวัน
	PaymentMethod                  int64         `json:"payment_method"` // WORKINGS, INSTALLMENT
	PaymentPercentage              float64       `json:"payment_percentage"`
	PaymentAmout                   float64       `json:"payment_amout"`
	PaymentInstallments            int64         `json:"payment_installments"`
	PaymentInstallmentItems        []Installment `json:"payment_installmentItems"`
	AdvancePaymentMethod           int64         `json:"advance_payment_method"` // WORKINGS, INSTALLMENT
	AdvancePaymentPercentage       float64       `json:"advance_payment_percentage"`
	AdvancePaymentAmout            float64       `json:"advance_payment_amout"`
	AdvancePaymentInstallments     int64         `json:"advance_payment_installments"`
	AdvancePaymentInstallmentItems []Installment `json:"advance_payment_installment_items"`
	DeductMethod                   int64         `json:"deduct_method"` //BEFORE, AFTER
	DeductPercentage               float64       `json:"deduct_percentage"`
	WarrantyMethod                 int64         `json:"warranty_method"` //BEFORE, AFTER
	WarrantyPercentage             float64       `json:"warranty_percentage"`
	PerformanceBondPercentage      float64       `json:"performance_bond_percentage"` //หลักประกันการปฏิบัติงานตามสัญญา
	RetentionMoneyMethod           int64         `json:"retention_money_method"`      //BEFORE, AFTER เงินประกันผลงาน
	RetentionMoneyPercentage       string        `json:"retention_money_percentage"`
	Note                           string        `json:"note"`
	Attachments                    []string      `json:"attachments"`
}
