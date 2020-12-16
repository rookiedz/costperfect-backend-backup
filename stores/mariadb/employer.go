package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Employer ...
type Employer struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewEmployer ...
func NewEmployer() Employer {
	var columns []string
	columns = []string{
		"employer_id",
		"employer_fullname",
		"project_id"}
	return Employer{TableName: "employers", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (e Employer) Create(employer models.Employer) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (employer_fullname, project_id, employer_created_at, employer_updated_at)VALUES(?,?,?,?)`, e.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, employer.Fullname, employer.ProjectID, cds, cds)
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
func (e Employer) Update(id int64, employer models.Employer) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET employer_fullname = ?, project_id = ?, employer_updated_at = ?  WHERE employer_id = ?`, e.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, employer.Fullname, employer.ProjectID, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (e Employer) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE employer_id = ?`, e.TableName))
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
func (e Employer) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	idsString = ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE employer_id IN (%s)`, e.TableName, idsString))
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
func (e Employer) FindByID(id int64) (models.Employer, error) {
	var mEmployer models.Employer
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE employer_id = ?`, e.QueryColumn, e.TableName))
	if err != nil {
		return mEmployer, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mEmployer.ID, &mEmployer.Fullname, &mEmployer.ProjectID); err != nil {
		if err == sql.ErrNoRows {
			return mEmployer, nil
		}
		return mEmployer, err
	}
	return mEmployer, nil
}

//FindAll ...
func (e Employer) FindAll(setters ...Option) ([]models.Employer, error) {
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
	var mEmployers []models.Employer
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, e.QueryColumn, e.TableName))
	if err != nil {
		return mEmployers, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mEmployers, err
	}
	defer rows.Close()

	for rows.Next() {
		var mEmployer models.Employer
		if err = rows.Scan(&mEmployer.ID, &mEmployer.Fullname, &mEmployer.ProjectID); err != nil {
			return mEmployers, err
		}
		mEmployers = append(mEmployers, mEmployer)
	}
	if err = rows.Close(); err != nil {
		return mEmployers, err
	}
	if err = rows.Err(); err != nil {
		return mEmployers, err
	}
	return mEmployers, nil
}

//GetTotal ...
func (e Employer) GetTotal() (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(employer_id) FROM %s`, e.TableName))
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
func (e Employer) FindByProject(projectID int64, setters ...Option) ([]models.Employer, error) {
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
	var mEmployers []models.Employer
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE project_id = ? LIMIT ?, ?`, e.QueryColumn, e.TableName))
	if err != nil {
		return mEmployers, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, projectID, args.Offset-1, args.Limit)
	if err != nil {
		return mEmployers, err
	}
	defer rows.Close()

	for rows.Next() {
		var mEmployer models.Employer
		if err = rows.Scan(&mEmployer.ID, &mEmployer.Fullname, &mEmployer.ProjectID); err != nil {
			return mEmployers, err
		}
		mEmployers = append(mEmployers, mEmployer)
	}
	if err = rows.Close(); err != nil {
		return mEmployers, err
	}
	if err = rows.Err(); err != nil {
		return mEmployers, err
	}
	return mEmployers, nil
}

//GetTotalByProject ...
func (e Employer) GetTotalByProject(projectID int64) (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(employer_id) FROM %s WHERE project_id = ?`, e.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, projectID).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
