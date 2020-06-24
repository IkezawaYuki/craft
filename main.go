package main

import (
	"IkezawaYuki/craft/domain/model"
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

	order := &model.Order{
		ProductCode:     os.Getenv("PRODUCT_CODE"),
		ChildOrderType:  "MARKET",
		Side:            "BUY",
		Size:            0.001,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
}
