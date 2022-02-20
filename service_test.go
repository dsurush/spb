package test

import (
	"fmt"
	"spb/pkg"
	"testing"
)

func TestCountDistanceDriven(t *testing.T) {
	defaultX := 1
	defaultY := 1

	var tests = []struct {
		data pkg.CoordinateData
		want float64
	}{
		{pkg.CoordinateData{X: 1, Y: 1, Diametr: 4}, 0},
		{pkg.CoordinateData{X: 2, Y: -1, Diametr: 4}, 5},
		{pkg.CoordinateData{X: -2, Y: 2, Diametr: 4}, 10},
		{pkg.CoordinateData{X: -2, Y: -3, Diametr: 4}, 25},
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
