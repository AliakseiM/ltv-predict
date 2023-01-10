package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AliakseiM/ltv-predict/internal/models"
)

type Datasource struct {
	filePath string
	data     []*models.JSONData
	grouped  map[string][]*models.JSONData
}

func NewDatasource(filePath string) *Datasource {
	return &Datasource{
		filePath: filePath,
		data:     make([]*models.JSONData, 0),
		grouped:  make(map[string][]*models.JSONData),
	}
}

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

	jsonData := make([]*models.JSONData, 0)

	if err := json.Unmarshal(b, &jsonData); err != nil {
		return err
	}

	ds.data = jsonData

	return nil
}

func (ds *Datasource) Print() {
	for _, d := range ds.data {
		fmt.Println(d)
	}
}

func (ds *Datasource) GroupBy(col models.AggregateType) {
	grouped := make(map[string][]*models.JSONData)

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
