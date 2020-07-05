package bitflyerclient

import (
	"IkezawaYuki/craft/domain/model"
)

type APIClient interface {
	Header(string, string, []byte) map[string]string
	DoRequest(string, string, map[string]string, []byte) ([]byte, error)
	GetBalance() ([]model.Balance, error)
	GetTicker(string) (*model.Ticker, error)
	GetRealTimeTicker(string, chan<- model.Ticker)
	SendOrder(*model.Order) (*model.ResponseSendChildOrder, error)
	ListOrder(map[string]string) ([]model.Order, error)
}
