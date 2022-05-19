package search

func BinSearch(array []int, value int) int {
	min := 0
	max := len(array) - 1

	for min <= max {
		index := int((max + min) / 2)

		switch {
		case array[index] == value:
			return index
		case value > array[index]:
			min = index + 1
		case value < array[index]:
			max = index - 1
		}

	}
	return -1
}
