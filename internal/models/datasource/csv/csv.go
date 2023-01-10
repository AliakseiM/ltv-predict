package csv

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/AliakseiM/ltv-predict/internal/models"
)

type Datasource struct {
	filePath string
	data     []*models.CSVData
	grouped  map[string][]*models.CSVData
}

func NewDatasource(filePath string) *Datasource {
	return &Datasource{filePath: filePath}
}

func (ds *Datasource) LoadData() error {
	f, err := os.Open(ds.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvData := make([]*models.CSVData, 0)

	if err := gocsv.UnmarshalFile(f, &csvData); err != nil {
		return err
	}

	ds.data = csvData

	return nil
}

func (ds *Datasource) Print() {
	for _, d := range ds.data {
		fmt.Println(d)
	}
}

func (ds *Datasource) GroupBy(col models.AggregateType) {
	grouped := make(map[string][]*models.CSVData)

	for _, item := range ds.data {
		switch col {
		case models.AggregateTypeCountry:
			grouped[item.Country] = append(grouped[item.Country], item)
		case models.AggregateTypeCampaign:
			grouped[item.CampaignID] = append(grouped[item.CampaignID], item)
		default:
		}
	}

	ds.grouped = grouped

	return
}

func (ds *Datasource) Prepare() (map[string][]float64, error) {
	return nil, nil
}
