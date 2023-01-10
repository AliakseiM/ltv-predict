package linearRegression

import (
	"github.com/sajari/regression"
)

type LinearRegression struct{}

func New() *LinearRegression {
	return &LinearRegression{}
}

func (LinearRegression) PredictForDay(data []float64, day int) (float64, error) {
	r := new(regression.Regression)

	r.SetObserved("ARPU")
	r.SetVar(0, "Day")

	dataPoints := make(regression.DataPoints, 0, len(data))

	for i, p := range data {
		dp := regression.DataPoint(p, []float64{float64(i + 2)})
		dataPoints = append(dataPoints, dp)
	}

	r.Train(dataPoints...)

	if err := r.Run(); err != nil {
		return 0, err
	}

	prediction, err := r.Predict([]float64{float64(day)})
	if err != nil {
		return 0, err
	}

	return prediction, nil
}
