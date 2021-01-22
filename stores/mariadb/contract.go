package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Contract ...
type Contract struct {
	TableName    string
	Columns      []string
	InsertColumn string
	QueryColumn  string
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
		"contract_advance_payment_method",
		"contract_advance_payment_percentage",
		"contract_advance_payment_amout",
		"contract_advance_payment_installments",
		"contract_deduct_method",
		"contract_deduct_percentage",
		"contract_warranty_method",
		"contract_warranty_percentage",
		"contract_performance_bond_percentage",
		"contract_retention_money_method",
		"contract_retention_money_percentage",
		"contract_note",
	}
	return Contract{TableName: "contracts", Columns: columns, InsertColumn: strings.Join(columns[1:], ","), QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (c Contract) Create(contract models.Contract) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, contract_created_at, contract_updated_at)VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, c.TableName, c.InsertColumn))
	if err != nil {
		return lastID, nil
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, contract.ProjectID, contract.ContractorID, contract.EmployerID,
		contract.Name, contract.ContractNo, contract.LetterOfIntentNo, contract.Value,
		contract.Tax, contract.TaxValue, contract.NetValue, contract.SigningDate,
		contract.BeginDate, contract.EndDate, contract.DeliveryDate, contract.WarrantyDays,
		contract.PaymentMethod, contract.PaymentPercentage, contract.PaymentAmout,
		contract.PaymentInstallments, contract.AdvancePaymentMethod, contract.AdvancePaymentPercentage,
		contract.AdvancePaymentAmout, contract.AdvancePaymentInstallments, contract.DeductMethod,
		contract.DeductPercentage, contract.WarrantyMethod, contract.WarrantyPercentage,
		contract.PerformanceBondPercentage, contract.RetentionMoneyMethod, contract.RetentionMoneyPercentage,
		contract.Note, cds, cds)
	if err != nil {
		return 0, err
	}
	lastID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

//Update ...
func (c Contract) Update(id int64, contract models.Contract) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET WHERE contract_id = ?`, c.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (c Contract) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE contract_id = ?`, c.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}

//DeleteByIDs ...
func (c Contract) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var idsString string
	var err error

	idsString = ArrayInt64ToString(ids, ",")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE contract_id IN (%s)`, c.TableName, idsString))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

//DeleteByProject ...
func (c Contract) DeleteByProject(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE project_id = ?`, c.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}

//DeleteByProjectIDs ...
func (c Contract) DeleteByProjectIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	idsString = ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE project_id IN (%s)`, c.TableName, idsString))
	if err != nil {
		return err
	}
	if _, err = stmt.ExecContext(ctx); err != nil {
		return err
	}
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
