package sort

import "fmt"

// SelSort implements a simple selection sort ascending order
func SelSort(data []int) {
	var i, j int
	fmt.Printf("sorting: %v\n", data)
	for i = range data {

		// search and select next smallest
		for j = i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}

}
