package entity

import (
	"fmt"
	"time"
)

type Candle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Time        time.Time     `json:"time"`
	Open        float64       `json:"open"`
	Close       float64       `json:"close"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Volume      float64       `json:"volume"`
}

func NewCandle(productCode string, duration time.Duration, timeDate time.Time,
	open, close, high, low, volume float64) *Candle {
	return &Candle{
		ProductCode: productCode,
		Duration:    duration,
		Time:        timeDate,
		Open:        open,
		Close:       close,
		High:        high,
		Low:         low,
		Volume:      volume,
	}
}

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}
