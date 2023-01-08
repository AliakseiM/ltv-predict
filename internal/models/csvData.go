package models

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	csvFile = "data/test_data.csv"
)

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

func PrintCSVInput() error {
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	csvData := make([]*CSVData, 0)

	if err := gocsv.UnmarshalFile(f, &csvData); err != nil {
		return err
	}

	for _, d := range csvData {
		fmt.Println(d)
		fmt.Println()
	}

	return nil
}
