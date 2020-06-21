package domain

type Balance struct {
	CurrentCode string  `json:"current_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}
