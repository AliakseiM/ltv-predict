package models

type SourceType string

const (
	SourceTypeCSV  SourceType = "csv"
	SourceTypeJSON SourceType = "json"
)

func (sc SourceType) String() string {
	return string(sc)
}

func (sc SourceType) IsValid() bool {
	switch sc {
	case SourceTypeCSV:
		return true
	case SourceTypeJSON:
		return true
	default:
		return false
	}
}
