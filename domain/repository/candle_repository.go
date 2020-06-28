package repository

import (
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/model"
	"time"
)

type CandleRepository interface {
	Create(*entity.Candle) error
	Update(*entity.Candle) error
	FindByTime(string, time.Duration, time.Time) *entity.Candle
	FindAllCandle(string, time.Duration, int) *model.DataFrameCandle
}
