package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAggregateType_String(t *testing.T) {
	testCases := []struct {
		at       AggregateType
		expected string
	}{
		{AggregateTypeCountry, "country"},
		{AggregateTypeCampaign, "campaign"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.at), func(t *testing.T) {
			actual := tc.at.String()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestAggregateType_IsValid(t *testing.T) {
	testCases := []struct {
		at       AggregateType
		expected bool
	}{
		{AggregateTypeCountry, true},
		{AggregateTypeCampaign, true},
		{"invalid", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.at), func(t *testing.T) {
			actual := tc.at.IsValid()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
