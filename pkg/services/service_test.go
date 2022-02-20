package services

import (
	"fmt"
	"testing"
)

/*
func TestCountDistanceDriven(t *testing.T) {
	defaultX := 1
	defaultY := 1

	var tests = []struct {
		data CoordinateData
		want float64
	}{
		{CoordinateData{X: 1, Y: 1, Diametr: 4}, 0},
		{CoordinateData{X: 2, Y: -1, Diametr: 4}, 5},
		{CoordinateData{X: -2, Y: 2, Diametr: 4}, 10},
		{CoordinateData{X: -2, Y: -3, Diametr: 4}, 25},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.data)
		t.Run(testname, func(t *testing.T) {
			ans := tt.data.countDistance(float64(defaultX), float64(defaultY))
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
		data CircleCoordinateData
		want string
	}{
		{CircleCoordinateData{X: 1, Y: 1, Diametr: 4}, "USD"},
		{CircleCoordinateData{X: 1, Y: 1, Diametr: 1}, "USD"},
		{CircleCoordinateData{X: 2, Y: -1, Diametr: 4}, "EUR"},
		{CircleCoordinateData{X: 2, Y: -1, Diametr: 9}, "USD"},
		{CircleCoordinateData{X: -2, Y: 2, Diametr: 4}, "EUR"},
		{CircleCoordinateData{X: -2, Y: 2, Diametr: 10}, "USD"},
		{CircleCoordinateData{X: -2, Y: -3, Diametr: 6}, "EUR"},
		{CircleCoordinateData{X: -2, Y: -3, Diametr: 10}, "USD"},
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
