package model

import (
	"IkezawaYuki/craft/domain/entity"
	"time"
)

type DataFrameCandle struct {
	ProductCode string          `json:"product_code"`
	Duration    time.Duration   `json:"duration"`
	Candles     []entity.Candle `json:"candles"`
}
