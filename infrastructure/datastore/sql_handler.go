package datastore

import (
	"IkezawaYuki/craft/config"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const (
	tableNameSignalEvents = "signal_events"
)

func Connect() *sql.DB {
	var err error
	db, err = sql.Open(config.ConfigList.SQLDriver, config.ConfigList.DbName)
	if err != nil {
		panic(err)
	}
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT
	`)
	db.Exec(cmd)

	return db
}
