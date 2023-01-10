package json

import (
	"encoding/json"
	"io"
	"os"

	"github.com/AliakseiM/ltv-predict/internal/datasource"
	"github.com/AliakseiM/ltv-predict/internal/models"
)

type jsonData struct {
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

func (d jsonData) getRevenues() []float64 {
	r2 := d.LTV2 - d.LTV1
	r3 := d.LTV3 - d.LTV2
	r4 := d.LTV4 - d.LTV3
	r5 := d.LTV5 - d.LTV4
	r6 := d.LTV6 - d.LTV5
	r7 := d.LTV7 - d.LTV6

	return []float64{r2, r3, r4, r5, r6, r7}
}

type Datasource struct {
	filePath string
	data     []*jsonData
	grouped  map[string][]*jsonData
}

func NewDatasource(filePath string) *Datasource {
	return &Datasource{
		filePath: filePath,
		data:     make([]*jsonData, 0),
		grouped:  make(map[string][]*jsonData),
	}
}

var _ datasource.Datasource = &Datasource{}

func (ds *Datasource) LoadData() error {
	f, err := os.Open(ds.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	data := make([]*jsonData, 0)

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	ds.data = data

	return nil
}

func (ds *Datasource) GroupBy(col models.AggregateType) {
	grouped := make(map[string][]*jsonData)

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
		prepared[group] = ds.getAverageWeightedRevenue(data)
	}

	return prepared, nil
}

func (ds *Datasource) getAverageWeightedRevenue(data []*jsonData) []float64 {
	weightedRevenuesByDay := make(map[int][]float64)
	var weightSum float64

	for _, d := range data {

		for day, rev := range d.getRevenues() {
			weightedRevenuesByDay[day] = append(weightedRevenuesByDay[day], rev*float64(d.Users))
		}

		weightSum += float64(d.Users)

	}

	avgWeightedByDay := make(map[int]float64)

	for day, wrbd := range weightedRevenuesByDay {
		var sum float64
		for _, v := range wrbd {
			sum += v
		}
		avgWeightedByDay[day] = sum / weightSum
	}

	res := make([]float64, len(avgWeightedByDay))
	for i, v := range avgWeightedByDay {
		res[i] = v
	}

	return res
}
