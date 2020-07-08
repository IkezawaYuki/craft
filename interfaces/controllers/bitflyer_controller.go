package controllers

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/domain/model"
	infrastructure "IkezawaYuki/craft/infrastructure/bitflyer_client"
	"IkezawaYuki/craft/interfaces/adapter"
	"IkezawaYuki/craft/interfaces/bitflyer"
	"IkezawaYuki/craft/interfaces/datastore"
	"IkezawaYuki/craft/logger"
	"IkezawaYuki/craft/usecase"
	"fmt"
	"strconv"
)

type BitlyerController interface {
	StreamIngestionData(c Context)
	ApiCandleHandler(c Context)
}

type bitflyerController struct {
	bitflyerClient infrastructure.APIClient
	bitlyerUsecase usecase.BitFlyerUsecase
}

func NewBitlyerController(sqlH datastore.SQLHandler) BitlyerController {
	bitRepo := adapter.NewCandleRepository(sqlH)
	apiClient := bitflyer.NewApiClient(config.ConfigList.APIKey, config.ConfigList.APISecret)
	bitUsecase := usecase.NewBitFlyerUsecase(bitRepo)
	return &bitflyerController{
		bitlyerUsecase: bitUsecase,
		bitflyerClient: apiClient,
	}
}

func (b *bitflyerController) StreamIngestionData(c Context) {
	var tickerChannel = make(chan model.Ticker)
	go b.bitflyerClient.GetRealTimeTicker(config.ConfigList.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		logger.Info("StreamIngestionData", fmt.Sprintf("ticker: %v", ticker))
		c.JSON(200, ticker)
		for _, duration := range config.ConfigList.Durations {
			isCreated := b.bitlyerUsecase.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.ConfigList.TradeDuration {
				// todo
			}
		}
	}
}

func (b *bitflyerController) ApiCandleHandler(c Context) {
	//productCode := r.URL.Query().Get("product_code")
	productCode := c.Query("product_code")

	if productCode == "" {
		// todo error
		return
	}
	//strLimit := r.URL.Query().Get("limit")
	strLimit := c.Query("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	//duration := r.URL.Query().Get("duration")
	duration := c.Query("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.ConfigList.Durations[duration]

	df := b.bitlyerUsecase.FindAllCandle(productCode, durationTime, limit)

	//byte, err := json.Marshal(df)
	//if err != nil {
	//	// todo error
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(byte)
	c.JSON(200, df)
}
