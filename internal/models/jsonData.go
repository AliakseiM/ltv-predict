package models

type JSONData struct {
	Users      int64   `json:"Users"`
	CampaignID string  `json:"CampaignId"`
	Country    string  `json:"Country"`
	LTV1       float64 `json:"Ltv1"`
	LTV2       float64 `json:"Ltv2"`
	LTV3       float64 `json:"Ltv3"`
	LTV4       float64 `json:"Ltv4"`
	LTV5       float64 `json:"Ltv5"`
	LTV6       float64 `json:"Ltv6"`
	LTV7       float64 `json:"Ltv7"`
}
