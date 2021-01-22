package mariadb

import (
	"costperfect/backend/models"
	"strings"
)

//Contract ...
type Contract struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewContract ...
func NewContract() Contract {
	var columns []string
	columns = []string{
		"contract_id",
		"project_id",
		"contractor_id",
		"employer_id",
		"contract_name",
		"contract_no",
		"contract_loi_no",
		"contract_value",
		"contract_tax",
		"contract_tax_value",
		"contract_net_value",
		"contract_signing_date",
		"contract_begin_date",
		"contract_end_date",
		"contract_delivery_date",
		"contract_warranty_days",
		"contract_payment_method",
		"contract_payment_percentage",
		"contract_payment_amout",
		"contract_payment_installments",
		"contract_payment_installmentItems",
		"contract_advance_payment_method",
		"contract_advance_payment_percentage",
		"contract_advance_payment_amout",
		"contract_advance_payment_installments",
		"contract_advance_payment_installment_items",
		"contract_deduct_method",
		"contract_deduct_percentage",
		"contract_warranty_method",
		"contract_warranty_percentage",
		"contract_performance_bond_percentage",
		"contract_retention_money_method",
		"contract_retention_money_percentage",
		"contract_note",
		"contract_attachments",
	}
	return Contract{TableName: "contracts", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (c Contract) Create(contract models.Contract) (int64, error) {
	return 0, nil
}

//Update ...
func (c Contract) Update(id int64, contract models.Contract) error {
	return nil
}

//Delete ...
func (c Contract) Delete(id int64) error {
	return nil
}

//DeleteByIDs ...
func (c Contract) DeleteByIDs(ids []int64) error {
	return nil
}

//DeleteByProject ...
func (c Contract) DeleteByProject(id int64) error {
	return nil
}

//DeleteByProjectIDs ...
func (c Contract) DeleteByProjectIDs(ids []int64) error {
	return nil
}

//FindByID ...
func (c Contract) FindByID(id int64) (models.Contract, error) {
	return models.Contract{}, nil
}

//FindAll ...
func (c Contract) FindAll(setters ...Option) ([]models.Contract, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50}
	for _, setter = range setters {
		setter(args)
	}
	return []models.Contract{}, nil
}

//FindAllByProject ...
func (c Contract) FindAllByProject(id int64, setters ...Option) ([]models.Contract, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50}
	for _, setter = range setters {
		setter(args)
	}
	return []models.Contract{}, nil
}
