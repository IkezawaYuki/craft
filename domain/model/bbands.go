package model

type BBands struct {
	N    int       `json:"n,omitempty"`
	K    []float64 `json:"k,omitempty"`
	Up   []float64 `json:"up,omitempty"`
	Mid  []float64 `json:"mid,omitempty"`
	Down []float64 `json:"down,omitempty"`
}
