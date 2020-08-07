package adapter

import (
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/repository"
	sqlitehandler "IkezawaYuki/craft/infrastructure/sqlite_handler"
	"IkezawaYuki/craft/interfaces/datastore"
	"fmt"
)

type eventsRepository struct {
	sqlHandler datastore.SQLHandler
}

func NewEventsRepository(sHandler datastore.SQLHandler) repository.EventsRepository {
	return &eventsRepository{sqlHandler: sHandler}
}

const saveStmt = `INSERT INTO %s (time, product_code, side, price, size) VALUES (?, ?, ?, ?, ?)`

func (r *eventsRepository) Save(events *entity.Events) bool {
	cmd := fmt.Sprintf(saveStmt, sqlitehandler.TableNameSignalEvents)
	_, err := r.sqlHandler.Exec(cmd)
}
