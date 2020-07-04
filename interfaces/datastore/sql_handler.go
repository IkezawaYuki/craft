package datastore

type SQLHandler interface {
	Exec(query string, args ...interface{}) (Result, error)
	QueryRow(query string, args ...interface{}) Row
	Query(query string, args ...interface{}) (Rows, error)
	Begin() (Tx, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Tx interface {
	Commit() error
	Rollback() error
	Prepare(query string) (Stmt, error)
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
}

type Stmt interface {
	Exec(args ...interface{}) (Result, error)
	Query(args ...interface{}) (Rows, error)
	QueryRow(args ...interface{}) Row
	Close() error
}
