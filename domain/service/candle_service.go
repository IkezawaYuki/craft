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
