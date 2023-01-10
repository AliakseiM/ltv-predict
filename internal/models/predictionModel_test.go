package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPredictionModel_String(t *testing.T) {
	testCases := []struct {
		pm       PredictionModel
		expected string
	}{
		{LinearRegression, "lr"},
		{ExponentialSmoothing, "es"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.pm), func(t *testing.T) {
			actual := tc.pm.String()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestPredictionModel_IsValid(t *testing.T) {
	testCases := []struct {
		pm       PredictionModel
		expected bool
	}{
		{LinearRegression, true},
		{ExponentialSmoothing, true},
		{"invalid", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.pm), func(t *testing.T) {
			actual := tc.pm.IsValid()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
