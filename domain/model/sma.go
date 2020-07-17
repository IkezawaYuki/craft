package model

type Sma struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}
