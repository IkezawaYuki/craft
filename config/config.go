package config

import (
	"os"
	"strconv"
	"time"
)

var ConfigList List

type List struct {
	APIKey      string
	APISecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Port          int
}

func Init() {
	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	ConfigList = List{
		APIKey:        os.Getenv("API_KEY"),
		APISecret:     os.Getenv("API_SECRET"),
		LogFile:       os.Getenv("LOG_FILE"),
		ProductCode:   os.Getenv("PRODUCT_CODE"),
		TradeDuration: durations[os.Getenv("TRADE_DURATION")],
		Durations:     durations,
		DbName:        os.Getenv("DB_NAME"),
		SQLDriver:     os.Getenv("SQL_DRIVER"),
		Port:          port,
	}
}
