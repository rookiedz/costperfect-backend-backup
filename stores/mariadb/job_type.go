package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//JobType ...
type JobType struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewJobType ...
func NewJobType() JobType {
	var columns []string
	columns = []string{
		"job_type_id",
		"job_type_label",
	}
	return JobType{TableName: "job_type", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (jt JobType) Create(jobType models.JobType) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var lastID int64
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (job_type_label, job_type_created_at, job_type_updated_at)VALUES(?, ?, ?)`, jt.TableName))
	if err != nil {
		return lastID, err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, jobType.Label, cds, cds)
	if err != nil {
		return lastID, err
	}
	lastID, err = res.LastInsertId()
	if err != nil {
		return lastID, err
	}
	return lastID, nil
}

//Update ...
func (jt JobType) Update(id int64, jobType models.JobType) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET job_type_label = ?, job_type_updated_at = ? WHERE job_type_id = ?`, jt.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, jobType.Label, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (jt JobType) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_type_id = ?`, jt.TableName))
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
func (jt JobType) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idsString = ArrayInt64ToString(ids, ",")
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_type_id IN (%s)`, jt.TableName, idsString))
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
func (jt JobType) FindByID(id int64) (models.JobType, error) {
	var mJobType models.JobType
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE job_type_id = ?`, jt.QueryColumn, jt.TableName))
	if err != nil {
		return mJobType, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mJobType.ID, &mJobType.Label); err != nil {
		if err == sql.ErrNoRows {
			return mJobType, nil
		}
		return mJobType, err
	}
	return mJobType, nil
}

//FindAll ...
func (jt JobType) FindAll(setters ...Option) ([]models.JobType, error) {
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
	var mJobTypes []models.JobType
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, jt.QueryColumn, jt.TableName))
	if err != nil {
		return mJobTypes, err
	}
	defer stmt.Close()
	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mJobTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var mJobType models.JobType
		if err = rows.Scan(&mJobType.ID, &mJobType.Label); err != nil {
			return mJobTypes, err
		}
		mJobTypes = append(mJobTypes, mJobType)
	}
	if err = rows.Close(); err != nil {
		return mJobTypes, err
	}
	if err = rows.Err(); err != nil {
		return mJobTypes, err
	}
	return mJobTypes, nil
}
