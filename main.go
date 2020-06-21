package main

import (
	interfaces "IkezawaYuki/craft/interfaces/bitflyer"
	"IkezawaYuki/craft/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
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
	balance, err := apiClient.GetBalance()
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)
}
