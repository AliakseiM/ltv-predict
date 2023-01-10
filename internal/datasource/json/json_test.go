package json

import (
	"reflect"
	"testing"

	"github.com/AliakseiM/ltv-predict/internal/models"
)

func Test_jsonData_getRevenues(t *testing.T) {
	testCases := []struct {
		name           string
		input          jsonData
		expectedOutput []float64
	}{
		{
			name: "test case 1",
			input: jsonData{
				Users:      1,
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
			expectedOutput: []float64{10, 10, -30, 0, 0, 0},
		},
		{
			name: "test case 2",
			input: jsonData{
				Users:      1,
				CampaignID: "campaign1",
				Country:    "US",
				LTV1:       10,
				LTV2:       20,
				LTV3:       30,
				LTV4:       40,
				LTV5:       50,
				LTV6:       60,
				LTV7:       0,
			},
			expectedOutput: []float64{10, 10, 10, 10, 10, -60},
		},
		{
			name: "test case 3",
			input: jsonData{
				Users:      1,
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
	ds := NewDatasource("test.json")

	if ds == nil {
		t.Error("Expected Datasource to be not nil")
	}

	if ds.filePath != "test.json" {
		t.Errorf("Expected filePath to be test.json, got %s", ds.filePath)
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
			filePath: "file_not_found.json",
			wantErr:  true,
		},
		{
			name:     "success",
			filePath: "../../../test/data/ok.json",
			wantErr:  false,
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
		data     []*jsonData
		a        models.AggregateType
		expected map[string][]*jsonData
	}{
		{
			name: "correct campaign",
			data: []*jsonData{{CampaignID: "1"}, {CampaignID: "2"}, {CampaignID: "2"}, {CampaignID: "3"}, {CampaignID: "3"}, {CampaignID: "3"}},
			a:    models.AggregateTypeCampaign,
			expected: map[string][]*jsonData{
				"1": {{CampaignID: "1"}},
				"2": {{CampaignID: "2"}, {CampaignID: "2"}},
				"3": {{CampaignID: "3"}, {CampaignID: "3"}, {CampaignID: "3"}},
			},
		},
		{
			name: "correct country",
			data: []*jsonData{{Country: "IT"}, {Country: "FR"}, {Country: "FR"}, {Country: "US"}, {Country: "US"}, {Country: "US"}},
			a:    models.AggregateTypeCountry,
			expected: map[string][]*jsonData{
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

func TestDatasource_getAverageWeightedRevenue(t *testing.T) {
	testCases := []struct {
		data []*jsonData
		want []float64
	}{
		{
			data: []*jsonData{
				{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
			},
			want: []float64{1, 1, 1, 1, 1, 1},
		},
		{
			data: []*jsonData{
				{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				{Users: 300, LTV1: 1, LTV2: 3, LTV3: 5, LTV4: 7, LTV5: 9, LTV6: 11, LTV7: 13},
			},
			want: []float64{1.75, 1.75, 1.75, 1.75, 1.75, 1.75},
		},
		{
			data: []*jsonData{
				{Users: 8, LTV1: 10, LTV2: 20, LTV3: 30, LTV4: 40, LTV5: 50, LTV6: 60, LTV7: 70},
				{Users: 7, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
			},
			want: []float64{5.8, 5.8, 5.8, 5.8, 5.8, 5.8},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ds := &Datasource{}
			if got := ds.getAverageWeightedRevenue(tc.data); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getAverageWeightedRevenue() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDatasource_Prepare(t *testing.T) {
	testCases := []struct {
		name    string
		grouped map[string][]*jsonData
		want    map[string][]float64
		wantErr bool
	}{
		{
			name: "",
			grouped: map[string][]*jsonData{
				"IT": {
					{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
				},
				"US": {
					{Users: 100, LTV1: 1, LTV2: 2, LTV3: 3, LTV4: 4, LTV5: 5, LTV6: 6, LTV7: 7},
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
