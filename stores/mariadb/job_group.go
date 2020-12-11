package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//JobGroup ...
type JobGroup struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewJobGroup ...
func NewJobGroup() JobGroup {
	var columns []string
	columns = []string{
		"job_group_id",
		"job_type_id",
		"job_group_label",
	}
	return JobGroup{TableName: "job_groups", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (jg JobGroup) Create(jobGroup models.JobGroup) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var lastID int64
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (job_type_id, job_group_label, job_group_created_at, job_group_updated_at)VALUES(?, ?, ?, ?)`, jg.TableName))
	if err != nil {
		return lastID, err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, jobGroup.TypeID, jobGroup.Label, cds, cds)
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
func (jg JobGroup) Update(id int64, jobGroup models.JobGroup) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET job_type_id = ?, job_group_label = ?, job_group_updated_at = ? WHERE job_group_id = ?`, jg.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, jobGroup.TypeID, jobGroup.Label, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (jg JobGroup) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_group_id = ?`, jg.TableName))
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
func (jg JobGroup) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idsString = ArrayInt64ToString(ids, ",")
	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE job_group_id IN (%s)`, jg.TableName, idsString))
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
func (jg JobGroup) FindByID(id int64) (models.JobGroup, error) {
	var mJobGroup models.JobGroup
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE job_group_id = ?`, jg.QueryColumn, jg.TableName))
	if err != nil {
		return mJobGroup, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mJobGroup.ID, &mJobGroup.TypeID, &mJobGroup.Label); err != nil {
		if err == sql.ErrNoRows {
			return mJobGroup, nil
		}
		return mJobGroup, err
	}
	return mJobGroup, nil
}

//FindAll ...
func (jg JobGroup) FindAll(setters ...Option) ([]models.JobGroup, error) {
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
	var mJobGroups []models.JobGroup
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, jg.QueryColumn, jg.TableName))
	if err != nil {
		return mJobGroups, err
	}
	defer stmt.Close()
	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mJobGroups, err
	}
	defer rows.Close()

	for rows.Next() {
		var mJobGroup models.JobGroup
		if err = rows.Scan(&mJobGroup.ID, &mJobGroup.TypeID, &mJobGroup.Label); err != nil {
			return mJobGroups, err
		}
		mJobGroups = append(mJobGroups, mJobGroup)
	}
	if err = rows.Close(); err != nil {
		return mJobGroups, err
	}
	if err = rows.Err(); err != nil {
		return mJobGroups, err
	}
	return mJobGroups, nil
}

//GetTotal ...
func (jg JobGroup) GetTotal() (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(job_group_id) FROM %s`, jg.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

//FindByType ...
func (jg JobGroup) FindByType(typeID int64, setters ...Option) ([]models.JobGroup, error) {
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
	var mJobGroups []models.JobGroup
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE job_type_id = ? LIMIT ?, ?`, jg.QueryColumn, jg.TableName))
	if err != nil {
		return mJobGroups, err
	}
	defer stmt.Close()
	rows, err = stmt.QueryContext(ctx, typeID, args.Offset-1, args.Limit)
	if err != nil {
		return mJobGroups, err
	}
	defer rows.Close()

	for rows.Next() {
		var mJobGroup models.JobGroup
		if err = rows.Scan(&mJobGroup.ID, &mJobGroup.TypeID, &mJobGroup.Label); err != nil {
			return mJobGroups, err
		}
		mJobGroups = append(mJobGroups, mJobGroup)
	}
	if err = rows.Close(); err != nil {
		return mJobGroups, err
	}
	if err = rows.Err(); err != nil {
		return mJobGroups, err
	}
	return mJobGroups, nil
}

//GetTotalByType ...
func (jg JobGroup) GetTotalByType(typeID int64) (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var total int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT COUNT(job_group_id) FROM %s WHERE job_type_id = ?`, jg.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, typeID).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
