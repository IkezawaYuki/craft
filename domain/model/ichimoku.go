package model

type IchimokuCloud struct {
	Tenkan  []float64 `json:"tenkan,omitempty"`
	Kijun   []float64 `json:"kijun,omitempty"`
	SenkouA []float64 `json:"senkou_a,omitempty"`
	SenkouB []float64 `json:"senkou_b,omitempty"`
	Chikou  []float64 `json:"chikou,omitempty"`
}
