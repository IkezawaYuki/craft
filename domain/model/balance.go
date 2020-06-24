package model

type Balance struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}
