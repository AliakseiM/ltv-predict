package exponentialSmoothing

import (
	"github.com/AliakseiM/ltv-predict/internal/predictor"
)

type ExponentialSmoothing struct {
	alpha, beta float64
}

func New(alpha, beta float64) *ExponentialSmoothing {
	return &ExponentialSmoothing{
		alpha: alpha,
		beta:  beta,
	}
}

var _ predictor.Predictor = &ExponentialSmoothing{}

func (es *ExponentialSmoothing) PredictForDay(data []float64, day int) (float64, error) {
	level := []float64{data[0]}
	trend := []float64{0}

	for i := 1; i < len(data); i++ {
		l := es.alpha*data[i] + (1-es.alpha)*(level[i-1]+trend[i-1])
		t := es.beta*(l-level[i-1]) + (1-es.beta)*trend[i-1]
		level = append(level, l)
		trend = append(trend, t)
	}

	futureValues := make([]float64, 0)
	for i := 0; i < day-len(data)-1; i++ {
		futureValue := level[len(level)-1] + trend[len(trend)-1]*(float64(i+1))
		futureValues = append(futureValues, futureValue)
	}

	return futureValues[len(futureValues)-1], nil
}
