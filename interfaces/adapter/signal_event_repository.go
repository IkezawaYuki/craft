package adapter

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/repository"
	sqlitehandler "IkezawaYuki/craft/infrastructure/sqlite_handler"
	"IkezawaYuki/craft/interfaces/datastore"
	"IkezawaYuki/craft/logger"
	"fmt"
	"strings"
	"time"
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

const selectSignalStmt = `SELECT * FROM (
	SELECT time, product_code, side, price, size FROM %s WHERE product_code = ? ORDER BY time DESC LIMIT ?)
	ORDER BY time ASC;
`

func (r *eventsRepository) GetByCount(loadEvents int) *entity.SignalEvents {
	cmd := fmt.Sprintf(selectSignalStmt, sqlitehandler.TableNameSignalEvents)
	rows, err := r.sqlHandler.Query(cmd, config.ConfigList.ProductCode, loadEvents)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents entity.SignalEvents
	for rows.Next() {
		var signalEvent entity.SignalEvent
		if err := rows.Scan(
			&signalEvent.Time,
			&signalEvent.ProductCode,
			&signalEvent.Side,
			&signalEvent.Price,
			&signalEvent.Size,
		); err != nil {
			return nil
		}
		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
	}
	return &signalEvents
}

const selectSignalEventAfterTimeQuery = `SELECT * FROM (
	SELECT time, product_code, side, price, size FROM %s
	WHERE DATETIME(time) >= DATETIME(?)
	ORDER BY time DESC
	) ORDER BY time ASC;`

func (r *eventsRepository) GetSignalEventsAfterTime(timeTime time.Time) *entity.SignalEvents {
	cmd := fmt.Sprintf(selectSignalEventAfterTimeQuery, sqlitehandler.TableNameSignalEvents)
	rows, err := r.sqlHandler.Query(cmd, timeTime.Format(time.RFC3339))
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents entity.SignalEvents
	for rows.Next() {
		var signalEvent entity.SignalEvent
		if err := rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side,
			&signalEvent.Price, &signalEvent.Size); err != nil {
			return nil
		}
		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
	}
	return &signalEvents
}

func (r *eventsRepository) Buy(productCode string, time time.Time, price, size float64, save bool) bool {

}

func (r *eventsRepository) Sell(productCode string, time time.Time, price, size float64, save bool) bool {

}
