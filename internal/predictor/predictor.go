package predictor

type Predictor interface {
	PredictForDay(data []float64, day int) (float64, error)
}
