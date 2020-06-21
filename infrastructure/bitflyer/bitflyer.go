package infrastructure

import "IkezawaYuki/craft/domain"

type APIClient interface {
	Header(string, string, []byte) map[string]string
	DoRequest(string, string, map[string]string, []byte) ([]byte, error)
	GetBalance() ([]domain.Balance, error)
	GetTicker(string) (*domain.Ticker, error)
	GetRealTimeTicker(string, chan<- domain.Ticker)
}
