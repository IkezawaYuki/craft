package infrastructure

import (
	"IkezawaYuki/craft/interfaces/datastore"
	"database/sql"
)

type sqlResult struct {
	result sql.Result
}

type sqlRows struct {
	rows *sql.Rows
}

type sqlRow struct {
	row *sql.Row
}

type sqlTx struct {
	tx *sql.Tx
}

type sqlStmt struct {
	stmt *sql.Stmt
}

func (r *sqlResult) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r *sqlResult) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

func (r *sqlRow) Scan(dest ...interface{}) error {
	return r.row.Scan()
}

func (r *sqlRows) Next() bool {
	return r.rows.Next()
}

func (r *sqlRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest)
}

func (r *sqlRows) Close() error {
	return r.rows.Close()
}

func (t *sqlTx) Commit() error {
	return t.tx.Commit()
}

func (t *sqlTx) Rollback() error {
	return t.tx.Rollback()
}

func (t *sqlTx) Prepare(query string) (datastore.Stmt, error) {
	stmt, err := t.tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	stmtWrap := new(sqlStmt)
	stmtWrap.stmt = stmt
	return stmtWrap, nil
}

func (t *sqlTx) Exec(query string, args ...interface{}) (datastore.Result, error) {
	result, err := t.tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	resultWrap := new(sqlResult)
	resultWrap.result = result
	return resultWrap, nil
}

func (t *sqlTx) Query(query string, args ...interface{}) (datastore.Rows, error) {
	rows, err := t.tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	rowsWrap := new(sqlRows)
	rowsWrap.rows = rows
	return rowsWrap, nil
}

func (t *sqlTx) QueryRow(query string, args ...interface{}) datastore.Row {
	row := t.tx.QueryRow(query, args...)
	rowWrap := new(sqlRow)
	rowWrap.row = row
	return rowWrap
}

func (s *sqlStmt) Exec(args ...interface{}) (datastore.Result, error) {
	result, err := s.stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	resultWrap := new(sqlResult)
	resultWrap.result = result
	return resultWrap, nil
}

func (s *sqlStmt) Query(args ...interface{}) (datastore.Rows, error) {
	rows, err := s.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	rowsWrap := new(sqlRows)
	rowsWrap.rows = rows
	return rowsWrap, nil
}

func (s *sqlStmt) QueryRow(args ...interface{}) datastore.Row {
	row := s.stmt.QueryRow(args...)
	rowWrap := new(sqlRow)
	rowWrap.row = row
	return rowWrap
}

func (s *sqlStmt) Close() error {
	return s.stmt.Close()
}
