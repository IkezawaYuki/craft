package infrastructure

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/interfaces/datastore"
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

type sqlHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() datastore.SQLHandler {
	return sqlHandler{Conn: Connect()}
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
	return db
}

type SqlResult struct {
	Result *sql.Result
}

type SqlRows struct {
	Rows *sql.Rows
}

type SqlRow struct {
	Row *sql.Row
}
