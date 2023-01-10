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
	prepared := make(map[string][]float64, len(ds.grouped))

	for group, data := range ds.grouped {
		prepared[group] = ds.getAverage(data)
	}

	return prepared, nil
}

func (ds *Datasource) getAverage(data []*models.CSVData) []float64 {
	revenues := make([][]float64, 0, len(data))

	for _, d := range data {
		revenues = append(revenues, d.GetRevenues())
	}

	sums := make([]float64, 6)
	daysCount := make(map[int]int)
	for _, r := range revenues {
		for i, rbd := range r {
			sums[i] += rbd
			daysCount[i]++
		}
	}

	res := make([]float64, 0, len(sums))
	for day, sum := range sums {
		res = append(res, sum/float64(daysCount[day]))
	}

	return res
}
