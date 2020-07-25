package service

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

func (s *candleService) IchimokuCloud(inReal []float64) ([]float64, []float64, []float64, []float64, []float64) {
	length := len(inReal)
	tenkan := make([]float64, min(9, length))
}
