package entity

import (
	"math"
	"testing"
)

func TestWeather_ToResponse(t *testing.T) {
	tests := []struct {
		name     string
		weather  Weather
		expected Response
	}{
		{
			name: "Zero values",
			weather: Weather{
				Current: struct {
					Temp_C float64 "json:\"temp_c\""
					Temp_F float64 "json:\"temp_f\""
				}{Temp_C: 0, Temp_F: 0},
			},
			expected: Response{Temp_C: 0, Temp_F: 0, Temp_K: 273.15},
		},
		{
			name: "Positive temperatures",
			weather: Weather{
				Current: struct {
					Temp_C float64 "json:\"temp_c\""
					Temp_F float64 "json:\"temp_f\""
				}{Temp_C: 25, Temp_F: 77},
			},
			expected: Response{Temp_C: 25, Temp_F: 77, Temp_K: 298.15},
		},
		{
			name: "Negative Celsius",
			weather: Weather{
				Current: struct {
					Temp_C float64 "json:\"temp_c\""
					Temp_F float64 "json:\"temp_f\""
				}{Temp_C: -10, Temp_F: 14},
			},
			expected: Response{Temp_C: -10, Temp_F: 14, Temp_K: 263.15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.weather.ToResponse()
			if math.Abs(resp.Temp_C-tt.expected.Temp_C) > 1e-9 {
				t.Errorf("Temp_C: got %v, want %v", resp.Temp_C, tt.expected.Temp_C)
			}
			if math.Abs(resp.Temp_F-tt.expected.Temp_F) > 1e-9 {
				t.Errorf("Temp_F: got %v, want %v", resp.Temp_F, tt.expected.Temp_F)
			}
			if math.Abs(resp.Temp_K-tt.expected.Temp_K) > 1e-9 {
				t.Errorf("Temp_K: got %v, want %v", resp.Temp_K, tt.expected.Temp_K)
			}
		})
	}
}
