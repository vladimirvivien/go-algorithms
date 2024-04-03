package search

import (
	"reflect"
	"testing"
)

func TestPeakSearch(t *testing.T) {
	tests := map[string]struct {
		data  []int
		found bool
		peaks []int
	}{
		"one peak": {
			data:  []int{1, 3, 20, 4, 1, 0},
			found: true,
			peaks: []int{2},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			peaks, found := PeakSearch(test.data)
			if found != test.found {
				t.Errorf("Unexpected search result: found: %#v: %#v", found, test.data)
			}
			if !reflect.DeepEqual(test.peaks, peaks) {
				t.Errorf("Unexpected peak position mismatched: %#v: %#v", peaks, test.peaks)
			}
		})
	}
}

func TestPeak2DSearch(t *testing.T) {
	tests := map[string]struct {
		data  [][]int
		found bool
		peaks [][]int
	}{
		"one peak": {
			data: [][]int{
				{1, 3, 2, 4},
				{5, 7, 18, 6},
				{9, 10, 11, 12},
				{15, 14, 13, 16},
			},
			found: true,
			peaks: [][]int{{1, 2}, {3, 0}, {3, 3}},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			peaks := Peak2DSearch(test.data)
			if len(peaks) != len(test.peaks) {
				t.Errorf("Unexpected peak count: found %d: expected %#v", len(peaks), len(test.peaks))
			}
			if !reflect.DeepEqual(test.peaks, peaks) {
				t.Errorf("Unexpected peak position mismatched: %#v: %#v", peaks, test.peaks)
			}
		})
	}
}
