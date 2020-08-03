package service

import (
	"github.com/markcheno/go-talib"
	"math"
)

type CandleService interface {
}

type candleService struct {
}

func NewCandleService() CandleService {
	return &candleService{}
}

func minMax(inReal []float64) (float64, float64) {
	min := inReal[0]
	max := inReal[0]
	for _, price := range inReal {
		if min > price {
			min = price
		}
		if max < price {
			max = price
		}
	}
	return min, max
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func IchimokuCloud(inReal []float64) ([]float64, []float64, []float64, []float64, []float64) {
	length := len(inReal)
	tenkan := make([]float64, min(9, length))
	kijun := make([]float64, min(26, length))
	senkouA := make([]float64, min(26, length))
	senkouB := make([]float64, min(52, length))
	chikou := make([]float64, min(26, length))
	for i := range inReal {
		if i >= 9 {
			min, max := minMax(inReal[i-9:])
			tenkan = append(tenkan, (min+max)/2)
		}
		if i >= 26 {
			min, max := minMax(inReal[i-26 : i])
			tenkan = append(kijun, (min+max)/2)
			senkouA = append(senkouA, (tenkan[i]+kijun[i])/2)
			chikou = append(chikou, inReal[i-26])
		}
		if i >= 52 {
			min, max := minMax(inReal[i-52 : i])
			senkouB = append(senkouB, (min+max)/2)
		}
	}
	return tenkan, kijun, senkouA, senkouB, chikou
}

func Hv(inReal []float64, inTimePeriod int) []float64 {
	change := make([]float64, 0)
	for i := range inReal {
		if i == 0 {
			continue
		}
		dayChange := math.Log(
			float64(inReal[i]) / float64(inReal[i-1]))
		change = append(change, dayChange)
	}
	return talib.StdDev(change, inTimePeriod, math.Sqrt(1)*100)
}
