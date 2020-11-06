package mariadb

import (
	"context"
	"costperfect/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Author ...
type Author struct {
	TableName   string
	Columns     []string
	QueryColumn string
}

//NewAuthor ...
func NewAuthor() Author {
	var columns []string
	columns = []string{
		"author_id",
		"author_username",
		"author_password",
		"author_salt",
		"user_id",
		"author_deleted",
		"author_created_at",
		"author_updated_at",
		"author_deleted_at",
	}
	return Author{TableName: "authors", Columns: columns, QueryColumn: strings.Join(columns[:], ",")}
}

//Create ...
func (a Author) Create(auth models.Author) (int64, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result

	var lastID int64
	var cds string

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (author_username, author_password, author_salt, user_id, author_created_at, author_updated_at)VALUES(?, ?, ?, ?, ?)`, a.TableName))
	if err != nil {
		return lastID, err
	}
	defer stmt.Close()

	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, auth.Username, auth.Password, auth.Salt, auth.UserID, cds, cds)
	if err != nil {
		return lastID, err
	}
	return res.LastInsertId()
}

//Update ...
func (a Author) Update(id int64, auth models.Author) error {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var cds string
	var no int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET author_password = ?, author_salt = ?, author_updated_at = ? WHERE author_id = ? AND author_deleted = 0`, a.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds = CurrentDatetimeString()
	res, err = stmt.ExecContext(ctx, auth.Password, auth.Salt, cds, id)
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
func (a Author) Delete(id int64) error {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var res sql.Result
	var no int64

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE author_id = ?`, a.TableName))
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
		return errors.New(`Can't delete author`)
	}
	return nil
}

//FindByID ...
func (a Author) FindByID(id int64) (models.Author, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt

	var ma models.Author

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE author_id = ? AND author_deleted = 0`, a.QueryColumn, a.TableName))
	if err != nil {
		return ma, err
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(ctx, id).Scan(&ma.ID, &ma.Username, &ma.Password, &ma.Salt, &ma.UserID, &ma.Deleted, &ma.CreatedAt, &ma.UpdatedAt, &ma.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return ma, errors.New(`Author not found`)
		}
		return ma, err
	}
	return ma, nil
}

//FindByUsername ...
func (a Author) FindByUsername(username string) (models.Author, error) {
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt

	var ma models.Author

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE author_username = ? AND author_deleted = 0`, a.QueryColumn, a.TableName))
	if err != nil {
		return ma, err
	}
	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, username).Scan(&ma.ID, &ma.Username, &ma.Password, &ma.Salt, &ma.UserID, &ma.Deleted, &ma.CreatedAt, &ma.UpdatedAt, &ma.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return ma, errors.New(`Author not found`)
		}
		return ma, err
	}
	return ma, nil
}

//FindAll ...
func (a Author) FindAll(setters ...Option) ([]models.Author, error) {
	var args *Options
	args = &Options{Offset: 1, Limit: 50, Column: "author_id", Sort: "DESC"}
	for _, setter := range setters {
		setter(args)
	}
	var err error
	var ctx context.Context
	var cancel context.CancelFunc
	var stmt *sql.Stmt
	var rows *sql.Rows

	var mas []models.Author

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err = db.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE author_deleted = 0 ORDER BY %s %s LIMIT ?,?`, a.QueryColumn, a.TableName, args.Column, args.Sort))
	if err != nil {
		return mas, err
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, args.Offset-1, args.Limit)
	if err != nil {
		return mas, err
	}
	defer rows.Close()

	for rows.Next() {
		var ma models.Author
		if err = rows.Scan(&ma.ID, &ma.Username, &ma.Password, &ma.Salt, &ma.UserID, &ma.Deleted, &ma.CreatedAt, &ma.UpdatedAt, &ma.DeletedAt); err != nil {
			return mas, err
		}
		mas = append(mas, ma)
	}
	if err = rows.Close(); err != nil {
		return mas, err
	}
	if err = rows.Err(); err != nil {
		return mas, err
	}
	return []models.Author{}, nil
}
