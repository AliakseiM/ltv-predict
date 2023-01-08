package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	jsonFile = "data/test_data.json"
)

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

func PrintJSONInput() error {
	f, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	jsonData := make([]*JSONData, 0)

	if err := json.Unmarshal(b, &jsonData); err != nil {
		return err
	}

	for _, d := range jsonData {
		fmt.Println(d)
		fmt.Println()
	}

	return nil
}