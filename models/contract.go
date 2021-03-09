package models

//Contract ...
type Contract struct {
	Model
	AdvancePaymentMethod      string   `json:"advance_payment_method"`
	AdvancePaymentTotal       float64  `json:"advance_payment_total,string"`
	AdvancePaymentValue       float64  `json:"advance_payment_value,string"`
	BeginDate                 JSONDate `json:"begin_date"`
	BondNo                    string   `json:"bond_no"`
	BondBank                  string   `json:"bond_bank"`
	BondDate                  JSONDate `json:"bond_date"`
	BondValue                 float64  `json:"bond_value,string"`
	ContractNo                string   `json:"contract_no"`
	ContractorID              int64    `json:"contractor_id"`
	DeductMethod              string   `json:"deduct_method"`
	DeductPercentage          string   `json:"deduct_percentage"`
	DeliverDate               JSONDate `json:"deliver_date"`
	EmployerID                int64    `json:"employer_id"`
	EndDate                   JSONDate `json:"end_date"`
	JobID                     int64    `json:"job_id"`
	LOINo                     string   `json:"loi_no"`
	Name                      string   `json:"name"`
	NetValue                  string   `json:"net_value,string"`
	Note                      string   `json:"note"`
	PaymentInstallments       int64    `json:"payment_installments"`
	PaymentLastInstallments   int64    `json:"payment_last_installments"`
	PaymentMethod             string   `json:"payment_method"`
	PaymentPeriod             int64    `json:"payment_period"`
	PerformanceBondPercentage int64    `json:"performance_bond_percentage"`
	ProjectID                 int64    `json:"project_id"`
	RetentionMoneyMethod      string   `json:"retention_money_method"`
	RetentionMoneyPercentage  float64  `json:"retention_money_percentage,string"`
	SigningDate               JSONDate `json:"singing_date"`
	Tax                       float64  `json:"tax,string"`
	TaxValue                  float64  `json:"tax_value,string"`
	Value                     float64  `json:"value,string"`
}

//ContractPointer ...
type ContractPointer struct {
	AdvancePaymentMethod      *string   `json:"advance_payment_method"`
	AdvancePaymentTotal       *float64  `json:"advance_payment_total,string"`
	AdvancePaymentValue       *float64  `json:"advance_payment_value,string"`
	BeginDate                 *JSONDate `json:"begin_date"`
	BondNo                    *string   `json:"bond_no"`
	BondBank                  *string   `json:"bond_bank"`
	BondDate                  *JSONDate `json:"bond_date"`
	BondValue                 *float64  `json:"bond_value,string"`
	ContractNo                *string   `json:"contract_no"`
	ContractorID              *int64    `json:"contractor_id"`
	DeductMethod              *string   `json:"deduct_method"`
	DeductPercentage          *string   `json:"deduct_percentage"`
	DeliverDate               *JSONDate `json:"deliver_date"`
	EmployerID                *int64    `json:"employer_id"`
	EndDate                   *JSONDate `json:"end_date"`
	JobID                     *int64    `json:"job_id"`
	LOINo                     *string   `json:"loi_no"`
	Name                      *string   `json:"name"`
	NetValue                  *string   `json:"net_value,string"`
	Note                      *string   `json:"note"`
	PaymentInstallments       *int64    `json:"payment_installments"`
	PaymentLastInstallments   *int64    `json:"payment_last_installments"`
	PaymentMethod             *string   `json:"payment_method"`
	PaymentPeriod             *int64    `json:"payment_period"`
	PerformanceBondPercentage *int64    `json:"performance_bond_percentage"`
	ProjectID                 *int64    `json:"project_id"`
	RetentionMoneyMethod      *string   `json:"retention_money_method"`
	RetentionMoneyPercentage  *float64  `json:"retention_money_percentage,string"`
	SigningDate               *JSONDate `json:"singing_date"`
	Tax                       *float64  `json:"tax,string"`
	TaxValue                  *float64  `json:"tax_value,string"`
	Value                     *float64  `json:"value,string"`
}

//Match ...
func (cp ContractPointer) Match(contract *Contract) {
	if cp.AdvancePaymentMethod != nil {
		contract.AdvancePaymentMethod = *cp.AdvancePaymentMethod
	}
	if cp.AdvancePaymentValue != nil {
		contract.AdvancePaymentValue = *cp.AdvancePaymentValue
	}
	if cp.AdvancePaymentTotal != nil {
		contract.AdvancePaymentTotal = *cp.AdvancePaymentTotal
	}
	if cp.BeginDate != nil {
		contract.BeginDate = *cp.BeginDate
	}
	if cp.BondNo != nil {
		contract.BondNo = *cp.BondNo
	}
	if cp.BondBank != nil {
		contract.BondBank = *cp.BondBank
	}
	if cp.BondDate != nil {
		contract.BondDate = *cp.BondDate
	}
	if cp.ContractNo != nil {
		contract.ContractNo = *cp.ContractNo
	}
	if cp.ContractorID != nil {
		contract.ContractorID = *cp.ContractorID
	}
	if cp.DeductMethod != nil {
		contract.DeductMethod = *cp.DeductMethod
	}
	if cp.DeductPercentage != nil {
		contract.DeductPercentage = *cp.DeductPercentage
	}
	if cp.DeliverDate != nil {
		contract.DeliverDate = *cp.DeliverDate
	}
	if cp.EmployerID != nil {
		contract.EmployerID = *cp.EmployerID
	}
	if cp.EndDate != nil {
		contract.EndDate = *cp.EndDate
	}
	if cp.JobID != nil {
		contract.JobID = *cp.JobID
	}
	if cp.LOINo != nil {
		contract.LOINo = *cp.LOINo
	}
	if cp.Name != nil {
		contract.Name = *cp.Name
	}
	if cp.NetValue != nil {
		contract.NetValue = *cp.NetValue
	}
	if cp.Note != nil {
		contract.Note = *cp.Note
	}
	if cp.PaymentInstallments != nil {
		contract.PaymentInstallments = *cp.PaymentInstallments
	}
	if cp.PaymentLastInstallments != nil {
		contract.PaymentLastInstallments = *cp.PaymentLastInstallments
	}
	if cp.PaymentMethod != nil {
		contract.PaymentMethod = *cp.PaymentMethod
	}
	if cp.PaymentPeriod != nil {
		contract.PaymentPeriod = *cp.PaymentPeriod
	}
	if cp.PerformanceBondPercentage != nil {
		contract.PerformanceBondPercentage = *cp.PerformanceBondPercentage
	}
	if cp.ProjectID != nil {
		contract.ProjectID = *cp.ProjectID
	}
	if cp.RetentionMoneyMethod != nil {
		contract.RetentionMoneyMethod = *cp.RetentionMoneyMethod
	}
	if cp.RetentionMoneyPercentage != nil {
		contract.RetentionMoneyPercentage = *cp.RetentionMoneyPercentage
	}
	if cp.SigningDate != nil {
		contract.SigningDate = *cp.SigningDate
	}
	if cp.Tax != nil {
		contract.Tax = *cp.Tax
	}
	if cp.TaxValue != nil {
		contract.TaxValue = *cp.TaxValue
	}
	if cp.Value != nil {
		contract.Value = *cp.TaxValue
	}
}
