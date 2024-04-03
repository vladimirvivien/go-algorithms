package search

// BinSearch is a simple binary search implementation that
// does not use recursion. It is assumed the data is sorted.

func BinSearch(val int, data []int) (int, bool) {
	lo := 0
	hi := len(data) - 1

	for lo <= hi {
		i := int((lo + hi) / 2)

		switch {
		case data[i] == val:
			return i, true
		case val > data[i]:
			lo = i + 1
		case val < data[i]:
			hi = i - 1
		}
	}

	return -1, false
}

// ReBinSearch does a recursive binary search.
// The initial call should pass the starting position of the search.
func ReBinSearch(val, lo, hi int, data []int) (int, bool) {
	if hi < lo {
		return -1, false
	}

	if lo <= hi {
		i := int((lo + hi) / 2)
		switch {
		case data[i] == val:
			return i, true
		case val > data[i]:
			return ReBinSearch(val, i+1, hi, data)
		case val < data[i]:
			return ReBinSearch(val, lo, i-1, data)
		}
	}

	return -1, false
}
