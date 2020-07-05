package sqlitehandler

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/interfaces/datastore"
	"IkezawaYuki/craft/logger"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var db *sql.DB

const (
	tableNameSignalEvents = "signal_events"
)

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

type sqliteHandler struct {
	conn *sql.DB
}

func NewSQLiteHandler(db *sql.DB) datastore.SQLHandler {
	return &sqliteHandler{conn: db}
}

func Connect() *sql.DB {
	var err error
	db, err = sql.Open(config.ConfigList.SQLDriver, config.ConfigList.DbName)
	if err != nil {
		panic(err)
	}
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT)`, tableNameSignalEvents)
	_, err = db.Exec(cmd)
	if err != nil {
		panic(err)
	}

	for _, duration := range config.ConfigList.Durations {
		tableName := GetCandleTableName(config.ConfigList.ProductCode, duration)
		fmt.Println(tableName)
		c := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low FLOAT,
				volume FLOAT)`, tableName)
		_, err = db.Exec(c)
		if err != nil {
			panic(err)
		}
	}
	logger.Info("database is initialize")
	return db
}

func (h sqliteHandler) Exec(query string, args ...interface{}) (datastore.Result, error) {
	res := sqlResult{}
	logger.Info("Exec method is invoked")
	result, err := h.conn.Exec(query, args...)
	if err != nil {
		return res, err
	}
	resultWrap := new(sqlResult)
	resultWrap.result = result
	return resultWrap, nil
}

func (h sqliteHandler) Query(query string, args ...interface{}) (datastore.Rows, error) {
	logger.Info("Query is invoked")
	rows, err := h.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	rowsWrap := new(sqlRows)
	rowsWrap.rows = rows
	return rowsWrap, nil
}

func (h sqliteHandler) QueryRow(query string, args ...interface{}) datastore.Row {
	logger.Info("QueryRow is invoked")
	row := h.conn.QueryRow(query, args...)
	rowWrap := new(sqlRow)
	rowWrap.row = row
	return rowWrap
}

func (h sqliteHandler) Begin() (datastore.Tx, error) {
	tx, err := h.conn.Begin()
	if err != nil {
		return nil, err
	}
	txWrap := new(sqlTx)
	txWrap.tx = tx
	return txWrap, nil
}
