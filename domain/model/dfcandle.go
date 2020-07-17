package model

import (
	"IkezawaYuki/craft/domain/entity"
	"time"
)

type DataFrameCandle struct {
	ProductCode string          `json:"product_code"`
	Duration    time.Duration   `json:"duration"`
	Candles     []entity.Candle `json:"candles"`
	Smas        []Sma           `json:"smas"`
}

func (df *DataFrameCandle) Highs() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.High
	}
	return s
}

func (df *DataFrameCandle) Low() []float64 {

}
