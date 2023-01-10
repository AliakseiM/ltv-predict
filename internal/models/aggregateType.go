package models

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
		return true
	case AggregateTypeCampaign:
		return true
	default:
		return false
	}
}
