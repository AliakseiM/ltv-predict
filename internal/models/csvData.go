package models

type CSVData struct {
	UserID     int64   `csv:"UserId"`
	CampaignID string  `csv:"CampaignId"`
	Country    string  `csv:"Country"`
	LTV1       float64 `csv:"Ltv1"`
	LTV2       float64 `csv:"Ltv2"`
	LTV3       float64 `csv:"Ltv3"`
	LTV4       float64 `csv:"Ltv4"`
	LTV5       float64 `csv:"Ltv5"`
	LTV6       float64 `csv:"Ltv6"`
	LTV7       float64 `csv:"Ltv7"`
}
