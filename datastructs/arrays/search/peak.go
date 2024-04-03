package search

// PeakSearch searches for the peak values in the array
// A peak is any element A where A[i-1]<= A[i] >= A[i+1]
func PeakSearch(data []int) ([]int, bool) {
	if len(data) == 1 {
		return data, true
	}

	var peaks []int
	found := false
	for i := 0; i < len(data); i++ {
		switch {
		case i == 0:
			if data[i] >= data[i+1] {
				peaks = append(peaks, i)
				found = true
				continue
			}
		case i == len(data)-1:
			if data[i] >= data[i-1] {
				peaks = append(peaks, i)
				found = true
				continue
			}
		case data[i-1] <= data[i] && data[i] >= data[i+1]:
			found = true
			peaks = append(peaks, i)
		}
	}
	return peaks, found
}
