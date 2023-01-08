package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSourceType_String(t *testing.T) {
	testCases := []struct {
		at       SourceType
		expected string
	}{
		{SourceTypeJSON, "json"},
		{SourceTypeCSV, "csv"},
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

func TestSourceType_IsValid(t *testing.T) {
	testCases := []struct {
		at       SourceType
		expected bool
	}{
		{SourceTypeJSON, true},
		{SourceTypeCSV, true},
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
