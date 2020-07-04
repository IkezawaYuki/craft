package main

import (
	"IkezawaYuki/craft/config"
	infrastructure "IkezawaYuki/craft/infrastructure/datastore"
	"IkezawaYuki/craft/interfaces/controllers"
	"IkezawaYuki/craft/logger"
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

	handler := infrastructure.NewSQLiteHandler(infrastructure.Connect())

	ctr := controllers.NewBitlyerController()
	ctr.StreamIngestionData()
}
