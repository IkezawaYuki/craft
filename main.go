package main

import (
	"IkezawaYuki/craft/config"
	infrastructure "IkezawaYuki/craft/infrastructure/datastore"
	interfaces "IkezawaYuki/craft/interfaces/bitflyer"
	"IkezawaYuki/craft/logger"
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	// 設定ファイルの読み込み
	config.Init()

	// ログの設定
	logger.SettingInit(config.ConfigList.LogFile)

	apiClient := interfaces.NewApiClient(config.ConfigList.APIKey, config.ConfigList.APISecret)

	fmt.Println(apiClient)

	conn := infrastructure.Connect()
	fmt.Println(conn)
}
