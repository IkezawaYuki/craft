package usecase

import (
	"IkezawaYuki/craft/config"
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/model"
	"IkezawaYuki/craft/domain/repository"
	"IkezawaYuki/craft/logger"
	"fmt"
	"time"
)

type bitflyerUsecase struct {
	candleRepo repository.CandleRepository
	eventRepo  repository.SignalEventRepository
}

func NewBitFlyerUsecase(canRepo repository.CandleRepository, eveRepo repository.SignalEventRepository) BitFlyerUsecase {
	return &bitflyerUsecase{
		candleRepo: canRepo,
		eventRepo:  eveRepo,
	}
}

type BitFlyerUsecase interface {
	CreateCandleWithDuration(model.Ticker, string, time.Duration) bool
	FindAllCandle(string, time.Duration, int) (*model.DataFrameCandle, error)
	Buy(productCode string, time time.Time, price, size float64, save bool) (*entity.SignalEvent, bool)
	Sell(productCode string, time time.Time, price, size float64, save bool) (*entity.SignalEvent, bool)
}

func (u *bitflyerUsecase) CreateCandleWithDuration(ticker model.Ticker, productCode string, duration time.Duration) bool {
	currentCandle := u.candleRepo.FindByTime(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	if currentCandle == nil {
		candle := entity.NewCandle(productCode, duration, ticker.TruncateDateTime(duration),
			price, price, price, price, ticker.Volume)
		_ = u.candleRepo.Create(candle)
		return true
	}

	if currentCandle.High <= price {
		currentCandle.High = price
	} else if currentCandle.Low >= price {
		currentCandle.Low = price
	}
	currentCandle.Volume += ticker.Volume
	currentCandle.Close = price
	currentCandle.ProductCode = productCode
	currentCandle.Duration = duration
	err := u.candleRepo.Update(currentCandle)
	if err != nil {
		panic(err)
	}
	return true
}

// FindAllCandle is ...
func (u *bitflyerUsecase) FindAllCandle(productCode string, durationTime time.Duration, limit int) (*model.DataFrameCandle, error) {
	df, err := u.candleRepo.FindAllCandle(productCode, durationTime, limit)
	if err != nil {
		logger.Error(fmt.Sprintf("FindAllCandle(%v, %v, %d)",
			config.ConfigList.ProductCode, durationTime, limit), err)
		return nil, err
	}
	return df, nil
}

// Buy is ...
func (u *bitflyerUsecase) Buy(productCode string, time time.Time, price, size float64, save bool,
) (*entity.SignalEvent, bool) {
	s := entity.NewSignalEvents()
	if s.CanBuy(time) {
		return nil, false
	}

	signalEvent := entity.SignalEvent{
		Time:        time,
		ProductCode: productCode,
		Side:        "BUY",
		Price:       price,
		Size:        size,
	}

	if save {
		u.eventRepo.Save(&signalEvent)
	}
	return &signalEvent, true
}

// Sell is ...
func (u *bitflyerUsecase) Sell(productCode string, time time.Time, price, size float64, save bool,
) (*entity.SignalEvent, bool) {
	s := entity.NewSignalEvents()
	if s.CanBuy(time) {
		return nil, false
	}

	signalEvent := entity.SignalEvent{
		Time:        time,
		ProductCode: productCode,
		Side:        "SELL",
		Price:       price,
		Size:        size,
	}

	if save {
		u.eventRepo.Save(&signalEvent)
	}
	return &signalEvent, true
}
