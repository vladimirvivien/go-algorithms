package sort

import (
	"fmt"
	"testing"
)

func TestSelSort(t *testing.T) {
	tests := []struct {
		data   []int
		sorted []int
	}{
		{data: []int{10, 5, 12, 6, 7}, sorted: []int{5, 6, 7, 10, 12}},
		{data: []int{100, 5, 2212, 611, 971}, sorted: []int{5, 100, 611, 971, 2212}},
		{data: []int{2100, 1201, 844, 371}, sorted: []int{371, 844, 1201, 2100}},
		{data: []int{100, 100, 100, 100, 100, 100}, sorted: []int{100, 100, 100, 100, 100, 100}},
	}

	for _, test := range tests {
		SelSort(test.data)
		if fmt.Sprintf("%v", test.data) != fmt.Sprintf("%v", test.sorted) {
			t.Errorf("data is not unsorted %v", test.data)
		}
	}
}
