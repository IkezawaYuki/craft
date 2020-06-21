package interfaces

import (
	"IkezawaYuki/craft/domain"
	infrastructure "IkezawaYuki/craft/infrastructure/bitflyer"
	"IkezawaYuki/craft/logger"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (a *apiClient) GetBalance() ([]domain.Balance, error) {
	url := "me/getbalance"
	resp, err := a.DoRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		logger.Error("getBalance() is error", err)
		return nil, err
	}
	fmt.Println(string(resp))
	logger.Info("getBalance()", fmt.Sprintf("url:%s", url), fmt.Sprintf("resp:%s", resp))
	var balance []domain.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		logger.Error("json unmarshal is failed", err)
		return nil, err
	}
	return balance, nil
}
