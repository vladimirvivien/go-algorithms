package search

import (
	"testing"

	"github.com/vladimirvivien/go-algoritms/datastructs/arrays/sorting"
)

var (
	tests = map[string]struct {
		data   []int
		val    int
		found  bool
		sorted bool
	}{
		"not found": {
			data:  []int{4, 3, 6, 799, 2, 17, 0, 12},
			val:   5,
			found: false,
		},

		"long list unsorted": {
			data:   []int{17, 410, 433, 19, 37, 17, 23, 29, 31, 155, 341, 367, 171, 19, 2, 3, 31, 37, 323, 339, 130, 31, 37, 47, 5, 7, 11, 13, 41, 43, 47, 53, 59, 61, 67, 300, 71, 73, 79, 89},
			val:    300,
			found:  true,
			sorted: false,
		},

		"long list sorted": {
			data:   []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 102, 130, 155, 171, 211, 213, 271, 300, 323, 339, 341, 367, 410, 433},
			val:    271,
			found:  true,
			sorted: true,
		},
	}
)

func TestLinSearch(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, found := LinSearch(test.val, test.data)
			if found != test.found {
				t.Errorf("Unexpected search result: found: %#v:", found)
			}
		})
	}
}

func TestBinSearch(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if !test.sorted {
				sorting.BubbSort(test.data)
			}
			_, found := BinSearch(test.val, test.data)
			if found != test.found {
				t.Errorf("Unexpected search result: found: %#v: %#v", found, test.data)
			}
		})
	}
}

func TestReBinSearch(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if !test.sorted {
				sorting.BubbSort(test.data)
			}
			_, found := ReBinSearch(test.val, 0, len(test.data)-1, test.data)
			if found != test.found {
				t.Errorf("Unexpected search result: found: %#v: %#v", found, test.data)
			}
		})
	}
}
