package models

type PredictionModel string

const (
	LinearRegression     PredictionModel = "lr"
	ExponentialSmoothing PredictionModel = "es"
)

func (pm PredictionModel) String() string {
	return string(pm)
}

func (pm PredictionModel) IsValid() bool {
	switch pm {
	case LinearRegression:
		return true
	case ExponentialSmoothing:
		return true
	default:
		return false
	}
}
