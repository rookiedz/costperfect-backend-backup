package mariadb

import (
	"context"
	"costperfect/models"
	"costperfect/utility/convert"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

//User ...
type User struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewUser ...
func NewUser() User {
	var columns []string
	columns = []string{
		"user_id",
		"user_employee_id",
		"user_name",
		"user_address",
		"user_telephone",
		"user_deleted",
		"user_created_at",
		"user_updated_at",
		"user_deleted_at",
	}
	return User{TableName: "users", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (u User) Create(user models.User) (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (user_employee_id, user_name, user_address, user_telephone, user_created_at, user_updated_at)VALUES(?, ?, ?, ?, ?, ?)`, u.TableName))
	if err != nil {
		return 0, err
	}
	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, user.EmployeeID, user.Name, user.Address, user.Telephone, cds, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

//Update ...
func (u User) Update(id int64, user models.User) error {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var cds string
	var no int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET user_employee_id = ?, user_name = ?, user_address = ?, user_telephone = ?, user_updated_at = ? WHERE user_deleted = 0 AND user_id = ?`, u.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, user.EmployeeID, user.Name, user.Address, user.Telephone, cds, id)
	if err != nil {
		return err
	}
	no, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if no < 1 {
		return errors.New(`Can't update user`)
	}
	return nil
}

//Delete ...
func (u User) Delete(id int64) error {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var no int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE user_id = ?`, u.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	no, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if no < 1 {
		return errors.New(`Can't delete user`)
	}
	return nil
}

//DeleteByIDs ...
func (u User) DeleteByIDs(ids []int64) error {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var no int64
	var idsString string

	idsString = convert.ArrayInt64ToString(ids, ",")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE user_id IN (%s)`, u.TableName, idsString))
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	no, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if no < 1 {
		return errors.New(`Can't delete user`)
	}
	return nil
}

//FindAll ...
func (u User) FindAll(settings ...Option) ([]models.User, error) {
	var args *Options
	var setter Option
	args = &Options{Offset: 1, Limit: 50, Column: "user_id", Sort: "DESC"}
	for _, setter = range settings {
		setter(args)
	}
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows
	var mus []models.User

	mus = []models.User{}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE user_deleted = 0 ORDER BY %s %s LIMIT ?,?`, u.QueryColumn, u.TableName, args.Column, args.Sort))
	if err != nil {
		return mus, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mus, err
	}
	defer rows.Close()

	for rows.Next() {
		var mu models.User
		if err = rows.Scan(&mu.ID, &mu.EmployeeID, &mu.Name, &mu.Address, &mu.Telephone, &mu.Deleted, &mu.CreatedAt, &mu.UpdatedAt, &mu.DeletedAt); err != nil {
			return mus, err
		}
		mus = append(mus, mu)
	}
	if err = rows.Close(); err != nil {
		return mus, err
	}
	if err = rows.Err(); err != nil {
		return mus, err
	}
	return mus, nil
}

//FindByID ...
func (u User) FindByID(id int64) (models.User, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var mu models.User

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE user_deleted = 0 AND user_id = ?`, u.QueryColumn, u.TableName))
	if err != nil {
		return mu, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, id).Scan(&mu.ID, &mu.EmployeeID, &mu.Name, &mu.Address, &mu.Telephone, &mu.Deleted, &mu.CreatedAt, &mu.UpdatedAt, &mu.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return mu, nil
		}
		return mu, err
	}
	return mu, nil
}
