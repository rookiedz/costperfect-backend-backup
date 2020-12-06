package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Job ...
type Job struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewJob ...
func NewJob() Job {
	var columns []string
	columns = []string{
		"job_id",
		"job_type_id",
		"job_group_id",
		"job_description",
	}
	return Job{TableName: "jobs", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (j Job) Create(job models.Job) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (job_type_id, job_group_id, job_description, job_created_at, job_updated_at)VALUES(?,?,?,?,?)`, j.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, job.TypeID, job.GroupID, job.Description, cds, cds)
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
func (j Job) Update(id int64, job models.Job) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET job_type_id = ?, job_group_id = ?, job_description = ?, job_updated_at = ? WHERE job_id = ?`, j.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, job.TypeID, job.GroupID, job.Description, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (j Job) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_id = ?`, j.TableName))
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
func (j Job) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idsString = ArrayInt64ToString(ids, ",")
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_id IN (%s)`, j.TableName, idsString))
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
func (j Job) FindByID(id int64) (models.Job, error) {
	var mJob models.Job
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE job_id = ?`, j.QueryColumn, j.TableName))
	if err != nil {
		return mJob, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mJob.ID, &mJob.TypeID, &mJob.GroupID, &mJob.Description); err != nil {
		if err == sql.ErrNoRows {
			return mJob, nil
		}
		return mJob, err
	}
	return mJob, nil
}

//FindAll ...
func (j Job) FindAll(setters ...Option) ([]models.Job, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50}
	for _, setter = range setters {
		setter(args)
	}
	var mJobs []models.Job
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, j.QueryColumn, j.TableName))
	if err != nil {
		return mJobs, err
	}
	defer stmt.Close()
	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mJobs, err
	}
	defer rows.Close()

	for rows.Next() {
		var mJob models.Job
		if err = rows.Scan(&mJob.ID, &mJob.TypeID, &mJob.GroupID, &mJob.Description); err != nil {
			return mJobs, err
		}
		mJobs = append(mJobs, mJob)
	}
	if err = rows.Close(); err != nil {
		return mJobs, err
	}
	if err = rows.Err(); err != nil {
		return mJobs, err
	}
	return mJobs, nil
}

//GetTotal ...
func (j Job) GetTotal() (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(job_id) FROM %s`, j.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
