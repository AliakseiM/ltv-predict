package models

import (
	"fmt"
)

type SourceType string

const (
	SourceTypeCSV  SourceType = "csv"
	SourceTypeJSON SourceType = "json"
)

func (sc SourceType) String() string {
	return string(sc)
}

func (sc SourceType) IsValid() bool {
	switch sc {
	case SourceTypeCSV:
		fmt.Println("csv source")
		return true
	case SourceTypeJSON:
		fmt.Println("json source")
		return true
	default:
		return false
	}
}
