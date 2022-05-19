package search

import (
	"testing"
)

func TestBinSearch(t *testing.T) {
	var data = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 102, 130, 155, 171, 211, 213, 271, 300, 323, 339, 341, 367, 410, 433}

	index := BinSearch(data, 5)
	if index != 2 {
		t.Fatal("expected value not found")
	}
}
