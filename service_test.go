package test

import (
	"fmt"
	"spb/pkg/services"
	"testing"
)

/*
func TestCountDistanceDriven(t *testing.T) {
	defaultX := 1
	defaultY := 1

	var tests = []struct {
		data services.CoordinateData
		want float64
	}{
		{services.CoordinateData{X: 1, Y: 1, Diametr: 4}, 0},
		{services.CoordinateData{X: 2, Y: -1, Diametr: 4}, 5},
		{services.CoordinateData{X: -2, Y: 2, Diametr: 4}, 10},
		{services.CoordinateData{X: -2, Y: -3, Diametr: 4}, 25},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.data)
		t.Run(testname, func(t *testing.T) {
			ans := tt.data.CountDistance(float64(defaultX), float64(defaultY))
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}
*/
func TestCheckInCircileDriven(t *testing.T) {

	defaultX := 1
	defaultY := 1

	var tests = []struct {
		data services.CoordinateData
		want string
	}{
		{services.CoordinateData{X: 1, Y: 1, Diametr: 4}, "USD"},
		{services.CoordinateData{X: 1, Y: 1, Diametr: 1}, "USD"},
		{services.CoordinateData{X: 2, Y: -1, Diametr: 4}, "EUR"},
		{services.CoordinateData{X: 2, Y: -1, Diametr: 9}, "USD"},
		{services.CoordinateData{X: -2, Y: 2, Diametr: 4}, "EUR"},
		{services.CoordinateData{X: -2, Y: 2, Diametr: 10}, "USD"},
		{services.CoordinateData{X: -2, Y: -3, Diametr: 6}, "EUR"},
		{services.CoordinateData{X: -2, Y: -3, Diametr: 10}, "USD"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.data)

		t.Run(testname, func(t *testing.T) {

			ans := tt.data.CheckInCircile(float64(defaultX), float64(defaultY))
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}

		})
	}
}
