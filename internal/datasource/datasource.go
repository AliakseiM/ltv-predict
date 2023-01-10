package datasource

import (
	"github.com/AliakseiM/ltv-predict/internal/models"
)

type Datasource interface {
	LoadData() error
	GroupBy(models.AggregateType)
	Prepare() (map[string][]float64, error)
}
