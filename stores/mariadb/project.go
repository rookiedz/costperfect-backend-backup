package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Project ...
type Project struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewProject ...
func NewProject() Project {
	var columns []string
	columns = []string{
		"project_id",
		"project_name",
		"owner_id"}
	return Project{TableName: "projects", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (p Project) Create(project models.Project) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (project_name, owner_id, project_created_at, project_updated_at)VALUES(?,?,?,?)`, p.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, project.Name, project.OwnerID, cds, cds)
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
func (p Project) Update(id int64, project models.Project) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET project_name = ?, owner_id = ?, project_updated_at = ?  WHERE project_id = ?`, p.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, project.Name, project.OwnerID, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (p Project) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE project_id = ?`, p.TableName))
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
func (p Project) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	idsString = ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE project_id IN (%s)`, p.TableName, idsString))
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
func (p Project) FindByID(id int64) (models.Project, error) {
	var mProject models.Project
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE project_id = ?`, p.QueryColumn, p.TableName))
	if err != nil {
		return mProject, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mProject.Name, &mProject.OwnerID); err != nil {
		if err == sql.ErrNoRows {
			return mProject, nil
		}
		return mProject, err
	}
	return mProject, nil
}

//FindAll ...
func (p Project) FindAll(setters ...Option) ([]models.Project, error) {
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
	var mProjects []models.Project
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, p.QueryColumn, p.TableName))
	if err != nil {
		return mProjects, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mProjects, err
	}
	defer rows.Close()

	for rows.Next() {
		var mProject models.Project
		if err = rows.Scan(&mProject.Name, &mProject.OwnerID); err != nil {
			return mProjects, err
		}
		mProjects = append(mProjects, mProject)
	}
	if err = rows.Close(); err != nil {
		return mProjects, err
	}
	if err = rows.Err(); err != nil {
		return mProjects, err
	}
	return mProjects, nil
}