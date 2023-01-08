package flags

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestFlag_String(t *testing.T) {
	testCases := []struct {
		f        Flag
		expected string
	}{
		{Model, "model"},
		{Source, "source"},
		{Aggregate, "aggregate"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.f), func(t *testing.T) {
			actual := tc.f.String()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestFlag_Shorthand(t *testing.T) {
	testCases := []struct {
		f        Flag
		expected string
	}{
		{Model, "m"},
		{Source, "s"},
		{Aggregate, "a"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.f), func(t *testing.T) {
			actual := tc.f.Shorthand()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestValidateValues(t *testing.T) {
	testCases := []struct {
		model, source, aggregate string
		expectedErr              error
	}{
		{"", "json", "country", nil},
		{"", "csv", "campaign", nil},
		{"", "invalid", "country", errUnsupportedSource},
		{"", "invalid", "campaign", errUnsupportedSource},
		{"", "json", "invalid", errUnsupportedAggregate},
		{"", "csv", "invalid", errUnsupportedAggregate},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v %v", tc.model, tc.source, tc.aggregate), func(t *testing.T) {
			err := ValidateValues(tc.model, tc.source, tc.aggregate)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
