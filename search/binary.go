package search

var data = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 102, 130, 155, 171, 211, 213, 271, 300, 323, 339, 341, 367, 410, 433}

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
