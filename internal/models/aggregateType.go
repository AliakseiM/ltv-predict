package models

import (
	"fmt"
)

type AggregateType string

const (
	AggregateTypeCountry  AggregateType = "country"
	AggregateTypeCampaign AggregateType = "campaign"
)

func (at AggregateType) String() string {
	return string(at)
}

func (at AggregateType) IsValid() bool {
	switch at {
	case AggregateTypeCountry:
		fmt.Println("group by country")
		return true
	case AggregateTypeCampaign:
		fmt.Println("group by campaign")
		return true
	default:
		return false
	}
}
