package usecase

import (
	"IkezawaYuki/craft/domain/entity"
	"IkezawaYuki/craft/domain/model"
	"IkezawaYuki/craft/domain/repository"
	"time"
)

type bitflyerUsecase struct {
	candleRepo repository.CandleRepository
}

func NewBitFlyerUsecase(canRepo repository.CandleRepository) BitFlyerUsecase {
	return &bitflyerUsecase{candleRepo: canRepo}
}

type BitFlyerUsecase interface {
	CreateCandleWithDuration(model.Ticker, string, time.Duration) bool
}

func (u *bitflyerUsecase) CreateCandleWithDuration(ticker model.Ticker, productCode string, duration time.Duration) bool {
	currentCandle := u.candleRepo.FindByTime(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	if currentCandle == nil {
		candle := entity.NewCandle(productCode, duration, ticker.TruncateDateTime(duration),
			price, price, price, price, ticker.Volume)
		err := u.candleRepo.Create(candle)
		if err != nil {
			panic(err)
		}
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
