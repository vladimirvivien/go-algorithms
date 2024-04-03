package search

// LinSearch implements a simple linear search that visits all elements consecutively.
// Returning the position and true if found
func LinSearch(val int, data []int) (int, bool) {
	for i := 0; i < len(data); i++ {
		if val == data[i] {
			return i, true
		}
	}

	return -1, false
}