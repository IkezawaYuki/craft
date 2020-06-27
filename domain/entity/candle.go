package entity

import (
	"fmt"
	"time"
)

type Candle struct {
	ProductCode string
	Duration    time.Duration
	Time        time.Time
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
}

func (c *Candle) GetCandleTableName() string {
	return fmt.Sprintf("%s_%s", c.ProductCode, c.Duration)
}
