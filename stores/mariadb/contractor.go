package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Contractor ...
type Contractor struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewContractor ...
func NewContractor() Contractor {
	var columns []string
	columns = []string{
		"contractor_id",
		"contractor_name",
		"contractor_name_eng",
		"contractor_acronym",
		"contractor_address",
		"contractor_telephone",
		"contractor_fax"}
	return Contractor{TableName: "contractors", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (c Contractor) Create(contractor models.Contractor) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (contractor_name, contractor_name_eng, contractor_acronym, contractor_address, contractor_telephone, contractor_fax, contractor_created_at, contractor_updated_at)VALUES(?,?,?,?,?,?,?,?)`, c.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, contractor.Name, contractor.NameEng, contractor.Acronym, contractor.Address, contractor.Telephone, contractor.Fax, cds, cds)
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
func (c Contractor) Update(id int64, contractor models.Contractor) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET contractor_name = ?, contractor_name_eng = ?, contractor_acronym = ?, contractor_address = ?, contractor_telephone = ?, contractor_fax = ?, contractor_updated_at = ? WHERE contractor_id = ?`, c.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, contractor.Name, contractor.NameEng, contractor.Acronym, contractor.Address, contractor.Telephone, contractor.Fax, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (c Contractor) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE contractor_id = ?`, c.TableName))
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
func (c Contractor) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	idsString = ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE contractor_id IN (%s)`, c.TableName, idsString))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

//FindByID ...
func (c Contractor) FindByID(id int64) (models.Contractor, error) {
	var mContractor models.Contractor
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE contractor_id = ?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContractor, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mContractor.ID, &mContractor.Name, &mContractor.NameEng, &mContractor.Acronym, &mContractor.Address, &mContractor.Telephone, &mContractor.Fax); err != nil {
		if err == sql.ErrNoRows {
			return mContractor, nil
		}
		return mContractor, err
	}
	return mContractor, nil
}

//FindAll ...
func (c Contractor) FindAll(setters ...Option) ([]models.Contractor, error) {
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
	var mContractors []models.Contractor
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContractors, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mContractors, err
	}
	defer rows.Close()

	for rows.Next() {
		var mContractor models.Contractor
		if err = rows.Scan(&mContractor.ID, &mContractor.Name, &mContractor.NameEng, &mContractor.Acronym, &mContractor.Address, &mContractor.Telephone, &mContractor.Fax); err != nil {
			return mContractors, err
		}
		mContractors = append(mContractors, mContractor)
	}
	if err = rows.Close(); err != nil {
		return mContractors, err
	}
	if err = rows.Err(); err != nil {
		return mContractors, err
	}
	return mContractors, nil
}

//GetTotal ...
func (c Contractor) GetTotal() (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(contractor_id) FROM %s`, c.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

//FindByProject ...
func (c Contractor) FindByProject(projectID int64, setters ...Option) ([]models.Contractor, error) {
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
	var mContractors []models.Contractor
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE project_id = ? LIMIT ?, ?`, c.QueryColumn, c.TableName))
	if err != nil {
		return mContractors, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, projectID, args.Offset-1, args.Limit)
	if err != nil {
		return mContractors, err
	}
	defer rows.Close()

	for rows.Next() {
		var mContractor models.Contractor
		if err = rows.Scan(&mContractor.ID, &mContractor.Name, &mContractor.NameEng, &mContractor.Acronym, &mContractor.Address, &mContractor.Telephone, &mContractor.Fax); err != nil {
			return mContractors, err
		}
		mContractors = append(mContractors, mContractor)
	}
	if err = rows.Close(); err != nil {
		return mContractors, err
	}
	if err = rows.Err(); err != nil {
		return mContractors, err
	}
	return mContractors, nil
}

//GetTotalByProject ...
func (c Contractor) GetTotalByProject(projectID int64) (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(contractor_id) FROM %s WHERE project_id = ?`, c.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, projectID).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
