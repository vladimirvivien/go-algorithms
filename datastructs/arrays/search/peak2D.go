package search

func Peak2DSearch(data [][]int) [][]int {

	var peaks [][]int

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if (i == 0 || (data[i][j] >= data[i-1][j])) && // top
				(i == len(data)-1 || (data[i][j] >= data[i+1][j])) &&
				(j == 0 || (data[i][j] >= data[i][j-1])) && // left
				(j == len(data[i])-1 || (data[i][j] >= data[i][j+1])) { // right
				peaks = append(peaks, [][]int{{i, j}}...)
			}
		}
	}

	return peaks
}
