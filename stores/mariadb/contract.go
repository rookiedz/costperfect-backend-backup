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
		"contract_advance_payment_method",
		"contract_advance_payment_total",
		"contract_advance_payment_value",
		"contract_begin_date",
		"contract_bond_no",
		"contract_bond_bank",
		"contract_bond_date",
		"contract_bond_value",
		"contract_contract_no",
		"contract_contractor_id",
		"contract_deduct_method",
		"contract_deduct_percentage",
		"contract_deliver_date",
		"contract_employer_id",
		"contract_end_date",
		"contract_job_id",
		"contract_loi_no",
		"contract_name",
		"contract_net_value",
		"contract_note",
		"contract_payment_installments",
		"contract_payment_last_installments",
		"contract_payment_method",
		"contract_payment_period",
		"contract_performance_bond_percentage",
		"contract_project_id",
		"contract_retention_money_method",
		"contract_retention_money_percentage",
		"contract_singing_date",
		"contract_tax",
		"contract_tax_value",
		"contract_value"}
	return Contract{TableName: "contracts", Columns: columns, InsertColumn: strings.Join(columns[1:], ","), QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (c Contract) Create(contract models.Contract) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, contract_created_at, contract_updated_at)VALUES()`, c.TableName, c.InsertColumn))
	if err != nil {
		return 0, nil
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, contract.AdvancePaymentMethod, contract.AdvancePaymentTotal, contract.AdvancePaymentValue,
		contract.BeginDate, contract.BondNo, contract.BondDate, contract.BondValue, contract.ContractNo, contract.ContractorID,
		contract.DeductMethod, contract.DeductPercentage, contract.DeliverDate, contract.EmployerID, contract.EndDate,
		contract.JobID, contract.LOINo, contract.Name, contract.NetValue, contract.Note, contract.PaymentInstallments,
		contract.PaymentLastInstallments, contract.PaymentMethod, contract.PaymentPeriod, contract.PerformanceBondPercentage,
		contract.ProjectID, contract.RetentionMoneyMethod, contract.RetentionMoneyPercentage, contract.SigningDate,
		contract.Tax, contract.TaxValue, contract.Value, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
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
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET contract_advance_payment_method = ?,
	contract_advance_payment_total = ?, contract_advance_payment_value = ?, contract_begin_date = ?,
	contract_bond_no = ?, contract_bond_bank = ?, contract_bond_date = ?, contract_bond_value = ?,
	contract_contract_no = ?, contract_contractor_id = ?, contract_deduct_method = ?, contract_deduct_percentage = ?,
	contract_deliver_date = ?, contract_employer_id = ?, contract_end_date = ?, contract_job_id = ?,
	contract_loi_no = ?, contract_name = ?, contract_net_value = ?, contract_note = ?,
	contract_payment_installments = ?, contract_payment_last_installments = ?, contract_payment_method = ?,
	contract_payment_period = ?, contract_performance_bond_percentage = ?, contract_project_id = ?,
	contract_retention_money_method = ?, contract_retention_money_percentage = ?, contract_singing_date = ?,
	contract_tax = ?, contract_tax_value = ?, contract_value = ?, contract_updated_at = ? 
	WHERE contract_id = ?`, c.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, contract.AdvancePaymentMethod, contract.AdvancePaymentTotal, contract.AdvancePaymentValue,
		contract.BeginDate, contract.BondNo, contract.BondBank, contract.BondDate, contract.BondValue, contract.ContractNo,
		contract.ContractorID, contract.DeductMethod, contract.DeductPercentage, contract.DeliverDate, contract.EmployerID,
		contract.EndDate, contract.JobID, contract.LOINo, contract.Name, contract.NetValue, contract.Note,
		contract.PaymentInstallments, contract.PaymentLastInstallments, contract.PaymentMethod, contract.PaymentPeriod,
		contract.PerformanceBondPercentage, contract.ProjectID, contract.RetentionMoneyMethod, contract.RetentionMoneyPercentage,
		contract.SigningDate, contract.Tax, contract.TaxValue, contract.Value, cds); err != nil {
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
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var mContract models.Contract

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE contract_id = ?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContract, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&mContract.ID, &mContract.AdvancePaymentMethod, &mContract.AdvancePaymentTotal, &mContract.AdvancePaymentValue,
		&mContract.BeginDate, &mContract.BondNo, &mContract.BondBank, &mContract.BondDate, &mContract.BondValue,
		&mContract.ContractNo, &mContract.ContractorID, &mContract.DeductMethod, &mContract.DeductPercentage, &mContract.DeliverDate,
		&mContract.EmployerID, &mContract.EndDate, &mContract.JobID, &mContract.LOINo, &mContract.Name, &mContract.NetValue,
		&mContract.Note, &mContract.PaymentInstallments, &mContract.PaymentLastInstallments, &mContract.PaymentMethod,
		&mContract.PaymentPeriod, &mContract.PerformanceBondPercentage, &mContract.SigningDate, &mContract.Tax,
		&mContract.TaxValue, &mContract.Value); err != nil {
		if err == sql.ErrNoRows {
			return mContract, nil
		}
		return mContract, err
	}
	return mContract, nil
}

//FindAll ...
func (c Contract) FindAll(setters ...Option) ([]models.Contract, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50}
	for _, setter = range setters {
		setter(args)
	}
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	var mContracts []models.Contract

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s ORDER BY contract_id DESC LIMIT ?, ?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContracts, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mContracts, err
	}
	defer rows.Close()

	for rows.Next() {
		var mContract models.Contract
		if err = rows.Scan(&mContract.ID, &mContract.AdvancePaymentMethod, &mContract.AdvancePaymentTotal, &mContract.AdvancePaymentValue,
			&mContract.BeginDate, &mContract.BondNo, &mContract.BondBank, &mContract.BondDate, &mContract.BondValue,
			&mContract.ContractNo, &mContract.ContractorID, &mContract.DeductMethod, &mContract.DeductPercentage,
			&mContract.DeliverDate, &mContract.EmployerID, &mContract.EndDate, &mContract.JobID, &mContract.LOINo, &mContract.Name,
			&mContract.NetValue, &mContract.Note, &mContract.PaymentInstallments, &mContract.PaymentLastInstallments,
			&mContract.PaymentMethod, &mContract.PaymentPeriod, &mContract.PerformanceBondPercentage, &mContract.ProjectID,
			&mContract.RetentionMoneyMethod, &mContract.RetentionMoneyPercentage, &mContract.SigningDate,
			&mContract.Tax, &mContract.TaxValue, &mContract.Value); err != nil {
			return mContracts, err
		}
		mContracts = append(mContracts, mContract)
	}
	if err = rows.Close(); err != nil {
		return mContracts, err
	}
	if err = rows.Err(); err != nil {
		return mContracts, err
	}
	return mContracts, nil
}

//GetTotal ...
func (c Contract) GetTotal() (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(contract_id) FROM %s`, c.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

//FindAllByProject ...
func (c Contract) FindAllByProject(id int64, setters ...Option) ([]models.Contract, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50}
	for _, setter = range setters {
		setter(args)
	}

	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows
	var mContracts []models.Contract
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE project_id = ? ORDER BY contract_id DESC LIMIT ?,?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContracts, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, id)
	if err != nil {
		return mContracts, err
	}
	defer rows.Close()
	for rows.Next() {
		var mContract models.Contract
		if err = rows.Scan(&mContract.ID, &mContract.AdvancePaymentMethod, &mContract.AdvancePaymentTotal, &mContract.AdvancePaymentValue,
			&mContract.BeginDate, &mContract.BondNo, &mContract.BondBank, &mContract.BondDate, &mContract.BondValue,
			&mContract.ContractNo, &mContract.ContractorID, &mContract.DeductMethod, &mContract.DeductPercentage,
			&mContract.DeliverDate, &mContract.EmployerID, &mContract.EndDate, &mContract.JobID, &mContract.LOINo, &mContract.Name,
			&mContract.NetValue, &mContract.Note, &mContract.PaymentInstallments, &mContract.PaymentLastInstallments,
			&mContract.PaymentMethod, &mContract.PaymentPeriod, &mContract.PerformanceBondPercentage, &mContract.ProjectID,
			&mContract.RetentionMoneyMethod, &mContract.RetentionMoneyPercentage, &mContract.SigningDate,
			&mContract.Tax, &mContract.TaxValue, &mContract.Value); err != nil {
			return mContracts, err
		}
		mContracts = append(mContracts, mContract)
	}
	if err = rows.Close(); err != nil {
		return mContracts, err
	}
	if err = rows.Err(); err != nil {
		return mContracts, err
	}
	return mContracts, nil
}

//GetTotalByProject ...
func (c Contract) GetTotalByProject(id int64) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(contract_id) FROM %s WHERE project_id = ?`, c.TableName))
	if err != nil {
		return total, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
