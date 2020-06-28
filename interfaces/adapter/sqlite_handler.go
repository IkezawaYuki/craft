package adapter

import (
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/model"
	"IkezawaYuki/craft/domain/repository"
	"IkezawaYuki/craft/logger"
	"database/sql"
	"fmt"
	"time"
)

type candleRepository struct {
	conn *sql.DB
}

func NewCandleRepository(conn *sql.DB) repository.CandleRepository {
	return &candleRepository{conn: conn}
}

const createStmt = `INSERT INTO %s (time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)`

func (c *candleRepository) Create(candle *entity.Candle) error {
	stmt := fmt.Sprintf(createStmt, entity.GetCandleTableName(candle.ProductCode, candle.Duration))
	_, err := c.conn.Exec(stmt, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
	if err != nil {
		return err
	}
	return nil
}

const updateStmt = `UPDATE %s SET open = ?, close = ?, high = ?, low = ?, volume = ? WHERE time = ?`

func (c *candleRepository) Update(candle *entity.Candle) error {
	stmt := fmt.Sprintf(updateStmt, entity.GetCandleTableName(candle.ProductCode, candle.Duration))
	_, err := c.conn.Exec(stmt, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume, candle.Time)
	if err != nil {
		return err
	}
	return nil
}

const selectStmt = `SELECT time, open, close, high, low, volume FROM %s WHERE time = ?`

func (c *candleRepository) FindByTime(productCode string, duration time.Duration, time time.Time) *entity.Candle {
	tableName := entity.GetCandleTableName(productCode, duration)
	stmt := fmt.Sprintf(selectStmt, tableName)
	row := c.conn.QueryRow(stmt, time)
	var candle entity.Candle
	if err := row.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume); err != nil {
		logger.Error("FindByTime", err)
		return nil
	}
	return &candle
}

const findAllStmt = `SELECT * FROM (
	SELECT time, open, close, high, low, volume FROM %s ORDER BY time DESC LIMIT ?
	) ORDER BY time ASC`

func (c *candleRepository) FindAllCandle(productCode string, duration time.Duration, limit int) *model.DataFrameCandle {
	tableName := entity.GetCandleTableName(productCode, duration)
	stmt := fmt.Sprintf(findAllStmt, tableName)
	rows, err := c.conn.Query(stmt, limit)
	if err != nil {
		logger.Error("FindAllCandle Query()", err)
		return nil
	}
	var dfCandle model.DataFrameCandle
	dfCandle.ProductCode = productCode
	dfCandle.Duration = duration
	for rows.Next() {
		var candle entity.Candle
		if err := rows.Scan(&candle.Time, &candle.Open, &candle.High, &candle.Low, &candle.Volume); err != nil {
			logger.Error("FindAllCandle rows.Scan()", err)
			return nil
		}
		dfCandle.Candles = append(dfCandle.Candles, candle)
	}
	return &dfCandle
}
