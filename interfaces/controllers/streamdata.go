package controllers

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/domain/model"
	infrastructure "IkezawaYuki/craft/infrastructure/bitflyer"
	"IkezawaYuki/craft/interfaces/adapter"
	"IkezawaYuki/craft/interfaces/bitflyer"
	"IkezawaYuki/craft/logger"
	"IkezawaYuki/craft/usecase"
	"database/sql"
	"fmt"
)

type BitlyerController interface {
	StreamIngestionData()
}

type bitflyerController struct {
	bitflyerClient infrastructure.APIClient
	bitlyerUsecase usecase.BitFlyerUsecase
}

func NewBitlyerController(db *sql.DB) BitlyerController {
	bitRepo := adapter.NewCandleRepository(db)
	apiClient := bitflyer.NewApiClient(config.ConfigList.APIKey, config.ConfigList.APISecret)
	bitUsecase := usecase.NewBitFlyerUsecase(bitRepo)
	return &bitflyerController{
		bitlyerUsecase: bitUsecase,
		bitflyerClient: apiClient,
	}
}

func (c *bitflyerController) StreamIngestionData() {
	var tickerChannel = make(chan model.Ticker)
	go c.bitflyerClient.GetRealTimeTicker(config.ConfigList.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		logger.Info("StreamIngestionData", fmt.Sprintf("ticker: %v", ticker))
		for _, duration := range config.ConfigList.Durations {
			isCreated := c.bitlyerUsecase.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.ConfigList.TradeDuration {
				// todo
			}
		}
	}
}