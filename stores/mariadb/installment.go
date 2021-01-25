package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Installment ...
type Installment struct {
	TableName    string
	Columns      []string
	QueryColumn  string
	InsertColumn string
}

//NewInstallment ...
func NewInstallment() Installment {
	var columns []string
	columns = []string{
		"installment_id",
		"installment_no",
		"installment_value",
		"installment_relations",
		"contract_id",
	}
	return Installment{TableName: "installments", Columns: columns, QueryColumn: strings.Join(columns[:], ","), InsertColumn: strings.Join(columns[1:], ",")}
}

//Create ...
func (i Installment) Create(installment models.Installment) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, installment_created_at, installment_updated_at)VALUES(?,?,?,?,?)`, i.TableName, i.InsertColumn))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, installment.ContractID, installment.No, installment.Value, installment.Relations, installment.ContractID, cds, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

//Update ...
func (i Installment) Update(id int64, installment models.Installment) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cds = CurrentDatetimeString()
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET installment_no = ?, installment_value = ?, installment_relations = ?, contract_id = ?, installment_updated_at = ? WHERE installment_id = ?`, i.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, installment.No, installment.Value, installment.Relations, installment.ContractID, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (i Installment) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE installment_id = ?`, i.TableName))
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
func (i Installment) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idsString = ArrayInt64ToString(ids, ",")
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE installment_id IN(%s)`, i.TableName, idsString))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

//DeleteByRelations ...
func (i Installment) DeleteByRelations(contractID int64, relations string) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE contract_id = ? AND relations = ?`, i.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, contractID, relations); err != nil {
		return err
	}
	return nil
}

//FindByID ...
func (i Installment) FindByID(id int64) (models.Installment, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var mInstallment models.Installment

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE installment_id = ?`, i.QueryColumn, i.TableName))
	if err != nil {
		return mInstallment, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&mInstallment.ID, &mInstallment.No, &mInstallment.Value, &mInstallment.Relations, &mInstallment.ContractID); err != nil {
		return mInstallment, err
	}
	return mInstallment, nil
}

//FindAll ...
func (i Installment) FindAll(setters ...Option) ([]models.Installment, error) {
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
	var mInstallments []models.Installment

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s ORDER BY installment_id DESC LIMIT ?,?`, i.QueryColumn, i.TableName))
	if err != nil {
		return mInstallments, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mInstallments, err
	}
	defer rows.Close()

	for rows.Next() {
		var mInstallment models.Installment
		if err = rows.Scan(&mInstallment.ID, &mInstallment.No, &mInstallment.Value, &mInstallment.Relations, &mInstallment.ContractID); err != nil {
			return mInstallments, err
		}
		mInstallments = append(mInstallments, mInstallment)
	}
	if err = rows.Close(); err != nil {
		return mInstallments, err
	}
	if err = rows.Err(); err != nil {
		return mInstallments, err
	}
	return mInstallments, nil
}

//GetTotal ...
func (i Installment) GetTotal() (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(installment_id) FROM %s`, i.TableName))
	if err != nil {
		return total, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

//FindByRelations ...
func (i Installment) FindByRelations(contractID int64, relations string) ([]models.Installment, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	var mInstallments []models.Installment

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE contract_id = ? AND installment_relations = ?`, i.QueryColumn, i.TableName))
	if err != nil {
		return mInstallments, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, contractID, relations)
	if err != nil {
		return mInstallments, err
	}
	defer rows.Close()

	for rows.Next() {
		var mInstallment models.Installment
		if err = rows.Scan(&mInstallment.ID, &mInstallment.No, &mInstallment.Value, &mInstallment.Relations, &mInstallment.ContractID); err != nil {
			return mInstallments, err
		}
		mInstallments = append(mInstallments, mInstallment)
	}
	if err = rows.Close(); err != nil {
		return mInstallments, err
	}
	if err = rows.Err(); err != nil {
		return mInstallments, err
	}
	return []models.Installment{}, nil
}

//GetTotalByRelations ...
func (i Installment) GetTotalByRelations(contractID int64, relations string) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(contract_id) FROM %s WHERE contract_id = ? AND installment_relations = ?`, i.TableName))
	if err != nil {
		return total, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, contractID, relations).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
