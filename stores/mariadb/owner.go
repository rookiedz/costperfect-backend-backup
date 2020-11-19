package mariadb

import (
	"context"
	"costperfect/backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Owner ...
type Owner struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewOwner ...
func NewOwner() Owner {
	var columns []string
	columns = []string{
		"owner_id",
		"owner_name",
		"owner_name_eng",
		"owner_director",
	}
	return Owner{TableName: "owners", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (o Owner) Create(owner models.Owner) (int64, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (owner_name, owner_name_eng, owner_director, owner_created_at, owner_updated_at)VALUES(?,?,?,?,?)`, o.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, owner.Name, owner.NameEng, owner.Director, cds, cds)
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
func (o Owner) Update(id int64, owner models.Owner) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET owner_name = ?, owner_name_eng = ?, owner_director = ?, owner_updated_at = ? WHERE owner_id = ?`, o.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	if _, err = stmt.ExecContext(ctx, owner.Name, owner.NameEng, owner.Director, cds, id); err != nil {
		return err
	}
	return nil
}

//Delete ...
func (o Owner) Delete(id int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE owner_id = ?`, o.TableName))
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
func (o Owner) DeleteByIDs(ids []int64) error {
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error
	var idsString string

	idsString = ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE owner_id IN (%s)`, o.TableName, idsString))
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
func (o Owner) FindByID(id int64) (models.Owner, error) {
	var mOwner models.Owner
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE owner_id = ?`, o.QueryColumn, o.TableName))
	if err != nil {
		return mOwner, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mOwner.ID, &mOwner.Name, &mOwner.NameEng, &mOwner.Director); err != nil {
		if err == sql.ErrNoRows {
			return mOwner, nil
		}
		return mOwner, err
	}
	return mOwner, nil
}

//FindAll ...
func (o Owner) FindAll(setters ...Option) ([]models.Owner, error) {
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
	var mOwners []models.Owner
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s LIMIT ?, ?`, o.QueryColumn, o.TableName))
	if err != nil {
		return mOwners, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mOwners, err
	}
	defer rows.Close()

	for rows.Next() {
		var mOwner models.Owner
		if err = rows.Scan(&mOwner.ID, &mOwner.Name, &mOwner.NameEng, &mOwner.Director); err != nil {
			return mOwners, err
		}
		mOwners = append(mOwners, mOwner)
	}
	if err = rows.Close(); err != nil {
		return mOwners, err
	}
	if err = rows.Err(); err != nil {
		return mOwners, err
	}
	return mOwners, nil
}
