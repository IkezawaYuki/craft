package main

import (
	"IkezawaYuki/craft/domain"
	interfaces "IkezawaYuki/craft/interfaces/bitflyer"
	"IkezawaYuki/craft/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	logger.SettingInit(os.Getenv("LOG_FILE"))
	fmt.Println(os.Getenv("API_KEY"))

	apiClient := interfaces.NewApiClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	ticker, _ := apiClient.GetTicker("BTC_JPY")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))

	tickerChannel := make(chan domain.Ticker)
	go apiClient.GetRealTimeTicker(os.Getenv("PRODUCT_CODE"), tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Second))
	}
}
