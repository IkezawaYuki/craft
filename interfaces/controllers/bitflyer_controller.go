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

	sma := c.Query("sma")
	if sma != "" {
		strSmaPeriod1 := c.Query("sma_period_1")
		strSmaPeriod2 := c.Query("sma_period_2")
		strSmaPeriod3 := c.Query("sma_period_3")
		period1, err := strconv.Atoi(strSmaPeriod1)
		if err != nil {
			period1 = 7
		}
		period2, err := strconv.Atoi(strSmaPeriod2)
		if err != nil {
			period2 = 14
		}
		period3, err := strconv.Atoi(strSmaPeriod3)
		if err != nil {
			period3 = 50
		}
		df.AddSma(period1)
		df.AddSma(period2)
		df.AddSma(period3)
	}

	ema := c.Query("ema")
	if ema != "" {
		strEmaPeriod1 := c.Query("ema_period_1")
		strEmaPeriod2 := c.Query("ema_period_2")
		strEmaPeriod3 := c.Query("ema_period_3")
		period1, err := strconv.Atoi(strEmaPeriod1)
		if err != nil {
			period1 = 7
		}
		period2, err := strconv.Atoi(strEmaPeriod2)
		if err != nil {
			period2 = 14
		}
		period3, err := strconv.Atoi(strEmaPeriod3)
		if err != nil {
			period3 = 50
		}
		df.AddEma(period1)
		df.AddEma(period2)
		df.AddEma(period3)
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
