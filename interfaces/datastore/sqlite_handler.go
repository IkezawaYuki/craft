package interfaces

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/model"
	"IkezawaYuki/craft/domain/repository"
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
	stmt := fmt.Sprintf(createStmt, candle.GetCandleTableName())
	_, err := c.conn.Exec(stmt, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
	if err != nil {
		return err
	}
	return nil
}

func (c *candleRepository) Update(candle *entity.Candle) error {
	panic("implement me")
}

func (c *candleRepository) FindByTime(s string, duration time.Duration, time time.Time) *entity.Candle {
	panic("implement me")
}

func (c *candleRepository) CreateCandleWithDuration(ticker model.Ticker, s string, duration time.Duration) {
	panic("implement me")
}
