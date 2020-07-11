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
	"net/http"
	"strconv"
)

type BitlyerController interface {
	StreamIngestionData()
	ApiCandleHandler(c Context)
	ViewChart(c Context)
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

func (b *bitflyerController) StreamIngestionData() {
	var tickerChannel = make(chan model.Ticker)
	go b.bitflyerClient.GetRealTimeTicker(config.ConfigList.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		logger.Info("StreamIngestionData", fmt.Sprintf("ticker: %v", ticker))
		for _, duration := range config.ConfigList.Durations {
			isCreated := b.bitlyerUsecase.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.ConfigList.TradeDuration {
				// todo
			}
		}
	}
}

func (b *bitflyerController) ApiCandleHandler(c Context) {
	productCode := c.Query("product_code")

	if productCode == "" {
		c.JSON(400, "product code is empty")
		return
	}
	//strLimit := r.URL.Query().Get("limit")
	strLimit := c.Query("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	duration := c.Query("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.ConfigList.Durations[duration]

	df, err := b.bitlyerUsecase.FindAllCandle(productCode, durationTime, limit)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, df)
}

func (b *bitflyerController) ViewChart(c Context) {
	limit := 100
	duration := "1m"
	durationTime := config.ConfigList.Durations[duration]
	df, err := b.bitlyerUsecase.FindAllCandle(config.ConfigList.ProductCode, durationTime, limit)
	if err != nil {
		logger.Error("FindAllCandle()", err)
	}
	if df == nil {
		logger.Info("candle is empty")
		return
	}
	c.HTML(http.StatusOK, "google.html", df.Candles)
}
