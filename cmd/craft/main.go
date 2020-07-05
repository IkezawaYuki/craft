package main

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/infrastructure/server"
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

	// サーバーの起動
	server.Init()
}
