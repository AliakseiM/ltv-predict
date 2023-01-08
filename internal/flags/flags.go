package flags

import (
	"errors"

	"github.com/AliakseiM/ltv-predict/internal/models"
)

type Flag string

const (
	Model     Flag = "model"
	Source    Flag = "source"
	Aggregate Flag = "aggregate"
)

func (f Flag) String() string {
	return string(f)
}

func (f Flag) Shorthand() string {
	return string(f[0])
}

var (
	//errUnsupportedModel     = errors.New("unsupported model")
	errUnsupportedSource    = errors.New("unsupported source")
	errUnsupportedAggregate = errors.New("unsupported aggregate")
)

func ValidateValues(model, source, aggregate string) error {
	// TODO: validate model

	if s := models.SourceType(source); !s.IsValid() {
		return errUnsupportedSource
	}

	if a := models.AggregateType(aggregate); !a.IsValid() {
		return errUnsupportedAggregate
	}

	return nil
}
