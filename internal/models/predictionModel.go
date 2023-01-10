package models

import (
	"fmt"
)

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
		fmt.Println("linear regression")
		return true
	case ExponentialSmoothing:
		fmt.Println("exponential smoothing")
		return true
	default:
		return false
	}
}
