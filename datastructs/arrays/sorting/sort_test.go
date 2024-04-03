package sorting

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	data   []int
	sorted []int
}{
	"all sorted":   {data: []int{1, 2, 3, 4, 5, 6}, sorted: []int{1, 2, 3, 4, 5, 6}},
	"all unsorted": {data: []int{9, 7, 3, 4, 2, 10}, sorted: []int{2, 3, 4, 7, 9, 10}},
	"twenty values": {
		data:   []int{3, 7, 60, 9, 87, 23, 9, 71, 93, 19, 11, 13, 4, 32, 10, 24, 2, 51, 14, 27, 101, 412, 2, 7},
		sorted: []int{2, 2, 3, 4, 7, 7, 9, 9, 10, 11, 13, 14, 19, 23, 24, 27, 32, 51, 60, 71, 87, 93, 101, 412},
	},
}

func TestBubbSort(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			BubbSort(test.data)
			if !reflect.DeepEqual(test.data, test.sorted) {
				t.Errorf("unexepected result: %#v, %#v", test.data, test.sorted)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			InsertionSort(test.data)
			if !reflect.DeepEqual(test.data, test.sorted) {
				t.Errorf("unexepected result: %#v, %#v", test.data, test.sorted)
			}
		})
	}
}
