package adapter

import (
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/repository"
	sqlitehandler "IkezawaYuki/craft/infrastructure/sqlite_handler"
	"IkezawaYuki/craft/interfaces/datastore"
	"IkezawaYuki/craft/logger"
	"fmt"
	"strings"
)

type eventsRepository struct {
	sqlHandler datastore.SQLHandler
}

func NewEventsRepository(sHandler datastore.SQLHandler) repository.SignalEventRepository {
	return &eventsRepository{sqlHandler: sHandler}
}

const saveStmt = `INSERT INTO %s (time, product_code, side, price, size) VALUES (?, ?, ?, ?, ?)`

func (r *eventsRepository) Save(events *entity.SignalEvent) bool {
	cmd := fmt.Sprintf(saveStmt, sqlitehandler.TableNameSignalEvents)
	_, err := r.sqlHandler.Exec(cmd, events.Time, events.ProductCode, events.Side, &events.Price, &events.Size)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			logger.Error("UNIQUE", err)
			return true
		}
		return false
	}
	return true
}
