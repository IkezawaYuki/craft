package bitflyer

import (
	"IkezawaYuki/craft/domain/model"
	infrastructure "IkezawaYuki/craft/infrastructure/bitflyer_client"
	"IkezawaYuki/craft/logger"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1/"

type apiClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func NewApiClient(key, secret string) infrastructure.APIClient {
	return &apiClient{
		key:        key,
		secret:     secret,
		httpClient: &http.Client{},
	}
}

func (a *apiClient) Header(method string, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	logger.Info(timestamp)
	message := timestamp + method + endpoint + string(body)
	mac := hmac.New(sha256.New, []byte(a.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       a.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (a *apiClient) DoRequest(method string, urlPath string, query map[string]string, data []byte) ([]byte, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	logger.Info("doRequest", fmt.Sprintf("endpoint:%s", endpoint))
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range a.Header(method, req.URL.RequestURI(), data) {
		req.Header.Add(k, v)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *apiClient) GetBalance() ([]model.Balance, error) {
	url := "me/getbalance"
	resp, err := a.DoRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		logger.Error("getBalance() is error", err)
		return nil, err
	}
	fmt.Println(string(resp))
	logger.Info("getBalance()", fmt.Sprintf("url:%s", url), fmt.Sprintf("resp:%s", resp))
	var balance []model.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		logger.Error("json unmarshal is failed", err)
		return nil, err
	}
	return balance, nil
}

func (a *apiClient) GetTicker(productCode string) (*model.Ticker, error) {
	url := "ticker"
	resp, err := a.DoRequest("GET", url, map[string]string{"product_code": productCode}, nil)
	if err != nil {
		logger.Error("GetTicker() is error", err)
		return nil, err
	}
	fmt.Println(string(resp))
	logger.Info("GetTicker()", fmt.Sprintf("url:%s", url), fmt.Sprintf("resp:%s", resp))
	var ticker model.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		logger.Error("json unmarshal is failed", err)
		return nil, err
	}
	return &ticker, nil
}

type JsonRPC2 struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	Id      *int        `json:"id,omitempty"`
}

type SubscribeParams struct {
	Channel string `json:"channel"`
}

func (a *apiClient) GetRealTimeTicker(symbol string, ch chan<- model.Ticker) {
	u := url.URL{
		Scheme: "wss",
		Host:   "ws.lightstream.bitflyer.com",
		Path:   "/json-rpc",
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	if err := c.WriteJSON(&JsonRPC2{
		Version: "2.0",
		Method:  "subscribe",
		Params:  &SubscribeParams{channel},
	}); err != nil {
		log.Fatal("subscribe:", err)
	}

OUTER:
	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			switch v := message.Params.(type) {
			case map[string]interface{}:
				for key, binary := range v {
					if key == "message" {
						marshalTic, err := json.Marshal(binary)
						if err != nil {
							continue OUTER
						}
						var ticker model.Ticker
						if err := json.Unmarshal(marshalTic, &ticker); err != nil {
							continue OUTER
						}
						ch <- ticker
					}
				}
			}
		}
	}
}

func (a *apiClient) SendOrder(order *model.Order) (*model.ResponseSendChildOrder, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	u := "me/sendchildorder"
	resp, err := a.DoRequest("POST", u, map[string]string{}, data)
	if err != nil {
		return nil, err
	}
	var response model.ResponseSendChildOrder
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *apiClient) ListOrder(query map[string]string) ([]model.Order, error) {
	resp, err := a.DoRequest(http.MethodGet, "me/getchildorders", query, nil)
	if err != nil {
		return nil, err
	}
	var responseListOrder []model.Order
	err = json.Unmarshal(resp, &responseListOrder)
	if err != nil {
		return nil, err
	}
	return responseListOrder, nil
}
