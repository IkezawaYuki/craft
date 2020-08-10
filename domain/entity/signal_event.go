package entity

import "time"

// SignalEvent is ...
type SignalEvent struct {
	Time        time.Time `json:"time"`
	ProductCode string    `json:"product_code"`
	Side        string    `json:"side"`
	Price       float64   `json:"price"`
	Size        float64   `json:"size"`
}

// SignalEvents is ...
type SignalEvents struct {
	Signals []SignalEvent `json:"signals"`
}

// NewSignalEvents is ...
func NewSignalEvents() *SignalEvents {
	return &SignalEvents{}
}
