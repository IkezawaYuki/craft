package main

import (
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

type ConfigList struct {
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

func main() {
	logger.SettingInit(os.Getenv("LOG_FILE"))
	fmt.Println(os.Getenv("API_KEY"))

	apiClient := interfaces.NewApiClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	fmt.Println(apiClient)

}
