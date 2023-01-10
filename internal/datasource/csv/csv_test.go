package csv

import (
	"reflect"
	"testing"

	"github.com/AliakseiM/ltv-predict/internal/models"
)

func Test_csvData_getRevenues(t *testing.T) {
	testCases := []struct {
		name           string
		input          csvData
		expectedOutput []float64
	}{
		{
			name: "test case 1",
			input: csvData{
				UserID:     1,
				CampaignID: "campaign1",
				Country:    "US",
				LTV1:       10,
				LTV2:       20,
				LTV3:       30,
				LTV4:       0,
				LTV5:       0,
				LTV6:       0,
				LTV7:       0,
			},
			expectedOutput: []float64{10, 10},
		},
		{
			name: "test case 2",
			input: csvData{
				UserID:     1,
				CampaignID: "campaign1",
				Country:    "US",
				LTV1:       10,
				LTV2:       20,
				LTV3:       30,
				LTV4:       40,
				LTV5:       0,
				LTV6:       0,
				LTV7:       0,
			},
			expectedOutput: []float64{10, 10, 10},
		},
		{
			name: "test case 3",
			input: csvData{
				UserID:     1,
				CampaignID: "campaign1",
				Country:    "US",
				LTV1:       10,
				LTV2:       20,
				LTV3:       30,
				LTV4:       40,
				LTV5:       50,
				LTV6:       60,
				LTV7:       70,
			},
			expectedOutput: []float64{10, 10, 10, 10, 10, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.input.getRevenues(); !reflect.DeepEqual(got, tc.expectedOutput) {
				t.Errorf("getRevenues() = %v, want %v", got, tc.expectedOutput)
			}
		})
	}
}

func TestNewDatasource(t *testing.T) {
	ds := NewDatasource("test.csv")

	if ds == nil {
		t.Error("Expected Datasource to be not nil")
	}

	if ds.filePath != "test.csv" {
		t.Errorf("Expected filePath to be test.csv, got %s", ds.filePath)
	}

	if ds.data == nil {
		t.Errorf("Expected data to be initialized as an empty slice")
	}

	if ds.grouped == nil {
		t.Errorf("Expected grouped to be initialized as an empty map")
	}
}

func TestDatasource_LoadData(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "no file",
			filePath: "file_not_found.csv",
			wantErr:  true,
		},
		{
			name:     "success",
			filePath: "../../../test/data/ok.csv",
			wantErr:  false,
		},
		{
			name:     "incorrect structure",
			filePath: "../../../test/data/incorrect.csv",
			wantErr:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ds := NewDatasource(tc.filePath)
			if err := ds.LoadData(); (err != nil) != tc.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestDatasource_GroupBy(t *testing.T) {
	testCases := []struct {
		name     string
		data     []*csvData
		a        models.AggregateType
		expected map[string][]*csvData
	}{
		{
			name: "correct campaign",
			data: []*csvData{{CampaignID: "1"}, {CampaignID: "2"}, {CampaignID: "2"}, {CampaignID: "3"}, {CampaignID: "3"}, {CampaignID: "3"}},
			a:    models.AggregateTypeCampaign,
			expected: map[string][]*csvData{
				"1": {{CampaignID: "1"}},
				"2": {{CampaignID: "2"}, {CampaignID: "2"}},
				"3": {{CampaignID: "3"}, {CampaignID: "3"}, {CampaignID: "3"}},
			},
		},
		{
			name: "correct country",
			data: []*csvData{{Country: "IT"}, {Country: "FR"}, {Country: "FR"}, {Country: "US"}, {Country: "US"}, {Country: "US"}},
			a:    models.AggregateTypeCountry,
			expected: map[string][]*csvData{
				"IT": {{Country: "IT"}},
				"FR": {{Country: "FR"}, {Country: "FR"}},
				"US": {{Country: "US"}, {Country: "US"}, {Country: "US"}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ds := &Datasource{
				data: tc.data,
			}
			ds.GroupBy(tc.a)

			if !reflect.DeepEqual(ds.grouped, tc.expected) {
				t.Errorf("getRevenues() = %v, want %v", ds.grouped, tc.expected)
			}
		})
	}
}

func TestDatasource_getAverageRevenue(t *testing.T) {
	testCases := []struct {
		data []*csvData
		want []float64
	}{
		{
			data: []*csvData{
				{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
			},
			want: []float64{1, 1, 1, 1, 1, 1},
		},
		{
			data: []*csvData{
				{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{LTV1: 1, LTV2: 3, LTV3: 5, LTV4: 7, LTV5: 9, LTV6: 11, LTV7: 13},
			},
			want: []float64{1.5, 1.5, 1.5, 1.5, 1.5, 1.5},
		},
		{
			data: []*csvData{
				{LTV1: 10, LTV2: 20, LTV3: 30, LTV4: 40, LTV5: 50, LTV6: 60, LTV7: 70},
				{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
			},
			want: []float64{5.5, 5.5, 5.5, 5.5, 5.5, 5.5},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ds := &Datasource{}
			if got := ds.getAverageRevenue(tc.data); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getAverageRevenue() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDatasource_Prepare(t *testing.T) {
	testCases := []struct {
		name    string
		grouped map[string][]*csvData
		want    map[string][]float64
		wantErr bool
	}{
		{
			name: "",
			grouped: map[string][]*csvData{
				"IT": {
					{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				},
				"US": {
					{LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				},
			},
			want: map[string][]float64{
				"IT": {1, 1, 1, 1, 1, 1},
				"US": {1, 1, 1, 1, 1, 1},
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ds := &Datasource{grouped: tc.grouped}

			got, err := ds.Prepare()
			if (err != nil) != tc.wantErr {
				t.Errorf("Prepare() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Prepare() got = %v, want %v", got, tc.want)
			}
		})
	}
}
