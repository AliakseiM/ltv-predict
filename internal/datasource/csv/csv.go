package csv

import (
	"os"

	"github.com/gocarina/gocsv"

	"github.com/AliakseiM/ltv-predict/internal/datasource"
	"github.com/AliakseiM/ltv-predict/internal/models"
)

type csvData struct {
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

func (d csvData) getRevenues() []float64 {
	r2 := d.LTV2 - d.LTV1
	r3 := d.LTV3 - d.LTV2

	r := []float64{r2, r3}

	if d.LTV4 != 0 {
		r = append(r, d.LTV4-d.LTV3)
	}
	if d.LTV5 != 0 {
		r = append(r, d.LTV5-d.LTV4)
	}
	if d.LTV6 != 0 {
		r = append(r, d.LTV6-d.LTV5)
	}
	if d.LTV7 != 0 {
		r = append(r, d.LTV7-d.LTV6)
	}

	return r
}

type Datasource struct {
	filePath string
	data     []*csvData
	grouped  map[string][]*csvData
}

func NewDatasource(filePath string) *Datasource {
	return &Datasource{
		filePath: filePath,
		data:     make([]*csvData, 0),
		grouped:  make(map[string][]*csvData),
	}
}

var _ datasource.Datasource = &Datasource{}

func (ds *Datasource) LoadData() error {
	f, err := os.Open(ds.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data := make([]*csvData, 0)

	if err := gocsv.UnmarshalFile(f, &data); err != nil {
		return err
	}

	ds.data = data

	return nil
}

func (ds *Datasource) GroupBy(col models.AggregateType) {
	grouped := make(map[string][]*csvData)

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
		prepared[group] = ds.getAverageRevenue(data)
	}

	return prepared, nil
}

func (ds *Datasource) getAverageRevenue(data []*csvData) []float64 {
	revenues := make([][]float64, 0, len(data))

	for _, d := range data {
		revenues = append(revenues, d.getRevenues())
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
