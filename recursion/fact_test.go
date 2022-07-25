package recursion

import "testing"

func TestFactorial(t * testing.T) {
	if fact(0) != 1 && fact(1) != 1 {
		t.Error("base cases failed")
	}

	vals := []int{3, 5, 7, 9, 11}
	results := []int{6, 120, 5040, 362880, 39916800}

	for i, n := range vals {
		if fact(n) != results[i] {
			t.Errorf("%d! should be %d", n, results[i])
		}
	}

}