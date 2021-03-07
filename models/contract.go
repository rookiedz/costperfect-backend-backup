package models

//Contract ...
type Contract struct {
	Model
	ProjectID                       int64     `json:"project_id"`
	ContractorID                    int64     `json:"contractor_id"`
	EmployerID                      int64     `json:"employer_id"`
	JobID                           int64     `json:"job_id"`
	Name                            string    `json:"name"`
	ContractNo                      string    `json:"contract_no"`
	LetterOfIntentNo                string    `json:"loi_no"`
	Value                           float64   `json:"value"`
	Tax                             float64   `json:"tax"`
	TaxValue                        float64   `json:"tax_value"`
	NetValue                        float64   `json:"net_value"`
	SigningDate                     JSONDate  `json:"signing_date"`
	BeginDate                       JSONDate  `json:"begin_date"`
	EndDate                         JSONDate  `json:"end_date"`
	DeliveryDate                    JSONDate  `json:"deliver_date"`
	WarrantyDays                    int64     `json:"warranty_days"`  //days จำนวนวัน
	PaymentMethod                   string    `json:"payment_method"` // WORKINGS, INSTALLMENT
	PaymentPeriod                   int64     `json:"payment_period"`
	PaymentInstallments             float64   `json:"payment_installments"`
	PaymentLastInstallments         float64   `json:"payment_last_installments"`
	AdvancePaymentMethod            string    `json:"advance_payment_method"` // WORKINGS, INSTALLMENT
	AdvancePaymentTotal             float64   `json:"advance_payment_total"`
	AdvancePaymentValue             float64   `json:"advance_payment_value"`
	AdvancePaymentPercentage        float64   `json:"advance_payment_percentage"`
	AdvancePaymentAmout             float64   `json:"advance_payment_amout"`
	AdvancePaymentInstallments      int64     `json:"advance_payment_installments"`
	AdvancePaymentInstallmentValues []float64 `json:"advance_payment_installment_values"`
	DeductMethod                    string    `json:"deduct_method"` //BEFORE, AFTER
	DeductValue                     float64   `json:"deduct_value"`
	DeductPercentage                float64   `json:"deduct_percentage"`
	WarrantyMethod                  string    `json:"warranty_method"` //BEFORE, AFTER
	WarrantyPercentage              float64   `json:"warranty_percentage"`
	PerformanceBondPercentage       float64   `json:"performance_bond_percentage"` //หลักประกันการปฏิบัติงานตามสัญญา
	RetentionMoneyMethod            string    `json:"retention_money_method"`      //BEFORE, AFTER เงินประกันผลงาน
	RetentionMoneyPercentage        float64   `json:"retention_money_percentage"`
	Note                            string    `json:"note"`
	Attachments                     []string  `json:"attachments"`
}

//PContract ...
type PContract struct {
	ProjectID                       *int64     `json:"project_id"`
	ContractorID                    *int64     `json:"contractor_id"`
	EmployerID                      *int64     `json:"employer_id"`
	JobID                           *int64     `json:"job_id"`
	Name                            *string    `json:"name"`
	ContractNo                      *string    `json:"contract_no"`
	LetterOfIntentNo                *string    `json:"loi_no"`
	Value                           *float64   `json:"value"`
	Tax                             *float64   `json:"tax"`
	TaxValue                        *float64   `json:"tax_value"`
	NetValue                        *float64   `json:"net_value"`
	SigningDate                     *JSONDate  `json:"signing_date"`
	BeginDate                       *JSONDate  `json:"begin_date"`
	EndDate                         *JSONDate  `json:"end_date"`
	DeliveryDate                    *JSONDate  `json:"delivery_date"`
	WarrantyDays                    *int64     `json:"warranty_days"`  //days จำนวนวัน
	PaymentMethod                   *string    `json:"payment_method"` // WORKINGS, INSTALLMENT
	PaymentPeriod                   *int64     `json:"payment_period"`
	AdvancePaymentMethod            *string    `json:"advance_payment_method"` // WORKINGS, INSTALLMENT
	AdvancePaymentPercentage        *float64   `json:"advance_payment_percentage"`
	AdvancePaymentAmout             *float64   `json:"advance_payment_amout"`
	AdvancePaymentInstallments      *int64     `json:"advance_payment_installments"`
	AdvancePaymentInstallmentValues []*float64 `json:"advance_payment_installment_values"`
	DeductMethod                    *string    `json:"deduct_method"` //BEFORE, AFTER
	DeductPercentage                *float64   `json:"deduct_percentage"`
	WarrantyMethod                  *string    `json:"warranty_method"` //BEFORE, AFTER
	WarrantyPercentage              *float64   `json:"warranty_percentage"`
	PerformanceBondPercentage       *float64   `json:"performance_bond_percentage"` //หลักประกันการปฏิบัติงานตามสัญญา
	RetentionMoneyMethod            *string    `json:"retention_money_method"`      //BEFORE, AFTER เงินประกันผลงาน
	RetentionMoneyPercentage        *float64   `json:"retention_money_percentage"`
	Note                            *string    `json:"note"`
	Attachments                     []*string  `json:"attachments"`
}

//Match ...
func (pc PContract) Match(contract *Contract) {
	if pc.ProjectID != nil {
		contract.ProjectID = *pc.ProjectID
	}
	if pc.ContractorID != nil {
		contract.ContractorID = *pc.ContractorID
	}
	if pc.EmployerID != nil {
		contract.EmployerID = *pc.EmployerID
	}
	if pc.JobID != nil {
		contract.JobID = *pc.JobID
	}
	if pc.Name != nil {
		contract.Name = *pc.Name
	}
	if pc.ContractNo != nil {
		contract.ContractNo = *pc.ContractNo
	}
	if pc.LetterOfIntentNo != nil {
		contract.LetterOfIntentNo = *pc.LetterOfIntentNo
	}
	if pc.Value != nil {
		contract.Value = *pc.Value
	}
	if pc.Tax != nil {
		contract.Tax = *pc.Tax
	}
	if pc.TaxValue != nil {
		contract.TaxValue = *pc.TaxValue
	}
	if pc.NetValue != nil {
		contract.NetValue = *pc.NetValue
	}
	if pc.SigningDate != nil {
		contract.SigningDate = *pc.SigningDate
	}
	if pc.BeginDate != nil {
		contract.BeginDate = *pc.BeginDate
	}
	if pc.EndDate != nil {
		contract.EndDate = *pc.EndDate
	}
	if pc.DeliveryDate != nil {
		contract.DeliveryDate = *pc.DeliveryDate
	}
	if pc.WarrantyDays != nil {
		contract.WarrantyDays = *pc.WarrantyDays
	}
	if pc.PaymentMethod != nil {
		contract.PaymentMethod = *pc.PaymentMethod
	}
	if pc.PaymentPeriod != nil {
		contract.PaymentPeriod = *pc.PaymentPeriod
	}
	if pc.AdvancePaymentMethod != nil {
		contract.AdvancePaymentMethod = *pc.PaymentMethod
	}
	if pc.AdvancePaymentPercentage != nil {
		contract.AdvancePaymentPercentage = *pc.AdvancePaymentPercentage
	}
	if pc.AdvancePaymentAmout != nil {
		contract.AdvancePaymentAmout = *pc.AdvancePaymentAmout
	}
	if pc.AdvancePaymentInstallments != nil {
		contract.AdvancePaymentInstallments = *pc.AdvancePaymentInstallments
	}
	if pc.DeductMethod != nil {
		contract.DeductMethod = *pc.DeductMethod
	}
	if pc.DeductPercentage != nil {
		contract.DeductPercentage = *pc.DeductPercentage
	}
	if pc.WarrantyMethod != nil {
		contract.WarrantyMethod = *pc.WarrantyMethod
	}
	if pc.WarrantyPercentage != nil {
		contract.WarrantyPercentage = *pc.WarrantyPercentage
	}
	if pc.PerformanceBondPercentage != nil {
		contract.PerformanceBondPercentage = *pc.PerformanceBondPercentage
	}
	if pc.RetentionMoneyMethod != nil {
		contract.RetentionMoneyMethod = *pc.RetentionMoneyMethod
	}
	if pc.RetentionMoneyPercentage != nil {
		contract.RetentionMoneyPercentage = *pc.RetentionMoneyPercentage
	}
	if pc.Note != nil {
		contract.Note = *pc.Note
	}
}
