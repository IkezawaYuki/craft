package interfaces

import "database/sql"

type SQLiteHandler struct {
	conn *sql.Conn
}

func NewSQLiteHandler(conn *sql.Conn) *SQLiteHandler {
	return &SQLiteHandler{conn: conn}
}

//func (h *SQLiteHandler)
