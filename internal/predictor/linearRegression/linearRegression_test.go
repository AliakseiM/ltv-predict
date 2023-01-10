package linearRegression

import (
	"math"
	"testing"
)

func TestLinearRegression_PredictForDay(t *testing.T) {
	lr := &LinearRegression{}

	testCases := []struct {
		data    []float64
		day     int
		want    float64
		wantErr bool
	}{
		{
			data:    []float64{1, 2, 3, 4, 5},
			day:     7,
			want:    6,
			wantErr: false,
		},
		{
			data:    []float64{1, 2, 3, 4, 5},
			day:     60,
			want:    59,
			wantErr: false,
		},
		{
			data:    []float64{1, 1, 1},
			day:     60,
			want:    1,
			wantErr: false,
		},
		{
			data:    []float64{10, 20, 30, 40, 50, 60},
			day:     10,
			want:    90,
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := lr.PredictForDay(tt.data, tt.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("PredictForDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got = math.Round(got)
			if got != tt.want {
				t.Errorf("PredictForDay() got = %v, want %v", got, tt.want)
			}
		})
	}
}
